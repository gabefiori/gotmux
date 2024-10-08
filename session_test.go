package gotmux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSession(t *testing.T) {
	config := &SessionConfig{
		Name: "test-session",
		Dir:  "/tmp",
	}

	session, err := NewSession(config)

	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.True(t, HasSession(config.Name))
	assert.Equal(t, config.Name, session.Name)

	err = KillSession(config.Name)

	assert.NoError(t, err)
}

func TestCreatedSessionKill(t *testing.T) {
	config := &SessionConfig{
		Name: "test-session",
		Dir:  "/tmp",
	}

	session, err := NewSession(config)
	assert.NoError(t, err)

	err = session.Kill()

	assert.NoError(t, err)
	assert.False(t, HasSession(config.Name))
}
