
package cmd

import (
	"github.com/cjp2600/cutos/interactor"

	"github.com/spf13/cobra"
)

// basicCmd represents the basic command
var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "Edit basic information",
	RunE: interactor.BasicCmd,
}

func init() {
	editCmd.AddCommand(basicCmd)
}
