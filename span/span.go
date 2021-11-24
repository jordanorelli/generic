package span

import (
	"constraints"

	"github.com/jordanorelli/generic/iter"
)

type Span[T constraints.Integer] struct {
	Start T
	End T
	Step T
}

type spanIter[T constraints.Integer] struct {
	start T
	end T
	step T
	next T
}

func (s Span[T]) Iter() iter.Ator[T] {
	return &spanIter[T]{
		start: s.Start,
		end: s.End,
		step: s.Step,
		next: s.Start,
	}
}

func (s *spanIter[T]) Next(n *T) bool {
	if s.step == 0 || s.next < s.start || s.next >= s.end {
		return false
	}
	*n = s.next
	s.next += s.step
	return true
}

func (s spanIter[T]) Iter() iter.Ator[T] { return &s }

// New creates a span of integers between start and end. The is analagous to
// the "range" function in Python, but since range already means something in
// Go, span is the chosen name to avoid confusion with Go's concept of range.
// The created span is given a step size of 1.
func New[T constraints.Integer](start, end T) Span[T] {
	return Span[T]{
		Start: start,
		End: end,
		Step: 1,
	}
}

// Step is the same as creating a span with a provided step value
func Step[T constraints.Integer](start, end, step T) Span[T] {
	return Span[T]{
		Start: start,
		End: end,
		Step: step,
	}
}
