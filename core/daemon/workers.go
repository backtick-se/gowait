package daemon

import (
	"context"
	"fmt"

	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
)

type Workers interface {
	executor.Handler

	Remove(context.Context, task.ID) error
	Request(ctx context.Context, image string) (Worker, error)
}

type workers struct {
	driver  cluster.Driver
	queue   task.TaskQueue
	byId    map[task.ID]Worker
	byImage map[string][]Worker
}

func NewWorkers(driver cluster.Driver, queue task.TaskQueue) Workers {
	return &workers{
		driver:  driver,
		queue:   queue,
		byId:    make(map[task.ID]Worker),
		byImage: make(map[string][]Worker),
	}
}

func (w *workers) Remove(ctx context.Context, id task.ID) error {
	worker, ok := w.get(id)
	if !ok {
		return core.ErrUnknownTask
	}

	delete(w.byId, id)
	worker.OnStop()

	// we know that this will never fail
	workers := w.byImage[worker.Image()]
	for idx, wi := range workers {
		if wi == worker {
			w.byImage[worker.Image()] = append(workers[:idx], workers[idx+1:]...)
		}
	}

	// todo: handle not found errors
	return w.driver.Kill(ctx, id)
}

func (w *workers) Request(ctx context.Context, image string) (Worker, error) {
	fmt.Println("requested executor for", image)
	workers, ok := w.byImage[image]
	if !ok {
		workers = []Worker{}
	}

	for _, worker := range workers {
		if worker.Image() == image && worker.Status() == executor.StatusIdle {
			fmt.Println("found existing executor", worker.ID())
			return worker, nil
		}
	}

	worker, err := w.spawn(ctx, image)
	if err != nil {
		return nil, err
	}
	fmt.Println("spawned new executor", worker.ID())
	workers = append(workers, worker)
	w.byImage[image] = workers
	return worker, nil
}

func (w *workers) get(id task.ID) (Worker, bool) {
	wr, ok := w.byId[id]
	return wr, ok
}

func (w *workers) spawn(ctx context.Context, image string) (Worker, error) {
	id := task.GenerateID("executor")
	worker := NewWorker(w.driver, id, image)

	if err := worker.Start(ctx); err != nil {
		return nil, err
	}

	workers, ok := w.byImage[worker.Image()]
	if !ok {
		workers = []Worker{}
	}
	w.byImage[worker.Image()] = append(workers, worker)
	w.byId[worker.ID()] = worker

	return worker, nil
}

//
// Executor Server implementation
//

func (t *workers) ExecInit(ctx context.Context, req *executor.MsgInit) error {
	id := task.ID(req.Header.ID)
	if worker, ok := t.get(id); ok {
		fmt.Println("executor init:", id)
		worker.OnInit()
	}
	return nil
}

func (t *workers) ExecAquire(ctx context.Context, req *executor.MsgAquire) (*task.Run, error) {
	id := task.ID(req.Header.ID)
	if worker, ok := t.get(id); ok {
		// find the next suitable work item for this executor
		// this will block until a new task is available.
		// its up to the caller to abort the call if the wait is too long.
		// this greatly reduces task startup latency
		instance, err := t.queue.Aquire(ctx, worker.Image())
		if err != nil {
			return nil, err
		}
		// todo: keep track of which instance is running what task
		// instance.Assign(worker)

		fmt.Println("executor/aquire", *instance.State())
		worker.OnAquire(instance)
		return instance.State(), nil
	}
	return nil, core.ErrUnknownTask
}

func (t *workers) ExecStop(ctx context.Context, req *executor.MsgStop) error {
	id := task.ID(req.Header.ID)
	return t.Remove(ctx, id)
}
