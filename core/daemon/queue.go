package daemon

type Queue struct {
	queue []*instance
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]*instance, 0, 32),
	}
}

func (q *Queue) Enqueue(instance *instance) {
	q.queue = append(q.queue, instance)
}

func (q *Queue) Dequeue(image string) *instance {
	for idx, item := range q.queue {
		if item.run.Image == image {
			q.queue = append(q.queue[:idx], q.queue[idx+1:]...)
			return item
		}
	}
	return nil
}
