package commandrunner

import (
	"os/exec"
	"testing"
)

func TestRunCommand(t *testing.T) {
	command := "sleep"
	args := "10"
	identifier := "sleep"

	cmdr := NewCMDCommandRunner()
	cmdr.Run(identifier, command, args)

	if len(cmdr.commands) != 1 {
		t.Errorf("Expected one running process, got, %d", len(cmdr.commands))
		return
	}

	cmdr.commands[identifier].Process.Kill()
}

func TestStopCommand(t *testing.T) {
	command := "sleep"
	args := "10"
	identifier := "sleep"

	cmd := exec.Command(command, args)
	cmdr := NewCMDCommandRunner()
	cmdr.commands[identifier] = cmd
	cmd.Start()

	err := cmdr.Stop(identifier)

	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

  state, err := cmd.Process.Wait()

  if state.String() != "signal: killed" {
    t.Errorf("Expected process to be killed but got %s", state.String())
    return
  }
}
