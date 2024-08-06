package gotmux

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	config := &SessionConfig{
		Name: "test-session",
		Dir:  "/tmp",
	}

	session, _, err := NewSession(config)

	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.True(t, HasSession(config.Name))
	assert.Equal(t, config.Name, session.Name)

	_, err = KillSession(config.Name)
	assert.NoError(t, err)

}

func TestCreatedSessionKill(t *testing.T) {
	config := &SessionConfig{
		Name: "test-session",
		Dir:  "/tmp",
	}

	session, _, err := NewSession(config)

	assert.NoError(t, err)

	_, err = session.Kill()

	assert.NoError(t, err)
	assert.False(t, HasSession(config.Name))
}
