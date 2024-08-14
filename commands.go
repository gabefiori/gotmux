// Package gotmux is a simple library for interacting with tmux.
package gotmux

import (
	"fmt"
)

// GetCurrentSession retrieves the name of the currently active tmux session.
func GetCurrentSession() (string, error) {
	cmd, err := NewTmuxCmd("display-message", "-p", "#S")

	if err != nil {
		return "", err
	}

	return cmd.ExecWithOutput()
}

// HasSession checks if a tmux session with the given name exists.
func HasSession(target string) bool {
	cmd, err := NewTmuxCmd("has-session", "-t", target)

	if err != nil {
		return false
	}

	return cmd.Exec() == nil
}

// ListSessions lists all tmux sessions with an optional custom format.
func ListSessions(formatter string) (*TmuxList, error) {
	args := []string{"list-sessions"}

	if formatter != "" {
		args = append(args, "-F", formatter)
	}

	cmd, err := NewTmuxCmd(args...)

	if err != nil {
		return nil, err
	}

	output, err := cmd.ExecWithOutput()

	list := &TmuxList{
		Output:   output,
		hasError: err != nil,
	}

	return list, err
}

// ListWindows lists all windows in the specified tmux session with an optional custom format.
func ListWindows(target string, formatter string) (*TmuxList, error) {
	args := []string{"list-windows", "-t", target}

	if formatter != "" {
		args = append(args, "-F", formatter)
	}

	cmd, err := NewTmuxCmd(args...)

	if err != nil {
		return nil, err
	}

	output, err := cmd.ExecWithOutput()

	list := &TmuxList{
		Output:   output,
		hasError: err != nil,
	}

	return list, err
}

// AddWindow creates a new window in the specified tmux session with the given name.
func AddWindow(target string, name string) error {
	cmd, err := NewTmuxCmd("new-window", "-t", target, "-n", name)

	if err != nil {
		return err
	}

	return cmd.Exec()
}

// AddWindowWithIdx creates a new window in the specified tmux session with the given name and index.
func AddWindowWithIdx(target string, name string, idx uint8) error {
	fmtTarget := fmt.Sprintf("%s:%d", target, idx)
	cmd, err := NewTmuxCmd("new-window", "-t", fmtTarget, "-n", name)

	if err != nil {
		return err
	}

	return cmd.Exec()
}

// SwitchTo switches to the specified tmux session.
func SwitchTo(target string) error {
	cmd, err := NewTmuxCmd("switch-client", "-t", target)
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// AttachOrSwitchTo attaches to or switches to the specified tmux session.
// If inside a tmux session, it switches to the target session; otherwise, it attaches to it.
//
// Warning: This action could replace the current process with tmux, ending the execution of the current program.
func AttachOrSwitchTo(target string) error {
	if IsInsideTmux() {
		return SwitchTo(target)
	}

	return AttachTo(target)
}

// Attach attaches the current terminal to a tmux session.
func Attach() error {
	cmd, err := NewTmuxCmd("attach")

	if err != nil {
		return err
	}

	return cmd.ExecSyscall()
}

// AttachTo attaches the current terminal to the specified tmux session.
//
// Warning: This action will replace the current process with tmux, ending the execution of the current program.
func AttachTo(target string) error {
	cmd, err := NewTmuxCmd("attach-session", "-t", target)

	if err != nil {
		return err
	}

	return cmd.ExecSyscall()
}

// Detach detaches the current terminal from the tmux session.
//
// Warning: This action will replace the current process with tmux, ending the execution of the current program.
func Detach() error {
	cmd, err := NewTmuxCmd("detach")

	if err != nil {
		return err
	}

	return cmd.ExecSyscall()
}

// KillSession terminates the specified tmux session.
func KillSession(target string) error {
	cmd, err := NewTmuxCmd("kill-session", "-t", target)

	if err != nil {
		return err
	}

	return cmd.Exec()
}

// KillServer terminates the tmux server, killing all sessions.
func KillServer() error {
	cmd, err := NewTmuxCmd("kill-server")

	if err != nil {
		return err
	}

	return cmd.Exec()
}
