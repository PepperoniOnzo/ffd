package internal

import (
	"fmt"
	"os"
)

var ffdConfigsFolder = "/.ffd"
var ffdConfigsName = "/ffd.yaml"
var gitignoreFile = "/.gitignore"

func GetSettingFileDirectory() string {
	return GetConfigsDirectory() + ffdConfigsName
}

func GetGitignoreFileDirectory() string {
	return GetRunningDirectory() + gitignoreFile
}

func GetConfigsDirectory() string {

	return GetRunningDirectory() + ffdConfigsFolder
}

func GetRunningDirectory() string {
	res, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("cant get running directory. Err: %w", err))
	}

	return res
}
