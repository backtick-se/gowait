package executor

import (
	"cowait/core"
	"cowait/core/client"
	"os"

	"context"
	"fmt"
	"time"
)

type Executor interface {
	Run(context.Context) error
}

type executor struct {
	client client.TaskClient
	task   *core.TaskDef
}

func NewFromEnv(client client.TaskClient) (Executor, error) {
	id := core.TaskID(os.Getenv(core.EnvTaskID))
	def, err := core.TaskDefFromEnv(os.Getenv(core.EnvTaskdef))
	if err != nil {
		return nil, err
	}

	fmt.Printf("executor %s: %+v\n", id, def)

	// connect to daemon
	hostname := "cowaitd.default.svc.cluster.local:1337"
	if err := client.Connect(hostname, id); err != nil {
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

	server, err := newServer()
	if err != nil {
		e.client.Failure(ctx, err)
		return err
	}
	defer server.Close()

	if err := e.client.Init(ctx); err != nil {
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

	exited := make(chan error)
	defer close(exited)

	go func() {
		exited <- proc.Wait()
	}()

	select {
	case req := <-server.completed:
		logger.Close()
		e.client.Complete(ctx, string(req.Result))
		break

	case req := <-server.failed:
		logger.Close()
		e.client.Failure(ctx, req.Error)
		break

	case req := <-server.log:
		logger.Log(req.File, req.Data)
		break

	case err := <-exited:
		logger.Close()
		if err == nil {
			// completed without any result?
			e.client.Complete(ctx, "{}")
		} else {
			e.client.Failure(ctx, err)
		}
		break

	case <-ctx.Done():
		e.client.Failure(ctx, fmt.Errorf("deadline exceeded"))
		break
	}

	return nil
}
