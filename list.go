package gotmux

import "strings"

// TmuxList represents the output of a tmux command and its error state.
type TmuxList struct {
	Output   string // Output contains the raw string output from a tmux command.
	hasError bool   // hasError indicates whether the tmux command resulted in an error.
}

// Iter splits the tmux command output into a slice of strings, each representing a line.
// If the tmux command resulted in an error, it returns an empty slice.
func (tl *TmuxList) Iter() []string {
	if tl.hasError {
		return []string{}
	}

	return strings.Split(tl.Output, "\n")
}

