package cmd

import (
	"github.com/PepperoniOnzo/ffd/internal/cmdchecker"
	"github.com/spf13/cobra"
)

var doctorCommand = &cobra.Command{
	Use:   "doctor",
	Short: "Check if everything installed for using ffd",
	Run: func(cmd *cobra.Command, args []string) {
		cmdchecker.RunCommandChecker()
	},
}
