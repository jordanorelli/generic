package list

type Iterable[T any] interface {
	Iter() Iter[T]
}

type Iter[T any] interface {
	Next(*T) bool
}

// This actually works, but for some reason I don't yet understand, defining
// Max with a constraint of constrains.Ordered that takes Iterable[T] cannot
// have its type parameter infered, but defining it with the same constraint of
// constraints.Ordered and taking List[T] does allow the type parameter to be
// inferred. Super confusing.
//
// func Max[T constraints.Ordered](l Iterable[T]) T {
// 	it := l.Iter()
// 
// 	var v T
// 	if !it.Next(&v) {
// 		return v
// 	}
// 	max := v
// 	for it.Next(&v) {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }


// Discarded Iterator types:
//
// This was my first attempt. I like to have a method that returns a bool so
// you can use it succinctly in a for loop, but I didn't love having to call
// both done and next
//
// type Iter[T any] interface {
// 	Done() bool
// 	Next() T
// }
//
// I was never optimistic about this. In practice, using this in a for loop is
// just annoying so I stopped doing it. But also there's another thing I find
// annoying: it's a copy every time. You're not really iterating over the
// values, you're iterating over copies of the values, it seemed like a lot of
// unecessary copying.
//
// type Iter[T any] interface {
//   Next() (T, bool)
// }
