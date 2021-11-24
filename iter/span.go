package iter

import (
	"constraints"
)

type span[T constraints.Integer] struct {
	start T
	end T
	step T
}

type spanIter[T constraints.Integer] struct {
	start T
	end T
	step T
	next T
}

func (s span[T]) Iter() Ator[T] {
	return &spanIter[T]{
		start: s.start,
		end: s.end,
		step: s.step,
		next: s.start,
	}
}

func (s *spanIter[T]) Next(n *T) bool {
	if s.next >= s.end {
		return false
	}
	*n = s.next
	s.next += s.step
	return true
}

func (s spanIter[T]) Iter() Ator[T] { return &s }

// Span creates a span of integers between start and end. The is analagous to
// the "range" function in Python, but since range already means something in
// Go, span is the chosen name to avoid confusion with Go's concept of range.
func Span[T constraints.Integer](start, end T) Able[T] {
	return &span[T]{
		start: start,
		end: end,
		step: 1,
	}
}

// Step is the same as span, but allows for step sizes greater than 1
func Step[T constraints.Integer](start, end, step T) Able[T] {
	return &span[T]{
		start: start,
		end: end,
		step: step,
	}
}
