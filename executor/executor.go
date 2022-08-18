package executor

import (
	"context"
	"cowait/core"
	"cowait/core/client"
	"encoding/json"
	"io"
	"time"

	"fmt"
)

type Executor interface {
	Run(context.Context) error
}

type executor struct {
	client client.Client
	task   *core.TaskDef
}

func NewFromEnv(envdef string) (Executor, error) {
	def, err := core.TaskDefFromEnv(envdef)
	if err != nil {
		return nil, err
	}

	client, err := client.New(def.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to connect upstream: %w", err)
	}

	return &executor{
		client: client,
		task:   def,
	}, nil
}

func (e *executor) Run(ctx context.Context) error {
	fmt.Printf("running task: %s$%s\n", e.task.Image, e.task.Command)

	// apply timeout if set
	if e.task.Timeout > 0 {
		deadline, cancel := context.WithTimeout(ctx, time.Duration(e.task.Timeout)*time.Second)
		defer cancel()
		ctx = deadline
	}

	if err := e.client.Init(ctx, e.task); err != nil {
		// init failed
		return err
	}

	logger, err := e.client.Log(ctx)
	if err != nil {
		// logging failed
		return err
	}

	proc, err := Exec(e.task.Command[0], e.task.Command[1:]...)
	if err != nil {
		return err
	}

	go LineLogger("stdout", proc.Stdout, logger)
	go LineLogger("stderr", proc.Stderr, logger)

	result := "{}"
	{
		var err error
		done := make(chan string)
		fail := make(chan error)
		complete := make(chan error)
		defer close(done)
		defer close(fail)
		go UpstreamHandler(proc.Upstream, &DownstreamWriter{proc.Downstream}, done, fail)

		go func() {
			complete <- proc.Wait()
		}()

		select {
		case result = <-done:
			break
		case err = <-fail:
			break
		case err = <-complete:
			break
		case <-ctx.Done():
			err = fmt.Errorf("deadline exceeded")
			break
		}

		if err != nil {
			// task error
			e.client.Failure(ctx, err)
			return err
		}
	}

	logger.Close()

	return e.client.Complete(ctx, result)
}

func UpstreamHandler(upstream LineReader, downstream *DownstreamWriter, done chan string, fail chan error) {
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
			done <- string(msg.Body)
			return

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
