package scrunner

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunScriptWithOutputs(path string) error {
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}

	executableDir := filepath.Dir(executablePath)
	scriptPath := filepath.Join(executableDir, path)

	cmd := exec.Command(bashDefault, scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	return err
}

func RunScriptWithoutOutputs(path string) error {
	cmd := exec.Command(bashDefault, path)

	err := cmd.Run()

	return err
}
