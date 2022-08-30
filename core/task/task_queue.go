package task

import (
	"context"
)

type TaskQueue interface {
	Queue(ctx context.Context, instance Instance) error
	Aquire(ctx context.Context, image string) (Instance, error)
}

type queue struct {
	length int

	// todo: make thread safe
	queues map[string]chan Instance
}

func NewTaskQueue(length int) TaskQueue {
	return &queue{
		length: length,
		queues: make(map[string]chan Instance),
	}
}

func (m *queue) getQueue(image string) chan Instance {
	queue, exists := m.queues[image]
	if !exists {
		queue = make(chan Instance, m.length)
		m.queues[image] = queue
	}
	return queue
}

func (m *queue) Aquire(ctx context.Context, image string) (Instance, error) {
	queue := m.getQueue(image)
	for {
		select {
		case instance := <-queue:
			// skip any task that may have been cancelled
			if instance.State().Status != StatusWait {
				continue
			}
			return instance, nil
		case <-ctx.Done():
			return nil, context.Canceled
		}
	}
}

func (m *queue) Queue(ctx context.Context, instance Instance) error {
	queue := m.getQueue(instance.Spec().Image)

	select {
	case queue <- instance:
		return nil
	case <-ctx.Done():
		return context.Canceled
	}
}
