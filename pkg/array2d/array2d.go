package array2d

type Array2D[T any] struct {
	s    []T
	w, h int32
}

func New[T any](w, h int32) Array2D[T] {
	length := w * h
	return Array2D[T]{
		s: make([]T, length),
		w: w,
		h: h,
	}
}

func (a Array2D[T]) Idx(x, y int32) (int32, error) {
	i := x + y*a.w
	if int(i) >= len(a.s) || x >= a.w || y >= a.h || x < 0 || y < 0 {
		return 0, ErrOutOfRange
	}
	return x + y*a.w, nil
}

func (a Array2D[T]) Get(x, y int32) (T, error) {
	var val T

	i, err := a.Idx(x, y)
	if err != nil {
		return val, err
	}
	val = a.s[i]

	return val, nil
}

func (a *Array2D[T]) Set(x, y int32, value T) error {
	i, err := a.Idx(x, y)
	if err != nil {
		return err
	}

	a.s[i] = value
	return nil
}

func (a Array2D[T]) GetW() int32 {
	return a.w
}

func (a Array2D[T]) GetH() int32 {
	return a.h
}
