package cmd

import (
	"os"
	"os/exec"
	"testing"
)

func TestNewCommand(t *testing.T) {
	err := NewCommand(
		"echo hello",
		WithCustomBaseCommand(exec.Command(`C:\windows\system32\cmd.exe`, "/C")),
		WithStandardStreams,
		WithCustomStdout(os.Stdout),
		WithCustomStderr(os.Stderr),
		WithTimeout(10),
		WithoutTimeout,
		WithWorkingDir("."),
		WithInheritedEnvironment(EnvVars{
			"PATH": os.Getenv("PATH"),
		}),
	).Execute()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
