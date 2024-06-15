package cmd

import (
	"fmt"
	"os"

	"github.com/PepperoniOnzo/ffd/internal"
	"github.com/PepperoniOnzo/ffd/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfigs)
	rootCmd.AddCommand(initCommand)
}

func initConfigs() {
	viper.SetConfigName("ffd")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(internal.GetConfigsDirectory())

	if err := viper.ReadInConfig(); err != nil {
		return
	}

	if err := viper.Unmarshal(&configs); err != nil {
		return
	}
}

var configs *models.EnvConfigs
var version string = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "ffd",
	Short: "FFD (fast flutter delivery) is CLI for flutter apps delivery",
	Long: `A Fast and Flexible flutter delivery CLI
		Complete documentation is available at http://github.com/PepperoniOnzo/ffd`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// cmd.Help()
			initCommand.Run(cmd, args)
			os.Exit(0)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
