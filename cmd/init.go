package cmd

import (
	"github.com/cjp2600/cutos/interactor"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialization of a new document, filling in of base values",
	RunE:   interactor.InitializationCmd,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
