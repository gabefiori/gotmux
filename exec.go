package gotmux

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// TmuxCmd represents a tmux command with its path and arguments.
type TmuxCmd struct {
	path string
	args []string
}

// NewTmuxCmd creates a new TmuxCmd instance with the given arguments.
// It looks up the tmux executable path and returns an error if tmux is not found.
func NewTmuxCmd(arg ...string) (*TmuxCmd, error) {
	tmuxPath, err := exec.LookPath("tmux")

	if err != nil {
		return nil, err
	}

	return &TmuxCmd{
		path: tmuxPath,
		args: arg,
	}, nil
}

// Exec executes the tmux command without returning its output.
// It returns an error if the command execution fails.
func (t *TmuxCmd) Exec() error {
	output, err := exec.Command(t.path, t.args...).CombinedOutput()

	if err != nil {
		err = errors.New(strings.TrimSpace(string(output)))
	}

	return err
}

// ExecWithOutput executes the tmux command and returns its output as a string.
// It returns both the output and an error if the command execution fails.
func (t *TmuxCmd) ExecWithOutput() (string, error) {
	output, err := exec.Command(t.path, t.args...).CombinedOutput()
	fmtOutput := strings.TrimSpace(string(output))

	if err != nil {
		err = errors.New(fmtOutput)
	}

	return fmtOutput, err
}

// ExecSyscall replaces the current process with a new tmux process.
// It uses the syscall.Exec function to execute tmux with the given arguments.
// This function does not return unless there's an error in starting the new process.
func (t *TmuxCmd) ExecSyscall() error {
	args := append([]string{t.path}, t.args...)
	return syscall.Exec(t.path, args, os.Environ())
}
