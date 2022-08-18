package executor

import (
	"io"
)

type LineWriter interface {
	Write(line string) error
	Close() error
	Wait() error
}

type linewriter struct {
	output  io.WriteCloser
	input   chan string
	failure chan error
	closed  bool
}

func NewLineWriter(output io.WriteCloser) LineWriter {
	lw := &linewriter{
		output:  output,
		input:   make(chan string),
		failure: make(chan error),
		closed:  false,
	}
	go lw.proc()
	return lw
}

func (w *linewriter) Write(line string) error {
	if w.closed {
		return io.EOF
	}
	w.input <- line
	err := <-w.failure
	return err
}

func (w *linewriter) Wait() error {
	if w.closed {
		return nil
	}
	return <-w.failure
}

func (w *linewriter) Close() error {
	return nil
}

func (w *linewriter) proc() {
	defer close(w.input)
	defer close(w.failure)

	for input := range w.input {
		bytes := []byte(input)
		sent := 0
		for sent < len(bytes) {
			n, err := w.output.Write(bytes)
			sent += n
			if err != nil {
				w.failure <- err
				return
			}
		}
		w.failure <- nil
	}
}
