package vis

import (
	"fmt"
	"os"
	"os/exec"
)

// https://pkg.go.dev/cmd/link
// -s	Omit the symbol table and debug information.
// -w	Omit the DWARF symbol table.
// FIXME if I pass ldflags as quoted, go build returns an error.
// If I don't, "go build" is happy, but the log output misses the quotes,
// so it cannot be copy and pasted :-(
// ldflags := "-w -s"

// GoBuild invokes "go build args...".
func GoBuild(goos string, args ...string) error {
	env := os.Environ()
	if goos != "" {
		env = append(env, "GOOS="+goos)
	}
	args2 := append([]string{"build"}, args...)
	stdout, stderr, err := ExecOutput(env, "go", args2...)
	if stdout != "" {
		fmt.Println("✅", stdout)
	}
	if stderr != "" {
		fmt.Println("❌", stderr)
	}
	return err
}

// GoTest invokes "go test args...".
func GoTest(args ...string) error {
	dir := "."
	env := os.Environ()
	args2 := append([]string{"test"}, args...)
	return ExecRun(dir, env, "go", args2...)
}

// GoTestSum invokes "gotestsum args..." if utility gotestsum is found,
// otherwhise it behaves like GoTest.
//
// To install gotestsum, run: go install gotest.tools/gotestsum@latest
func GoTestSum(args ...string) error {
	if _, err := exec.LookPath("gotestsum"); err != nil {
		return GoTest(args...)
	}
	dir := "."
	env := os.Environ()
	args2 := append([]string{"--"}, args...)
	return ExecRun(dir, env, "gotestsum", args2...)
}

func GoCoverageBrowser(coverfile string) {
	dir := "."
	env := os.Environ()
	ExecRun(dir, env, "go", "tool", "cover", "-html", coverfile)
}
