package cmdchecker

import (
	"fmt"
	"os/exec"

	"github.com/PepperoniOnzo/ffd/internal"
)

func existCommand(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

func RunCommandChecker() {
	flutterExist := existCommand(flutter)
	firebaseCliExist := existCommand(firebaseCli)

	fmt.Println("Doctor summary")
	internal.LogStatus("Flutter - app development", flutterExist)
	internal.LogStatus("Firebase - app distribution", firebaseCliExist)
}
