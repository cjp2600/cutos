package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new path",
}

func init() {
	rootCmd.AddCommand(addCmd)
}
