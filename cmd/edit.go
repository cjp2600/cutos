package cmd

import (
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit swagger information",
}

func init() {
	rootCmd.AddCommand(editCmd)
}
