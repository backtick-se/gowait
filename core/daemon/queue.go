package daemon

import "github.com/backtick-se/gowait/core/task"

type Queue interface {
	Push(*task.Spec) Instance
	Pop(string) Instance
	Get(task.ID) (Instance, bool)
}

type queue struct {
	byId  map[task.ID]Instance
	queue []Instance
}

func NewQueue() Queue {
	return &queue{
		queue: make([]Instance, 0, 32),
		byId:  make(map[task.ID]Instance),
	}
}

func (q *queue) Get(id task.ID) (Instance, bool) {
	i, ok := q.byId[id]
	return i, ok
}

func (q *queue) Push(spec *task.Spec) Instance {
	instance := newInstance(spec)
	q.queue = append(q.queue, instance)
	q.byId[instance.ID()] = instance
	return instance
}

func (q *queue) Pop(image string) Instance {
	for idx, instance := range q.queue {
		if instance.State().Image == image {
			q.queue = append(q.queue[:idx], q.queue[idx+1:]...)
			// delete(q.byId, instance.ID())
			return instance
		}
	}
	return nil
}
