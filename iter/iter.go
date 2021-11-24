package iter

import "constraints"

// Able is anything that is iter.Able
type Able[T any] interface {
	Iter() Ator[T]
}

// Ator is an iter.Ator
type Ator[T any] interface {
	// Next assigns the value pointed at by *T to the next value in the
	// iteration. Next should return true if a value was set. An Ator that
	// returns false should continue to return false; once iteration is
	// complete it should be complete forever.
	Next(*T) bool
}

func Ate[T any](a Able[T]) (T, Ator[T]) {
	var v T
	return v, a.Iter()
}

func Min[T constraints.Ordered](it Ator[T]) T {
	var v T
	if !it.Next(&v) {
		var zero T
		return zero
	}
	min := v
	for it.Next(&v) {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T constraints.Ordered](it Ator[T]) T {
	var v T
	if !it.Next(&v) {
		var zero T
		return zero
	}
	max := v
	for it.Next(&v) {
		if v > max {
			max = v
		}
	}
	return max
}
