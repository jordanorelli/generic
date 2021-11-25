package iter

import "constraints"

type slice[T any] []T

type sliceIter[T any] struct {
	s slice[T]
	i int
}

func (it *sliceIter[T]) Next(v *T) bool {
	if it.i >= len(it.s) {
		return false
	}
	*v = it.s[it.i]
	it.i++
	return true
}

func (it *sliceIter[T]) Iter() Ator[T] { return &sliceIter[T]{s: it.s} }

func (s slice[T]) Iter() Ator[T] { return &sliceIter[T]{s: s} }

// Slice takes a slice and returns an iterable backed by that slice
func Slice[T constraints.Integer](s []T) Able[T] { return slice[T](s) }
