package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "v0.0.5"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print current version",
	DisableFlagsInUseLine: true,
	Long:                  ``,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
