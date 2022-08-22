package executor

import (
	"cowait/core"

	"bytes"
	"fmt"
	"io"
	"os"
)

type LineReader interface {
	Read() (string, error)
	Wait() error
	Close() error
}

type linereader struct {
	input   io.ReadCloser
	output  chan string
	failure chan error
	closed  bool
}

func NewLineReader(input io.ReadCloser) LineReader {
	lp := &linereader{
		input:   input,
		output:  make(chan string),
		failure: make(chan error),
		closed:  false,
	}
	go lp.proc()
	return lp
}

func (p *linereader) Wait() error {
	if p.closed {
		return nil
	}
	err := <-p.failure
	if err != io.EOF {
		return err
	}
	return nil
}

func (p *linereader) Close() error {
	if p.closed {
		return nil
	}
	if err := p.input.Close(); err != nil {
		return err
	}
	return p.Wait()
}

func (p *linereader) Read() (string, error) {
	if p.closed {
		return "", io.EOF
	}
	select {
	case line := <-p.output:
		return line, nil
	case err := <-p.failure:
		p.closed = true
		return "", err
	}
}

func (p *linereader) proc() {
	defer close(p.output)
	defer close(p.failure)

	buffer := make([]byte, 10240)
	offset := 0
	for {
		n, err := p.input.Read(buffer[offset:])
		pos := 0
		end := n + offset
		for pos < end {
			nl := bytes.IndexByte(buffer[pos:], '\n')
			if nl >= 0 && pos+nl < end {
				line := string(buffer[pos : pos+nl+1])
				p.output <- line
				pos += nl + 1
			} else {
				if pos > 0 {
					copy(buffer, buffer[pos:end])
				}
				break
			}
		}
		offset = end - pos

		if err == io.EOF {
			// if we have data in the buffer, submit it.
			if offset > 0 {
				p.output <- string(buffer[:offset])
			}
			p.failure <- io.EOF
			return
		}
		if err != nil {
			p.failure <- err
			return
		}
	}
}

func LinePrinter(name string, pump LineReader) {
	prefix := []byte(name + " | ")
	for {
		line, err := pump.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(name, "read error:", err)
			}
			break
		}
		if len(line) > 0 {
			os.Stdout.Write(prefix)
			os.Stdout.Write([]byte(line))
		}
	}
}

func LineLogger(name string, pump LineReader, log core.Logger) {
	for {
		line, err := pump.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(name, "read error:", err)
			}
			break
		}
		if len(line) > 0 {
			log.Log(name, line)
		}
	}
}
