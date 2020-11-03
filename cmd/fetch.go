package cmd

import (
	"github.com/cjp2600/cutos/interactor"
	"github.com/cjp2600/cutos/parser"
	"github.com/spf13/cobra"
)

// fetchCmd represents the curl command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "add new path from fetch",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := interactor.ListeningCmd(cmd, args, parser.FetchType)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	addCmd.AddCommand(fetchCmd)
}
