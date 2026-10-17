package constructor

type OptionFunc[T any] func(*T)

func New[T any](t *T, opts ...OptionFunc[T]) *T {
	if t == nil {
		t = new(T)
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}
