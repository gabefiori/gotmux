package gotmux

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// RunCmd executes the specified tmux command with the given arguments and returns the combined output and error.
//
// Use this function when you need to capture the output of the tmux command and when the tmux command is non-interactive.
// It is suitable for commands like creating, killing, or listing tmux sessions.
//
// This function uses exec.Command to run the tmux command and captures both standard output and standard error.
//
// Example:
//
//	output, err := RunCmd("list-sessions", "-F", "#S")
//	if err != nil {
//	    log.Fatalf("Error running tmux command: %v", err)
//	}
//	fmt.Println("Output:", output)
func RunCmd(arg ...string) (string, error) {
	tmuxPath, err := exec.LookPath("tmux")

	if err != nil {
		panic(err)
	}

	output, err := exec.Command(tmuxPath, arg...).CombinedOutput()
	fmtOutput := strings.TrimSpace(string(output))

	return fmtOutput, err
}

// RunCmdWithSyscall replaces the current process with a tmux process running with the given arguments.
//
// Use this function when you want to completely hand over control of the current process to tmux. It is particularly
// useful in scenarios where you want the Go program to run tmux directly and exit the Go process, such as in
// scripting environments or when starting tmux as part of an automated setup.
//
// This function uses syscall.Exec to replace the current process with tmux.
//
// Example:
//	// Attach to an existing tmux session named "my-session"
//	err := RunCmdWithSyscall("attach-session", "-t", "my-session")
//	if err != nil {
//	    log.Fatalf("Error running tmux with syscall: %v", err)
//	}
//	// This code will not be reached if syscall.Exec succeeds
func RunCmdWithSyscall(arg ...string) error {
	tmuxPath, err := exec.LookPath("tmux")

	if err != nil {
		panic(err)
	}

	args := append([]string{tmuxPath}, arg...)

	if err :=  syscall.Exec(tmuxPath, args, os.Environ()); err != nil {
		return err
	}

	return nil
}
