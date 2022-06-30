package fn

// Queue not goroutine-safe
type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

func (q *Queue[T]) Pop() (T, error) {
	if q.Empty() {
		var d T
		return d, ErrEmpty
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v, nil
}

func (q *Queue[T]) PopAll() []T {
	ret := q.values[:]
	q.values = q.values[0:]
	return ret
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return len(q.values)
}
