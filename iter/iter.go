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

	// An iterator must also be iterable. That is, a valid iterator must be
	// able to create a new iterator beginning at the same position of that
	// iterator, such that calling Iter() on an Iterator returns a new
	// Iterator, and that both can be iterated without affecting one another.
	Able[T]
}

// Start starts the iteration for an Iterable. This is a convenience function
// to facilitate iterating in a for loop. E.g., given an Iterable value l, such
// as a list, the following would iterate over the entire iterable:
//
//     for v, it := iter.Start(l); it.Next(&v); {
//         // utilize v here
//     }
//
// Without such a construct, a developer would have to do something similar instead:
// 
//     for v, it := 0, l.Iter(); it.Next(&v); {
//
//     }
//
// Doing that would require that the developer type out the zero-value for that
// type, when that value could just as easily be inferred.
func Start[T any](a Able[T]) (T, Ator[T]) {
	var v T
	return v, a.Iter()
}

func Min[T constraints.Ordered](src Able[T]) T {
	it := src.Iter()
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

func Max[T constraints.Ordered](src Able[T]) T {
	it := src.Iter()
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
