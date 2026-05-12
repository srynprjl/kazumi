package misc

import (
	"fmt"
	"os/exec"
)

func CheckDependencies(command string) (bool, string) {
	_, err := exec.LookPath(command)
	if err != nil {
		return false, fmt.Sprintf("%s not found", command)
	} else {
		return true, ""
	}
}
