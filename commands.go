// Package gotmux is a simple library for interacting with tmux.
package gotmux

import (
	"fmt"
)

// GetCurrentSession retrieves the name of the currently active tmux session.
func GetCurrentSession() (string, error) {
	return RunCmd("display-message", "-p", "#S")
}

// HasSession checks if a tmux session with the given name exists.
func HasSession(target string) bool {
	_, err := RunCmd("has-session", "-t", target)
	return err == nil
}

// ListSessions lists all tmux sessions with an optional custom format.
func ListSessions(formatter string) (*TmuxList, error) {
	cmd := []string{"list-sessions"}

	if formatter != "" {
		cmd = append(cmd, "-F", formatter)
	}

	output, err := RunCmd(cmd...)
	list := &TmuxList{
		Output:   output,
		hasError: err != nil,
	}

	return list, err
}

// ListWindows lists all windows in the specified tmux session with an optional custom format.
func ListWindows(target string, formatter string) (*TmuxList, error) {
	cmd := []string{"list-windows", "-t", target}

	if formatter != "" {
		cmd = append(cmd, "-F", formatter)
	}

	output, err := RunCmd(cmd...)
	list := &TmuxList{
		Output:   output,
		hasError: err != nil,
	}

	return list, err
}

// AddWindow creates a new window in the specified tmux session with the given name.
func AddWindow(target string, name string) (string, error) {
	return RunCmd("new-window", "-t", target, "-n", name)
}

// AddWindowWithIdx creates a new window in the specified tmux session with the given name and index.
func AddWindowWithIdx(target string, name string, idx uint8) (string, error) {
	fmtTarget := fmt.Sprintf("%s:%d", target, idx)
	return RunCmd("new-window", "-t", fmtTarget, "-n", name)
}

// SwitchTo switches to the specified tmux session.
func SwitchTo(target string) (string, error) {
	return RunCmd("switch-client", "-t", target)
}

// AttachOrSwitchTo attaches to or switches to the specified tmux session.
// If inside a tmux session, it switches to the target session; otherwise, it attaches to it.
//
// Warning: This action could replace the current process with tmux, ending the execution of the current program.
func AttachOrSwitchTo(target string) (string, error) {
	if IsInsideTmux() {
		return SwitchTo(target)
	}

	err := AttachTo(target)
	return "", err
}

// Attach attaches the current terminal to a tmux session.
func Attach() error {
	return RunCmdWithSyscall("attach")
}

// AttachTo attaches the current terminal to the specified tmux session.
//
// Warning: This action will replace the current process with tmux, ending the execution of the current program.
func AttachTo(target string) error {
	return RunCmdWithSyscall("attach-session", "-t", target)
}

// Detach detaches the current terminal from the tmux session.
//
// Warning: This action will replace the current process with tmux, ending the execution of the current program.
func Detach() error {
	return RunCmdWithSyscall("detach")
}

// KillSession terminates the specified tmux session.
func KillSession(target string) (string, error) {
	return RunCmd("kill-session", "-t", target)
}

// KillServer terminates the tmux server, killing all sessions.
func KillServer() (string, error) {
	return RunCmd("kill-server")
}
