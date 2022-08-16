package client

import (
	"bytes"
	"cowait/core"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"fmt"
)

type Executor interface {
	Run() error
}

type executor struct {
	task *core.TaskDef
}

func ExecutorFromEnv(envdef string) (Executor, error) {
	def, err := core.TaskDefFromEnv(envdef)
	if err != nil {
		return nil, err
	}

	return &executor{
		task: def,
	}, nil
}

func (e *executor) Run() error {
	fmt.Printf("11 running task: %+v\n", e.task)

	cmd := exec.Command(e.task.Command[0], e.task.Command[1:]...)

	upstream, upstreamWriter, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer upstream.Close()

	downstreamReader, downstream, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer downstream.Close()

	cmd.ExtraFiles = []*os.File{
		upstreamWriter,
		downstreamReader,
	}

	var stdout, stderr io.ReadCloser

	if stdout, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}
	if stderr, err = cmd.StderrPipe(); err != nil {
		log.Fatal(err)
	}

	stdoutDone := make(chan error)
	stderrDone := make(chan error)
	go ReaderFunc(stdout, func(s string) { print(s) }, stdoutDone)
	go ReaderFunc(stderr, func(s string) { print(s) }, stderrDone)

	// start subprocess
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// close handles that were handed over to the subprocess
	if err := upstreamWriter.Close(); err != nil {
		fmt.Println("warning: error closing upstream writer:", err)
	}
	if err := downstreamReader.Close(); err != nil {
		fmt.Println("warning: error closing downstream reader:", err)
	}

	downstream.Write([]byte("{\"result\": 10}\n"))

	upstreamOut, err := ioutil.ReadAll(upstream)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("upstream output:")
	fmt.Println(string(upstreamOut))

	// wait for pipes
	if err := <-stdoutDone; err != nil {
		log.Fatal("error reading stdout:", err)
	}
	if err := <-stderrDone; err != nil {
		log.Fatal("error reading stderr:", err)
	}

	return cmd.Wait()
}

func ReaderFunc(input io.ReadCloser, fn func(string), done chan<- error) {
	output := make(chan string)
	go func() {
		for line := range output {
			fn(line)
		}
	}()
	done <- ReaderPump(input, output)
	close(done)
}

func ReaderPump(input io.ReadCloser, output chan<- string) error {
	defer close(output)
	buffer := make([]byte, 10240)
	offset := 0
	for {
		n, err := input.Read(buffer[offset:])
		off := 0
		for off < n {
			nl := bytes.IndexByte(buffer[off:], '\n')
			if nl >= 0 {
				line := string(buffer[off : off+nl+1])
				output <- line
				off += nl + 1
				offset = 0
			} else {
				copy(buffer, buffer[off:n])
				offset = n - off
				break
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
