package executor

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

type ID string

type T interface {
	Run(context.Context, task.ID, string) error
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

func (e *executor) Run(ctx context.Context, id task.ID, daemonEndpoint string) error {
	// connect to daemon
	if err := e.client.Connect(daemonEndpoint, id); err != nil {
		return fmt.Errorf("failed to connect upstream: %w", err)
	}

	// register executor
	if err := e.client.ExecInit(ctx, nil); err != nil {
		return fmt.Errorf("failed to register: %w", err)
	}

	for {
		// grab next task
		spec, err := e.aquire()
		if err != nil {
			fmt.Println("aquire failed. stopping.")
			return e.client.ExecStop(ctx)
		}

		// store a copy in the server so that it may respond to SDK Aquire()
		e.server.SetRun(spec)

		// apply timeout if set
		if spec.Timeout > 0 {
			deadline, cancel := context.WithTimeout(ctx, time.Duration(spec.Timeout)*time.Second)
			defer cancel()
			ctx = deadline
		}

		// execute task
		result, err := e.exec(spec)

		if err != nil {
			fmt.Println("error:", err)
			e.client.Failure(ctx, spec.ID, err)
		} else {
			fmt.Println("complete:", string(result))
			e.client.Complete(ctx, spec.ID, string(result))
		}
	}
}

func (e *executor) aquire() (*task.Run, error) {
	fmt.Println("waiting...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	spec, err := e.client.ExecAquire(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("aquired task:", spec)
	return spec, nil
}

func (e *executor) exec(spec *task.Run) (task.Result, error) {
	ctx := context.Background()

	if err := e.client.Init(ctx, spec.ID); err != nil {
		// init failed
		return nil, err
	}

	logger, err := e.client.Log(ctx, spec.ID)
	if err != nil {
		// logging failed
		return nil, err
	}
	defer logger.Close()

	// apply timeout if set
	if spec.Timeout > 0 {
		deadline, cancel := context.WithTimeout(ctx, time.Duration(spec.Timeout)*time.Second)
		defer cancel()
		ctx = deadline
	}

	logger.Log("system", fmt.Sprintf("running command %s", spec.Command))
	proc, err := Exec(spec.Command[0], spec.Command[1:]...)
	if err != nil {
		return nil, err
	}

	go LineLogger("stdout", proc.Stdout(), logger)
	go LineLogger("stderr", proc.Stderr(), logger)

	select {
	case req := <-e.server.AwaitInit():
		// sdk initialization
		logger.Log("system", fmt.Sprintf("sdk init: %s\n", req.Version))
		break

	case err := <-proc.Done():
		if err == nil {
			// if the task exits with code 0 without sending init, we assume its running without SDK
			// so we consider it a successful completion with an empty result
			logger.Log("system", "process exit ok without sdk init")
			return task.NoResult, nil
		} else {
			logger.Log("system", fmt.Sprintf("process exit without sdk init: %s\n", err.Error()))
			return nil, err
		}

	case <-ctx.Done():
		return nil, fmt.Errorf("deadline exceeded")
	}

	select {
	case req := <-e.server.AwaitComplete():
		return req.Result, nil

	case req := <-e.server.AwaitFailure():
		return nil, req.Error

	case err := <-proc.Done():
		// exiting without sending complete/fail is an error
		if err == nil {
			return nil, fmt.Errorf("task exited unexpectedly with status 0")
		} else {
			return nil, fmt.Errorf("task exited unexpectedly: %w", err)
		}

	case <-ctx.Done():
		return nil, fmt.Errorf("deadline exceeded")
	}
}
