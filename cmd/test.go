package cmd

import (
	"github.com/PepperoniOnzo/ffd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type FlavorsTest struct {
	AndroidFlavorName string `yaml:"android_flavor_name"`
	IosFlavorName     string `yaml:"ios_flavor_name"`
	FirebaseId        string `yaml:"firebase_id"`
	AppleId           string `yaml:"apple_id"`
}

var testCommand = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {

		flavorTest := &FlavorsTest{
			AndroidFlavorName: "dsds",
			IosFlavorName:     "dassd",
			FirebaseId:        "sdaw423",
			AppleId:           "fweacs",
		}

		flavorTest2 := &FlavorsTest{
			AndroidFlavorName: "dsds",
			IosFlavorName:     "dassd",
			FirebaseId:        "sdaw423",
			AppleId:           "fweacs",
		}

		flavors := map[string]FlavorsTest{"dev": *flavorTest, "stage": *flavorTest2}

		viper.Set("flavors", flavors)
		if err := viper.WriteConfig(); err != nil {
			internal.LogError("Failed to create settings file.")
			return
		}
	},
}
