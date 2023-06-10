package sets

type Set[T comparable] struct {
	m map[T]struct{}
}

func (s *Set[T]) Add(item T) {
	s.m[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exist := s.m[item]
	return exist
}

func (s *Set[T]) Clear() {
	s.m = map[T]struct{}{}
}

func (s *Set[T]) Raw() []T {
	var (
		ret []T
	)
	for k := range s.m {
		ret = append(ret, k)
	}
	return ret
}

func (s *Set[T]) Len() int64 {
	return int64(len(s.m))
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		m: map[T]struct{}{},
	}
}

func FromRawSlice[T comparable](s []T) *Set[T] {
	ret := New[T]()
	for _, item := range s {
		ret.Add(item)
	}
	return ret
}
