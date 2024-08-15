package gotmux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createSession(name string) error {
	cmd, err := NewTmuxCmd("new-session", "-d", "-s", name)
	if err != nil {
		return err
	}
	return cmd.Exec()
}

func TestGetCurrentSession(t *testing.T) {
	if !IsInsideTmux() {
		t.Skip("Not running inside a tmux session")
	}
	session, err := GetCurrentSession()
	assert.NoError(t, err)
	assert.NotEmpty(t, session)
}

func TestKillSession(t *testing.T) {
	sessionName := "test-kill-session"

	err := createSession(sessionName)
	assert.NoError(t, err)

	err = KillSession(sessionName)

	assert.NoError(t, err)
	assert.False(t, HasSession(sessionName), "Expected session %s to be killed", sessionName)
}

func TestHasSession(t *testing.T) {
	sessionName := "test-session"

	err := createSession(sessionName)
	assert.NoError(t, err)

	defer KillSession(sessionName)

	assert.True(t, HasSession(sessionName))
	assert.False(t, HasSession("nonexistent-session"))
}

func TestListSessions(t *testing.T) {
	session1 := "test-session-1"
	session2 := "test-session-2"

	err := createSession(session1)
	assert.NoError(t, err)

	defer KillSession(session1)

	err = createSession(session2)
	assert.NoError(t, err)

	defer KillSession(session2)

	sessions, err := ListSessions("")

	assert.NoError(t, err)
	assert.Contains(t, sessions.Output, session1)
	assert.Contains(t, sessions.Output, session2)
}

func TestAddWindow(t *testing.T) {
	sessionName := "test-new-window-session"

	err := createSession(sessionName)
	assert.NoError(t, err)

	defer KillSession(sessionName)

	err = AddWindow(sessionName, "new-window")
	assert.NoError(t, err)

	windows, err := ListWindows(sessionName, "")

	assert.NoError(t, err)
	assert.Contains(t, windows.Output, "new-window")
}

func TestAddWindowWithIdx(t *testing.T) {
	sessionName := "test-new-window-session"

	err := createSession(sessionName)
	assert.NoError(t, err)

	defer KillSession(sessionName)

	err = AddWindowWithIdx(sessionName, "new-window", uint8(10))
	assert.NoError(t, err)

	windows, err := ListWindows(sessionName, "")

	assert.NoError(t, err)
	assert.Contains(t, windows.Output, "10: new-window")
}

func TestListWindows(t *testing.T) {
	sessionName := "test-window-session"
	err := createSession(sessionName)

	assert.NoError(t, err)

	defer KillSession(sessionName)

	err = AddWindow(sessionName, "window1")
	assert.NoError(t, err)

	err = AddWindow(sessionName, "window2")
	assert.NoError(t, err)

	windows, err := ListWindows(sessionName, "")

	assert.NoError(t, err)
	assert.Contains(t, windows.Output, "window1")
	assert.Contains(t, windows.Output, "window2")
}
