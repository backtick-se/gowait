package executor

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
	"fmt"
	"time"
)

type T interface {
	Run(context.Context, task.ID) error
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

func (e *executor) Run(ctx context.Context, id task.ID) error {
	// connect to daemon
	hostname := "cowaitd.default.svc.cluster.local:1337"
	if err := e.client.Connect(hostname, id); err != nil {
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
	timeout := time.After(time.Minute)
	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("aquire timeout")
		default:
		}

		spec, err := e.client.ExecAquire(context.Background())
		if err != nil {
			// todo: differentiate between different errors
			// we should only continue on "no task available" errors
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("aquired task:", spec)
		return spec, nil
	}
}

func (e *executor) exec(spec *task.Run) (task.Result, error) {
	ctx := context.Background()

	fmt.Println("sending init..")
	if err := e.client.Init(ctx, spec.ID); err != nil {
		// init failed
		return nil, err
	}

	fmt.Println("open log init..")
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

	fmt.Println("run exe")
	proc, err := Exec(spec.Command[0], spec.Command[1:]...)
	if err != nil {
		return nil, err
	}

	go LineLogger("stdout", proc.Stdout(), logger)
	go LineLogger("stderr", proc.Stderr(), logger)

	select {
	case req := <-e.server.OnInit():
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
	case req := <-e.server.OnComplete():
		return req.Result, nil

	case req := <-e.server.OnFailure():
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
