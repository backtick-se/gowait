package client

import (
	"cowait/core"
	"encoding/json"
	"io"

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
	fmt.Printf("running task: %s$%s\n", e.task.Image, e.task.Command)

	proc, err := Exec(e.task.Command[0], e.task.Command[1:]...)
	if err != nil {
		return err
	}

	go LinePrinter("stdout", proc.Stdout)
	go LinePrinter("stderr", proc.Stderr)

	go UpstreamHandler(proc.Upstream, &DownstreamWriter{proc.Downstream})

	return proc.Wait()
}

func UpstreamHandler(upstream LineReader, downstream *DownstreamWriter) {
	for {
		line, err := upstream.Read()
		if err == io.EOF {
			break
		}

		// unpack command
		msg := MsgHeader{}
		err = json.Unmarshal([]byte(line), &msg)
		if err != nil {
			fmt.Println("invalid upstream msg:", err)
			continue
		}

		switch msg.Type {
		case MsgInvoke:
			invoke := core.TaskDef{}
			err := json.Unmarshal(msg.Body, &invoke)
			if err != nil {
				fmt.Println("failed to parse invoke:", err)
				continue
			}
			fmt.Println("INVOKE:", invoke.Name)
			if err := downstream.Result(map[string]any{
				"result": 10,
			}); err != nil {
				fmt.Println("failed to write result:", err)
			}

		case MsgResult:
			fmt.Println("RESULT:", string(msg.Body))

		default:
			fmt.Println("unknown upstream msg:", msg.Type)
		}
	}
}

type MsgHeader struct {
	ID   uint64
	Type MsgType
	Body json.RawMessage
}

type MsgType string

const (
	MsgInvoke MsgType = "cowait/invoke"
	MsgResult MsgType = "cowait/result"
	MsgError  MsgType = "cowait/error"
)

type DownstreamWriter struct {
	w LineWriter
}

func (d *DownstreamWriter) Result(r any) error {
	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	msg := MsgHeader{
		Type: MsgResult,
		Body: body,
	}
	msgjson, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	msgjson = append(msgjson, '\n')

	return d.w.Write(string(msgjson))
}
