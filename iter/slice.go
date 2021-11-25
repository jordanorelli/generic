package iter

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

// Slice takes a slice and returns an iterable backed by that slice.
//
// ok so, asn an aside: normally you would write a function to return the
// concrete implementation, not an interface, but if I return the concrete
// implementation for an iter.Able, the type inference doesn't work correctly
// with the type parameters in some cases.
//
// So let's say we had an exported type Slice that was defined as:
//
//     type Slice[T any] []T
//
// And we tried to convert a []int to a Slice[int], we could so that with:
//
//     Slice([]int{1, 2, 3})
//
// And that would be fine until we try to do soething like pass it to Start.
// Returning Able[T] here means we can say Start(Slice([]int{1, 2, 3})), but if
// we had type Slice[T any] []T, we would have to say:
//
//     Start[int](Slice([]int{1, 2, 3}))
//
// ...and I think that's annoying.
func Slice[T any](s []T) Able[T] { return slice[T](s) }
