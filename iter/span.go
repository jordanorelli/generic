package iter

import (
	"constraints"
)

type span[T constraints.Integer] struct {
	next T
	final T
	step T
}

func (s *span[T]) Next(n *T) bool {
	if s.next >= s.final {
		return false
	}
	*n = s.next
	s.next += s.step
	return true
}

// Span creates a span of integers between start and end
func Span[T constraints.Integer](start, end T) Ator[T] {
	return &span[T]{next: start, final: end, step: 1}
}

// Step is the same as span, but allows for step sizes greater than 1
func Step[T constraints.Integer](start, end, step T) Ator[T] {
	return &span[T]{next: start, final: end, step: step}
}
