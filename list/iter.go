package list

type Iterable[T any] interface {
	Iter() Iter[T]
}

type Iter[T any] interface {
	Next(*T) bool
}

/*
type Iter[T any] interface {
	Done() bool
	Next() T
}
*/
