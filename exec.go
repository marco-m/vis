package vis

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func ExecOutput(env []string, name string, args ...string) (string, string, error) {
	// We don't use cmd.String() because we don't want the full path of 'name'.
	Out(cmdString(name, args...))
	cmd := exec.Command(name, args...)
	cmd.Env = env
	// Output runs the command and returns its standard output. Any returned
	// error will usually be of type *ExitError. If c.Stderr was nil, Output
	// populates [ExitError.Stderr].
	stdout, err := cmd.Output()
	var stderr []byte
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			stderr = exitError.Stderr
			exitError.Stderr = nil
			return string(stdout), string(stderr),
				fmt.Errorf("%s: %w", name, err)
		}
		return string(stdout), "", fmt.Errorf("%s: %w", name, err)
	}
	return string(stdout), "", nil
}

func ExecRun(dir string, env []string, name string, args ...string) error {
	// We don't use cmd.String() because we don't want the full path of 'name'.
	Out(cmdString(name, args...))
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return nil
}

func ExecRunFunc(dir string, env []string, name string, args1 ...string,
) func(arg2 ...string) error {
	return func(args2 ...string) error {
		Out(cmdString(name, slices.Concat(args1, args2)...))
		cmd := exec.Command(name, append(args1, args2...)...)
		cmd.Dir = dir
		cmd.Env = env
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		return nil
	}
}

// Return a string that is pastable directly in a shell (that is, with proper
// quotes when arguments contain spaces).
func cmdString(name string, args ...string) string {
	var bld strings.Builder
	fmt.Fprint(&bld, name)
	for _, arg := range args {
		if strings.Contains(arg, " ") {
			fmt.Fprintf(&bld, " %q", arg)
		} else {
			fmt.Fprintf(&bld, " %s", arg)
		}
	}
	return bld.String()
}
