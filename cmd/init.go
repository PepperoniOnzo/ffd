package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/PepperoniOnzo/ffd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Init ffd for project",
	Run: func(cmd *cobra.Command, args []string) {

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				//
				// 	Creating settings file
				//
				if err := os.Mkdir(internal.GetConfigsDirectory(), os.ModePerm); err != nil {
					internal.LogError("Failed to create setting folder.")
					return
				}

				if _, err := os.Create(internal.GetSettingFileDirectory()); err != nil {
					internal.LogError("Failed to create setting file.")
					return
				}

				viper.Set("version", version)
				if err := viper.WriteConfig(); err != nil {
					internal.LogError("Failed to create settings file.")
					return
				}

				internal.LogSuccess("Successfully initialized ffd.")

				//
				// Updating gitignore if exist
				//
				var gitignoreFile = internal.GetGitignoreFileDirectory()
				if _, err := os.Stat(gitignoreFile); os.IsNotExist(err) {
					return
				}

				fileText, err := os.ReadFile(gitignoreFile)
				if err != nil {
					return
				}

				if strings.Contains(string(fileText), ".ffd\n") {
					return
				}

				file, err := os.Open(gitignoreFile)
				if err != nil {
					file.Close()
					return
				}
				defer file.Close()

				writer := bufio.NewWriter(file)
				defer writer.Flush()

				writer.WriteString("\n ffd related\n.ffd\n")

				writer.Flush()
				file.Close()
				return
			} else {
				internal.LogError("ffd file is broken. Delete file and call again.")
				return
			}
		}

		internal.LogMedium("ffd already initialized.")
	},
}
