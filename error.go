package gotmux

import "strings"

// TmuxError represents different types of errors that can occur with tmux commands.
type TmuxError int

const (
	Unknown TmuxError = iota // An unknown error occurred
	SessionNotCreated       // The session could not be created
	NestedSession           // Sessions should be nested with care
	DuplicateSession        // The session name is duplicated
	SessionNotFound         // The session was not found
	CommandFailed           // The command failed to execute
	InvalidArgument         // An invalid argument was provided
	PermissionDenied        // Permission was denied
	Timeout                 // The command timed out
)

// IdentifyError classifies the error based on the command output.
func IdentifyError(output string) TmuxError {
	if strings.Contains(output, "sessions should be nested with care") {
		return NestedSession
	}

	if strings.Contains(output, "duplicate session name") {
		return DuplicateSession
	}

	if strings.Contains(output, "session not found") {
		return SessionNotFound
	}

	if strings.Contains(output, "command failed") {
		return CommandFailed
	}

	if strings.Contains(output, "invalid argument") {
		return InvalidArgument
	}

	if strings.Contains(output, "permission denied") {
		return PermissionDenied
	}

	if strings.Contains(output, "timed out") {
		return Timeout
	}

	return Unknown
}

