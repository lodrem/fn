package fn

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func NewSetFromSlice[T comparable](xs []T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]struct{}, len(xs)),
	}

	for _, x := range xs {
		s.Add(x)
	}
	return s
}

func (s *Set[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.m[value]
	return found
}

func (s *Set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Empty() bool {
	return len(s.m) == 0
}

func (s *Set[T]) ToSlice() []T {
	return Keys(s.m)
}

func (s *Set[T]) ForEach(fn func(T)) {
	for k := range s.m {
		fn(k)
	}
}

// Intersect returns a set containing the elements present in both the sets.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	res := NewSet[T]()

	var (
		s1 = s
		s2 = other
	)
	if s1.Size() > s2.Size() {
		s1, s2 = s2, s1
	}
	s1.ForEach(func(v T) {
		if s2.Contains(v) {
			res.Add(v)
		}
	})

	return res
}
