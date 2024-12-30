package vis

import (
	"bytes"
	"os/exec"
)

// v0.7.0
func GitShortVersion() (string, error) {
	out, err := exec.Command("git", "describe",
		"--tags", "--always", "--abbrev=0").Output()
	return string(bytes.TrimSpace(out)), err
}

// v0.7.0-18-g92e5e86d65
func GitLongVersion() (string, error) {
	out, err := exec.Command("git", "describe",
		"--tags", "--always", "--dirty").Output()
	return string(bytes.TrimSpace(out)), err
}
