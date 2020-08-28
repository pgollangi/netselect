package netselect

import (
	"fmt"
	"net"
	"net/url"
	"runtime"
	"sort"
	"time"

	"github.com/pgollangi/go-ping"
)

// NetSelector represents the instance of a NetSelector
type NetSelector struct {
	Hosts      []*Host
	Debug      bool
	Attempts   int
	Timeout    time.Duration
	Privileged bool
	Threads    int
}

// Host represents a input address to NetSelector
type Host struct {
	// Unique ID
	ID string
	// Address of the Host. If URL provided, Host name will be extracted.
	Address string
}

// HostStats represents the results of one particular host
type HostStats struct {
	Host *Host

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

func sanitizeHost(host *Host) error {
	_, err := url.ParseRequestURI(host.Address)
	if err == nil {
		// Its a URL
		u, err := url.Parse(host.Address)
		if err != nil {
			return err
		}
		if u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("invalid host address %s", host.Address)
		}
		host.Address = u.Hostname()
	}

	return nil
}

// NewHost creates and returns new Host instance
func NewHost(id string, address string) (host *Host, err error) {
	// TODO validate address
	host = &Host{
		ID:      id,
		Address: address,
	}
	err = sanitizeHost(host)
	return host, err
}

// NewNetSelector instantiate new instance of NetSelector
func NewNetSelector(hosts []*Host) (*NetSelector, error) {
	return &NetSelector{
		Hosts:      hosts,
		Attempts:   3,
		Timeout:    time.Second * 30,
		Privileged: isWindows(),
	}, nil
}

func executePing(host *Host, s *NetSelector) *HostStats {
	pinger, err := ping.NewPinger(host.Address)
	if err != nil {
		return &HostStats{
			Host:    host,
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

	return &HostStats{
		Host:        host,
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

type allResults []*HostStats

func (r allResults) Len() int           { return len(r) }
func (r allResults) Less(i, j int) bool { return r[i].AvgRtt < r[j].AvgRtt }
func (r allResults) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Select finds the faster hosts among the provided inputs, and sort the resulted host in ASC order
func (s *NetSelector) Select() ([]*HostStats, error) {
	return s.performSelection()
}

func (s *NetSelector) performSelection() ([]*HostStats, error) {

	hosts := s.Hosts

	for _, host := range hosts {
		err := sanitizeHost(host)
		if err != nil {
			return nil, err
		}
	}

	mLen := len(hosts)

	jobs := make(chan *Host, mLen)
	results := make(chan *HostStats, mLen)

	pingResults := []*HostStats{}

	for t := 0; t < s.Threads; t++ {
		go func() {
			for host := range jobs {
				r := executePing(host, s)
				results <- r
			}
		}()
	}
	for _, host := range hosts {
		jobs <- host
	}
	close(jobs)

	success := []*HostStats{}
	failed := []*HostStats{}

	for range hosts {
		result := <-results
		if result.Success {
			success = append(success, result)
		} else {
			failed = append(failed, result)
		}
		pingResults = append(pingResults, result)
	}

	sort.Sort(allResults(success))
	pingResults = append(success, failed...)
	return pingResults, nil
}

func worker(s *NetSelector, jobs <-chan *Host, results chan<- *HostStats) {
	for host := range jobs {
		r := executePing(host, s)
		results <- r
	}
}
