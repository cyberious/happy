package artifact_builder

import "os/exec"

type Executor interface {
	Run(command *exec.Cmd) error
	Output(command *exec.Cmd) ([]byte, error)
}

type DefaultExecutor struct{}

func (e DefaultExecutor) Run(command *exec.Cmd) error {
	return command.Run()
}

func (e DefaultExecutor) Output(command *exec.Cmd) ([]byte, error) {
	return command.Output()
}

func NewDefaultExecutor() Executor {
	return DefaultExecutor{}
}

type DummyExecutor struct{}

func (e DummyExecutor) Run(command *exec.Cmd) error {
	return nil
}

func (e DummyExecutor) Output(command *exec.Cmd) ([]byte, error) {
	return []byte{}, nil
}

func NewDummyExecutor() Executor {
	return DummyExecutor{}
}