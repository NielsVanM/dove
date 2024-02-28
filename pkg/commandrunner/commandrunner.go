package commandrunner

import (
	"io"
	"os/exec"
)

type CommandRunner interface {
	Run(identifier string, command string, args ...string)
	Stop(identifier string)

	GetReader(identifier string) io.Reader
	GetWriter(identifier string) io.Writer
}

type CMDCommandRunner struct {
	commands map[string]*exec.Cmd
}

func NewCMDCommandRunner() *CMDCommandRunner {
	return &CMDCommandRunner{
		make(map[string]*exec.Cmd),
	}
}

func (ccr *CMDCommandRunner) Run(identifier string, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Start()
	ccr.commands[identifier] = cmd
}

func (ccr *CMDCommandRunner) Stop(identifier string) error {
	cmd := ccr.getCmd(identifier)
	err := cmd.Process.Kill()
  if err != nil {
    return err
  }
  
  return nil
}

func (ccr *CMDCommandRunner) getCmd(identifier string) *exec.Cmd {
	return ccr.commands[identifier]
}
