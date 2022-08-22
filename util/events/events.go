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
	in    chan T
	subs  map[*sub[T]]bool
	queue int
}

func New[T Event]() Pub[T] {
	queue := 100
	pub := &pub[T]{
		subs:  make(map[*sub[T]]bool, queue),
		in:    make(chan T, queue),
		queue: queue,
	}
	go pub.proc()
	return pub
}

func (p *pub[T]) proc() {
	for {
		ev, ok := <-p.in
		if !ok {
			break
		}

		for sub := range p.subs {
			sub.out <- ev
		}
	}
}

func (p *pub[T]) Publish(event T) {
	p.in <- event
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
