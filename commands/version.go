package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "show netselect version information",
	Long:    ``,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("netselect version %s (%s)\n", Version, Build)
	},
}

func init() {
	// RootCmd.AddCommand(versionCmd)
}
