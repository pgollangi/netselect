package netselect

import (
	"fmt"
	"net"
	"time"

	"github.com/pgollangi/go-ping"
)

type Selector struct {
	Mirrors []string
	Debug   bool
}

type Stats struct {
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

func NewSelector(mirrors []string) (*Selector, error) {
	return &Selector{
		Mirrors: mirrors,
	}, nil
}

func DoPing(mirror string, s *Selector) *MirrorStats {
	pinger, err := ping.NewPinger(mirror)
	if err != nil {
		return &MirrorStats{
			Mirror:  mirror,
			Success: false,
			Error:   err,
		}
	}
	pinger.Timeout = time.Second * 30
	pinger.Count = 3
	pinger.Debug = s.Debug
	pinger.SetPrivileged(true)
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

func worker(id int, jobs <-chan string, results chan<- *MirrorStats) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		stats := DoPing(j, nil)
		fmt.Println("worker", id, "finished job", j)
		results <- stats
	}
}

func (s *Selector) Select() []*MirrorStats {
	pingResults := []*MirrorStats{}

	for _, mirror := range s.Mirrors {
		r := DoPing(mirror, s)
		pingResults = append(pingResults, r)
	}

	// // numJobs := len(s.mirrors)
	// jobs := make(chan string)
	// results := make(chan *MirrorStats)

	// for w := 1; w <= 2; w++ {
	// 	go worker(w, jobs, results)
	// }

	// for _, mirror := range s.mirrors {
	// 	jobs <- mirror
	// }

	// close(jobs)

	// // for {
	// // 	r, ok := <-results
	// // 	fmt.Print(ok)
	// // 	pingResults = append(pingResults, r)
	// // 	if ok != true {
	// // 		break
	// // 	}
	// // }
	// fmt.Print("JOBS DONE")
	// for _, mirror := range s.mirrors {
	// 	fmt.Print(mirror)
	// 	pingResults = append(pingResults, <-results)
	// }

	// for r := range results {
	// 	pingResults = append(pingResults, r)
	// }
	// sort.Sort(AllResults(pingResults))
	return pingResults
}
