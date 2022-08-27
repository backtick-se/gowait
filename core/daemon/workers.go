package daemon

import (
	"fmt"

	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
)

type Workers interface {
	Add(Worker)
	Remove(Worker)
	Get(id task.ID) (Worker, bool)
	Request(image string) Worker
}

type workers struct {
	driver  cluster.Driver
	byId    map[task.ID]Worker
	byImage map[string][]Worker
}

func NewWorkers(driver cluster.Driver) Workers {
	return &workers{
		driver:  driver,
		byId:    make(map[task.ID]Worker),
		byImage: make(map[string][]Worker),
	}
}

func (w *workers) Get(id task.ID) (Worker, bool) {
	wr, ok := w.byId[id]
	return wr, ok
}

func (w *workers) Add(worker Worker) {
	workers, ok := w.byImage[worker.Image()]
	if !ok {
		workers = []Worker{}
	}
	workers = append(workers, worker)
	w.byImage[worker.Image()] = workers
	w.byId[worker.ID()] = worker
}

func (w *workers) Remove(worker Worker) {
	workers, ok := w.byImage[worker.Image()]
	if ok {
		for idx, wi := range workers {
			if wi == worker {
				w.byImage[worker.Image()] = append(workers[:idx], workers[idx+1:]...)
				delete(w.byId, worker.ID())
				return
			}
		}
	}
}

func (w *workers) Request(image string) Worker {
	fmt.Println("requested executor for", image)
	workers, ok := w.byImage[image]
	if !ok {
		workers = []Worker{}
	}

	for _, worker := range workers {
		if worker.Image() == image && worker.Status() == executor.StatusIdle {
			fmt.Println("found existing executor", worker.ID())
			return worker
		}
	}

	worker := w.Spawn(image)
	fmt.Println("spawning new executor", worker.ID())
	workers = append(workers, worker)
	w.byImage[image] = workers
	return worker
}

func (w *workers) Spawn(image string) Worker {
	id := task.GenerateID("executor")
	worker := NewWorker(w.driver, id, image)
	w.Add(worker)
	return worker
}