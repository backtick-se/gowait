package client

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Process struct {
	Upstream   LineReader
	Stdout     LineReader
	Stderr     LineReader
	Downstream LineWriter

	command *exec.Cmd
}

func Exec(command string, args ...string) (*Process, error) {
	cmd := exec.Command(command, args...)

	var err error
	var stdout, stderr io.ReadCloser
	var upstream, upstreamWriter *os.File
	var downstream, downstreamReader *os.File

	cleanup := func() {
		if stdout != nil {
			stdout.Close()
		}
		if stderr != nil {
			stderr.Close()
		}
		if upstream != nil {
			upstream.Close()
			upstreamWriter.Close()
		}
		if downstream != nil {
			downstream.Close()
			downstreamReader.Close()
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

	// create upstream pipe
	if upstream, upstreamWriter, err = os.Pipe(); err != nil {
		cleanup()
		return nil, err
	}

	// create downstream pipe
	if downstreamReader, downstream, err = os.Pipe(); err != nil {
		cleanup()
		return nil, err
	}

	// pass extra file handles to the command
	cmd.ExtraFiles = []*os.File{
		upstreamWriter,
		downstreamReader,
	}

	// start subprocess
	if err := cmd.Start(); err != nil {
		cleanup()
		return nil, err
	}

	// close handles that were handed over to the subprocess
	if err := upstreamWriter.Close(); err != nil {
		fmt.Printf("failed to close upstream writer: %s\n", err)
	}
	if err := downstreamReader.Close(); err != nil {
		fmt.Printf("failed to close downstream reader: %s\n", err)
	}

	return &Process{
		command: cmd,

		Stdout:     NewLineReader(stdout),
		Stderr:     NewLineReader(stderr),
		Upstream:   NewLineReader(upstream),
		Downstream: NewLineWriter(downstream),
	}, nil
}

func (p *Process) Wait() error {
	defer p.Upstream.Close()
	defer p.Downstream.Close()

	// wait for pipes
	if err := p.Stdout.Wait(); err != nil {
		fmt.Println("error reading stdout:", err)
	}
	if err := p.Stderr.Wait(); err != nil {
		fmt.Println("error reading stderr:", err)
	}

	return p.command.Wait()
}
