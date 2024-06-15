package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ffd",
	Short: "FFD (fast flutter delivery) is CLI for flutter apps delivery",
	Long: `A Fast and Flexible flutter delivery CLI
		Complete documentation is available at http://github.com/PepperoniOnzo/ffd`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
