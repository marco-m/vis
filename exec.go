package vis

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecOutput(env []string, name string, arg ...string) (string, string, error) {
	// We don't use cmd.String() because we don't want the full path of 'name'.
	Out(strings.Join(append([]string{name}, arg...), " "))
	cmd := exec.Command(name, arg...)
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

func ExecRun(dir string, env []string, name string, arg ...string) error {
	// We don't use cmd.String() because we don't want the full path of 'name'.
	Out(strings.Join(append([]string{name}, arg...), " "))
	cmd := exec.Command(name, arg...)
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

func ExecRunFunc(dir string, env []string, name string, arg1 ...string,
) func(arg2 ...string) error {
	return func(arg2 ...string) error {
		cmd := exec.Command(name, append(arg1, arg2...)...)
		Out("cmd", cmd.String(), "dir", dir)
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
