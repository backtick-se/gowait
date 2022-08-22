package executor

import (
	"fmt"
	"io"
	"os/exec"
)

type Process interface {
	Stdout() LineReader
	Stderr() LineReader
	Done() <-chan error
}

type process struct {
	stdout LineReader
	stderr LineReader

	command *exec.Cmd
	done    chan error
}

func Exec(command string, args ...string) (Process, error) {
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

	proc := &process{
		command: cmd,
		done:    make(chan error),
		stdout:  NewLineReader(stdout),
		stderr:  NewLineReader(stderr),
	}

	go func() {
		proc.done <- proc.wait()
		close(proc.done)
	}()

	return proc, nil
}

func (p *process) Stdout() LineReader { return p.stdout }
func (p *process) Stderr() LineReader { return p.stderr }
func (p *process) Done() <-chan error { return p.done }

func (p *process) wait() error {
	// wait for pipes
	if err := p.stdout.Wait(); err != nil {
		fmt.Println("error reading stdout:", err)
	}
	if err := p.stderr.Wait(); err != nil {
		fmt.Println("error reading stderr:", err)
	}

	return p.command.Wait()
}
