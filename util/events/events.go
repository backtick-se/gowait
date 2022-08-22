package events

type Event interface{}

type Pub[T Event] interface {
	Publish(T)
	Subscribe() Sub[T]
}

type Sub[T Event] interface {
	Next() <-chan T
}

type pub[T Event] struct {
	subs  map[*sub[T]]bool
	queue int
}

func New[T Event]() Pub[T] {
	queue := 100
	return &pub[T]{
		subs:  make(map[*sub[T]]bool),
		queue: queue,
	}
}

func (p *pub[T]) Publish(event T) {
	for sub := range p.subs {
		sub.out <- event
	}
}

func (p *pub[T]) Subscribe() Sub[T] {
	sub := &sub[T]{
		pub: p,
		out: make(chan T, p.queue),
	}
	p.subs[sub] = true
	return sub
}

type sub[T Event] struct {
	pub *pub[T]
	out chan T
}

func (s *sub[T]) Next() <-chan T {
	return s.out
}

func (s *sub[T]) Unsub() {
	defer close(s.out)
	delete(s.pub.subs, s)
}
