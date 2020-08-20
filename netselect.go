package netselect

import (
	"fmt"
	"net"
	"runtime"
	"sort"
	"time"

	"github.com/pgollangi/go-ping"
)

type Selector struct {
	Mirrors    []string
	Debug      bool
	Attempts   int
	Timeout    time.Duration
	Privileged bool
}

type MirrorStats struct {
	Mirror string

	Success bool

	Error error
	// PacketsRecv is the number of packets received.
	PacketsRecv int

	// PacketsSent is the number of packets sent.
	PacketsSent int

	// PacketLoss is the percentage of packets lost.
	PacketLoss float64

	// IPAddr is the address of the host being pinged.
	IPAddr *net.IPAddr

	// Addr is the string address of the host being pinged.
	Addr string

	// Rtts is all of the round-trip times sent via this pinger.
	Rtts []time.Duration

	// MinRtt is the minimum round-trip time sent via this pinger.
	MinRtt time.Duration

	// MaxRtt is the maximum round-trip time sent via this pinger.
	MaxRtt time.Duration

	// AvgRtt is the average round-trip time sent via this pinger.
	AvgRtt time.Duration

	// StdDevRtt is the standard deviation of the round-trip times sent via
	// this pinger.
	StdDevRtt time.Duration
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func NewSelector(mirrors []string) (*Selector, error) {
	return &Selector{
		Mirrors:    mirrors,
		Attempts:   3,
		Timeout:    time.Second * 30,
		Privileged: isWindows(),
	}, nil
}

func executePing(mirror string, s *Selector) *MirrorStats {
	pinger, err := ping.NewPinger(mirror)
	if err != nil {
		return &MirrorStats{
			Mirror:  mirror,
			Success: false,
			Error:   err,
		}
	}
	pinger.Timeout = s.Timeout
	pinger.Count = s.Attempts
	pinger.Debug = s.Debug
	pinger.SetPrivileged(s.Privileged)
	if s.Debug {
		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v \n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
		}
		pinger.OnFinish = func(stats *ping.Statistics) {
			fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
			fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
				stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}
	}
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	return &MirrorStats{
		Mirror:      mirror,
		Success:     true,
		PacketsRecv: stats.PacketsRecv,
		PacketsSent: stats.PacketsSent,
		PacketLoss:  stats.PacketLoss,
		IPAddr:      stats.IPAddr,
		Addr:        stats.Addr,
		Rtts:        stats.Rtts,
		MinRtt:      stats.MinRtt,
		MaxRtt:      stats.MaxRtt,
		AvgRtt:      stats.AvgRtt,
		StdDevRtt:   stats.StdDevRtt,
	}
}

type AllResults []*MirrorStats

func (r AllResults) Len() int           { return len(r) }
func (r AllResults) Less(i, j int) bool { return r[i].AvgRtt < r[j].AvgRtt }
func (r AllResults) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (s *Selector) Select() []*MirrorStats {

	pingResults := []*MirrorStats{}

	for _, mirror := range s.Mirrors {
		r := executePing(mirror, s)
		pingResults = append(pingResults, r)
	}

	sort.Sort(AllResults(pingResults))
	return pingResults
}
