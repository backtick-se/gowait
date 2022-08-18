package executor

import (
	"fmt"
	"io"
	"os/exec"
)

type Process struct {
	Stdout LineReader
	Stderr LineReader

	command *exec.Cmd
}

func Exec(command string, args ...string) (*Process, error) {
	cmd := exec.Command(command, args...)

	var err error
	var stdout, stderr io.ReadCloser

	cleanup := func() {
		if stdout != nil {
			stdout.Close()
		}
		if stderr != nil {
			stderr.Close()
		}
	}

	// create stdout/stderr pipes
	if stdout, err = cmd.StdoutPipe(); err != nil {
		cleanup()
		return nil, err
	}
	if stderr, err = cmd.StderrPipe(); err != nil {
		cleanup()
		return nil, err
	}

	// start subprocess
	if err := cmd.Start(); err != nil {
		cleanup()
		return nil, err
	}

	return &Process{
		command: cmd,

		Stdout: NewLineReader(stdout),
		Stderr: NewLineReader(stderr),
	}, nil
}

func (p *Process) Wait() error {
	// wait for pipes
	if err := p.Stdout.Wait(); err != nil {
		fmt.Println("error reading stdout:", err)
	}
	if err := p.Stderr.Wait(); err != nil {
		fmt.Println("error reading stderr:", err)
	}

	return p.command.Wait()
}
