package daemon

import (
	"github.com/backtick-se/gowait/core/task"

	"context"
)

type Queue interface {
	Push(ctx context.Context, spec *task.Spec) (Instance, error)
	Pop(ctx context.Context, image string) (Instance, error)
}

type queue struct {
	length int

	// todo: make thread safe
	queues map[string]chan Instance
}

func NewQueue(length int) Queue {
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

func (m *queue) Pop(ctx context.Context, image string) (Instance, error) {
	queue := m.getQueue(image)
	select {
	case instance := <-queue:
		return instance, nil
	case <-ctx.Done():
		return nil, context.Canceled
	}
}

func (m *queue) Push(ctx context.Context, spec *task.Spec) (Instance, error) {
	instance := newInstance(spec)
	queue := m.getQueue(spec.Image)

	select {
	case queue <- instance:
		return instance, nil
	case <-ctx.Done():
		return nil, context.Canceled
	}
}
