package slices

func Map[T any, S any](items []T, transform func(T) S) []S {
	out := make([]S, len(items))
	for i, item := range items {
		out[i] = transform(item)
	}
	return out
}
