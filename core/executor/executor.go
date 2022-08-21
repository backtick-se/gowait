package executor

import (
	"cowait/core"
	"cowait/core/client"

	"context"
	"fmt"
	"time"
)

type Executor interface {
	Run(context.Context, core.TaskID, *core.TaskSpec) error
}

type executor struct {
	server Server
	client client.Executor
}

func New(client client.Executor, server Server) (Executor, error) {
	return &executor{
		server: server,
		client: client,
	}, nil
}

func (e *executor) Run(ctx context.Context, id core.TaskID, task *core.TaskSpec) error {
	// apply timeout if set
	if task.Timeout > 0 {
		deadline, cancel := context.WithTimeout(ctx, time.Duration(task.Timeout)*time.Second)
		defer cancel()
		ctx = deadline
	}

	// connect to daemon
	hostname := "cowaitd.default.svc.cluster.local:1337"
	if err := e.client.Connect(hostname, id); err != nil {
		return fmt.Errorf("failed to connect upstream: %w", err)
	}

	// start sdk rpc server
	if err := e.server.Listen(1337); err != nil {
		e.client.Failure(ctx, err)
		return err
	}

	if err := e.client.Init(ctx); err != nil {
		// init failed
		return err
	}

	logger, err := e.client.Log(ctx)
	if err != nil {
		// logging failed
		return err
	}

	proc, err := Exec(task.Command[0], task.Command[1:]...)
	if err != nil {
		logger.Close()
		e.client.Failure(ctx, err)
		return err
	}

	go LineLogger("stdout", proc.Stdout(), logger)
	go LineLogger("stderr", proc.Stderr(), logger)

	select {
	case req := <-e.server.Completed():
		logger.Close()
		e.client.Complete(ctx, string(req.Result))
		break

	case req := <-e.server.Failed():
		logger.Close()
		e.client.Failure(ctx, req.Error)
		break

	case req := <-e.server.Logs():
		logger.Log(req.File, req.Data)
		break

	case err := <-proc.Done():
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
