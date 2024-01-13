package utils

import (
	// import built-in packages
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":

		exec.Command("cmd", "/c", "cls")
	default:
		exec.Command("export TERM=xterm").Run()
		exec.Command("clear").Run()
	}
}
