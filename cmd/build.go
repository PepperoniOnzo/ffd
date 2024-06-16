package cmd

import (
	"fmt"

	"github.com/PepperoniOnzo/ffd/internal"
	"github.com/PepperoniOnzo/ffd/internal/scrunner"
	"github.com/spf13/cobra"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "Build flutter project",
	Run: func(cmd *cobra.Command, args []string) {
		if configs == nil {
			internal.LogError("Init ffd to use this command.")
			return
		}

		err := scrunner.RunScriptWithOutputs(scrunner.Build)

		if err != nil {
			internal.LogError(fmt.Sprintf("Failed to build app.\nDetails: %v", err))
			return
		}

		internal.LogSuccess("Successfully builded app.")
	},
}
