package gotmux

import (
	"os"
	"os/exec"
	"strings"
)

func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

func IsTmuxInstalled() bool {
	_, err := exec.LookPath("tmux")
	return err == nil
}

func ValidateSessionName(name string) bool {
	if name == "" ||
		strings.Contains(name, ".") ||
		strings.Contains(name, ":") {
		return false
	}

	return true
}
