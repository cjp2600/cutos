package cmd

import (
	"github.com/cjp2600/cutos/interactor"
	"github.com/cjp2600/cutos/parser"
	"github.com/spf13/cobra"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "add new path from curl",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := interactor.ListeningCmd(cmd, args, parser.CurlType)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	addCmd.AddCommand(curlCmd)
}
