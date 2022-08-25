package executor

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

type T interface {
	Run(context.Context, task.ID, *task.Spec) error
}

type executor struct {
	server Server
	client Client
}

func New(client Client, server Server) (T, error) {
	return &executor{
		server: server,
		client: client,
	}, nil
}

func (e *executor) Run(ctx context.Context, id task.ID, task *task.Spec) error {
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
	case req := <-e.server.OnInit():
		// sdk initialization
		logger.Log("system", fmt.Sprintf("sdk init: %s\n", req.Version))
		break

	case err := <-proc.Done():
		logger.Log("system", fmt.Sprintf("process exit without sdk init: %s\n", err.Error()))
		logger.Close()
		if err == nil {
			// if the task exits with code 0 without sending init, we assume its running without SDK
			// so we consider it a successful completion with an empty result
			e.client.Complete(ctx, "{}")
			return nil
		} else {
			e.client.Failure(ctx, err)
			return err
		}

	case <-ctx.Done():
		logger.Close()
		err := fmt.Errorf("deadline exceeded")
		e.client.Failure(ctx, err)
		return err
	}

	select {
	case req := <-e.server.OnComplete():
		logger.Close()
		e.client.Complete(ctx, string(req.Result))
		break

	case req := <-e.server.OnFailure():
		logger.Close()
		e.client.Failure(ctx, req.Error)
		break

	case err := <-proc.Done():
		// exiting without sending complete/fail is an error
		logger.Close()
		if err == nil {
			e.client.Failure(ctx, fmt.Errorf("task exited unexpectedly with status 0"))
		} else {
			e.client.Failure(ctx, fmt.Errorf("task exited unexpectedly: %w", err))
		}
		break

	case <-ctx.Done():
		logger.Close()
		e.client.Failure(ctx, fmt.Errorf("deadline exceeded"))
		break
	}

	return nil
}
