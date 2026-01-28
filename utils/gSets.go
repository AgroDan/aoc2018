package utils

// and finally, gSets, a generic finite set implementation.
// Note that I'm setting it to use a struct{} because it uses
// zero bytes in memory. If I use bool, it will use one bit.

var gExists = struct{}{}

type GSet[T comparable] struct {
	elements map[T]struct{}
}

func NewGSet[T comparable]() *GSet[T] {
	s := &GSet[T]{
		elements: make(map[T]struct{}),
	}
	return s
}

func (s *GSet[T]) Add(value T) {
	s.elements[value] = gExists
}

func (s *GSet[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *GSet[T]) Contains(value T) bool {
	_, c := s.elements[value]
	return c
}

func (s *GSet[T]) Size() int {
	return len(s.elements)
}

func (s *GSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// merge two sets
func (s *GSet[T]) Union(other *GSet[T]) *GSet[T] {
	result := NewGSet[T]()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

// only elements in both sets
func (s *GSet[T]) Intersection(other *GSet[T]) *GSet[T] {
	result := NewGSet[T]()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
