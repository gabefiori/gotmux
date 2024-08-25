package gotmux

import (
	"errors"

	"github.com/mitchellh/go-homedir"
)

// Session represents a tmux session.
type Session struct {
	Name string // Name of the tmux session
	Dir  string // Working directory for the session
}

// SessionConfig holds the configuration options for creating a new tmux session.
type SessionConfig struct {
	Name       string // Name of the tmux session
	Dir        string // Working directory for the session (optional)
	WindowName string // Name of the initial window in the session (optional)
}

// NewSession creates a new tmux session with the specified configuration.
// If the session name is invalid, an error is returned.
func NewSession(config *SessionConfig) (*Session, error) {
	if !ValidateSessionName(config.Name) {
		return nil, errors.New("invalid session name")
	}

	args := []string{"new-session", "-d", "-s", config.Name}

	if config.Dir != "" {
		expanded, err := homedir.Expand(config.Dir)

		if err != nil {
			return nil, err
		}

		args = append(args, "-c", expanded)
	}

	if config.WindowName != "" {
		args = append(args, "-n", config.WindowName)
	}

	cmd, err := NewTmuxCmd(args...)

	if err != nil {
		return nil, err
	}

	err = cmd.Exec()

	if err != nil {
		return nil, err
	}

	newSession := &Session{
		Name: config.Name,
		Dir:  config.Dir,
	}

	return newSession, nil
}

// AddWindow adds a new window with the specified name to the session.
func (s *Session) AddWindow(name string) error {
	return AddWindow(s.Name, name)
}

// AddWindowWithIdx adds a new window with the specified name and index to the session.
func (s *Session) AddWindowWithIdx(name string, idx uint8) error {
	return AddWindowWithIdx(s.Name, name, idx)
}

// Kill terminates the session.
func (s *Session) Kill() error {
	return KillSession(s.Name)
}

// Attach attaches the current terminal to the session.
func (s *Session) Attach() error {
	return AttachTo(s.Name)
}

// Switch switches to the session.
func (s *Session) Switch() error {
	return SwitchTo(s.Name)
}

// AttachOrSwitch attaches to the session if not already inside a tmux session,
// or switches to it if inside tmux.
func (s *Session) AttachOrSwitch() error {
	return AttachOrSwitchTo(s.Name)
}
