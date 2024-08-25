# Go tmux

A simple library for interacting with [tmux](https://github.com/tmux/tmux) using Go.

Currently, the library supports only a few basic functionalities and is not stable yet. Additional features will be added as needed.

## Getting Started

To install the library, use the following command:
```sh
go get -u github.com/gabefiori/gotmux
```

Here’s an example of how to use the library:
```go
package main

import (
	"fmt"
	"github.com/gabefiori/gotmux"
	"log"
)

func main() {
	// Create a new tmux session
	session, err := gotmux.NewSession(&gotmux.SessionConfig{
		Name:       "session-name",
		WindowName: "window-name", // Optional name for the window
		Dir:        "/tmp",        // Optional working directory for the session
	})

	if err != nil {
		log.Fatal(err)
	}

	// Most commands now return only an error
	err = session.AddWindow("new-window")

	if err != nil {
		log.Fatal(err)
	}

	// Switch or attach to the created session.
	// Warning: If the session is attached, code execution will stop
	err = session.AttachOrSwitch()

	if err != nil {
		log.Fatal(err)
	}

	// Kill another tmux session if it exists
	if gotmux.HasSession("other") {
		gotmux.KillSession("other")
	}

	// List all sessions with a custom format and print them
	allSessions, err := gotmux.ListSessions("#S")

	if err != nil {
		log.Fatal(err)
	}

	for _, s := range allSessions.Iter() {
		fmt.Println(s)
	}

	// Create a custom command
	cmd, err := NewTmuxCmd("new-session", "-d", "-s", "new session")

	if err != nil {
		return err
	}

	err = cmd.Exec()                   // Execute
	output, err = cmd.ExecWithOutput() // Execute with output
	err = cmd.ExecSyscall()            // Execute with syscall
}
```

Be sure to check the code for the complete documentation and more examples.
