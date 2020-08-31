package commands

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"text/tabwriter"

	"github.com/pgollangi/netselect"
)

// Version is the version for netselect
var Version string

// Build holds the date bin was released
var Build string

// RootCmd is the main root/parent command
var RootCmd = &cobra.Command{
	Use:           "netselect [flags] <host(s)>",
	Short:         "A netselect CLI Tool",
	Long:          `netselect is an open source CLI tool to select the fastest host based on the lowest ICMP latency.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Example: heredoc.Doc(`
		$ netselect m1.example.com m2.example.com m3.example.com
		$ netselect -v
		`),
	RunE: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) error {
	if ok, _ := cmd.Flags().GetBool("version"); ok {
		executeVersionCmd()
		return nil
	} else if len(args) == 0 {
		cmd.Usage()
		return nil
	}

	debug, _ := cmd.Flags().GetBool("debug")
	output, _ := cmd.Flags().GetInt("output")
	concurrent, _ := cmd.Flags().GetInt("concurrent")
	attempts, _ := cmd.Flags().GetInt("attempts")
	privileged, _ := cmd.Flags().GetBool("privileged")

	hosts := make([]*netselect.Host, len(args))

	for i, h := range args {
		hosts[i], _ = netselect.NewHost(h, h)
	}

	selector, err := netselect.NewNetSelector(hosts[:])
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}
	selector.Debug = debug
	selector.Threads = concurrent
	selector.Attempts = attempts
	selector.Privileged = privileged

	result, err := selector.Select()
	if err != nil {
		return err
	}

	// initialize tabwriter
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	for i := 0; i < output; i++ {
		if i >= len(hosts) {
			break
		}
		r := result[i]
		var (
			successPercent, avgRtt int64
			successPackets         string
		)
		if r.Success {
			successPercent = (int64)(100 - r.PacketLoss)
			avgRtt = r.AvgRtt.Milliseconds()
		}
		successPackets = fmt.Sprintf("(%2d/%2d)", r.PacketsRecv, r.PacketsSent)

		fmt.Fprintf(w, "%s \t %d ms\t%d%% ok\t%s\t\n", r.Host.Address, avgRtt, successPercent, successPackets)
	}
	return nil
}

// Execute performs netselect command execution
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.Flags().BoolP("version", "v", false, "show netselect version information")
	RootCmd.Flags().BoolP("debug", "d", false, "show debug information")
	RootCmd.Flags().IntP("concurrent", "c", 1, "use <n> concurrent threads. Default to 3.")
	RootCmd.Flags().IntP("output", "o", 3, "output top ranked <n> results. Default to 3.")
	RootCmd.Flags().IntP("attempts", "a", 3, "no.of ping attempts to perform for each host. Default to 3.")
	RootCmd.Flags().BoolP("privileged", "p", true, `use to send "privileged" raw ICMP ping. Default to TRUE.`)

}

func executeVersionCmd() {
	fmt.Printf("netselect version %s (%s)\n", Version, Build)
	fmt.Println("For more info: pgollangi.com/netselect")
}
