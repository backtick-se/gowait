package daemon

import (
	"github.com/backtick-se/gowait/core"
	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"

	"context"

	"go.uber.org/zap"
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
	log     *zap.Logger
}

func NewWorkers(driver cluster.Driver, queue task.TaskQueue, log *zap.Logger) Workers {
	return &workers{
		driver:  driver,
		queue:   queue,
		byId:    make(map[task.ID]Worker),
		byImage: make(map[string][]Worker),
		log:     log,
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

	w.log.Info("removed executor", zap.String("executor", string(id)))

	// todo: handle not found errors
	return w.driver.Kill(ctx, id)
}

func (w *workers) Request(ctx context.Context, image string) (Worker, error) {
	log := w.log.With(zap.String("image", image))
	log.Info("requested executor")

	workers, ok := w.byImage[image]
	if !ok {
		workers = []Worker{}
	}

	for _, worker := range workers {
		if worker.Image() == image && worker.Status() == executor.StatusIdle {
			log.Info("found existing executor", zap.String("executor", string(worker.ID())))
			return worker, nil
		}
	}

	// create new executor
	worker, err := w.spawn(ctx, image)
	if err != nil {
		return nil, err
	}

	w.log.Info("spawned new executor", zap.String("executor", string(worker.ID())))
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
	worker := NewWorker(w.driver, id, image, w.log)

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

		worker.OnAquire(instance)
		return instance.State(), nil
	}
	return nil, core.ErrUnknownTask
}

func (t *workers) ExecStop(ctx context.Context, req *executor.MsgStop) error {
	id := task.ID(req.Header.ID)
	return t.Remove(ctx, id)
}
