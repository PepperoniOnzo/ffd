package internal

import (
	"fmt"
	"os"
)

var ffdConfigsFolder = "/.ffd"
var ffdConfigsName = "/ffd.yaml"
var gitignoreFile = "/.gitignore"

type DirectoryUtil struct {
	Path string
}

func NewDirectoryUtil() *DirectoryUtil {
	res, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("cant get running directory. Err: %w", err))
	}

	return &DirectoryUtil{
		Path: res,
	}
}

func (du *DirectoryUtil) GetConfigFolder() string {
	return du.Path + ffdConfigsFolder
}

func (du *DirectoryUtil) GetConfigFile() string {
	return du.Path + ffdConfigsFolder + ffdConfigsName
}

func (du *DirectoryUtil) GetGitignoreFile() string {
	return du.Path + gitignoreFile
}
