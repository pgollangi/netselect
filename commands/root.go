package commands

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/pgollangi/go-ping"
	"github.com/spf13/cobra"
)

// RootCmd is the main root/parent command
var RootCmd = &cobra.Command{
	Use:           "netselect [flags] <mirror(s)>",
	Short:         "A netselect CLI Tool",
	Long:          `netselect is an open source CLI tool to select the fastest mirror based on the lowest ICMP latency.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Example: heredoc.Doc(`
	$ netselect -s 
	$ netselect -v
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// if ok, _ := cmd.Flags().GetBool("version"); ok {
		// 	// versionCmd.Run(cmd, args)
		// 	return
		// }
		pinger, err := ping.NewPinger(args[0])
		if err != nil {
			panic(err)
		}
		pinger.Count = 10
		pinger.SetPrivileged(true)
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
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		fmt.Println(stats)
	},
}

// Start netselect command execution
func Execute() error {
	RootCmd.Flags().BoolP("version", "v", false, "show netselect version information")
	return RootCmd.Execute()
}
