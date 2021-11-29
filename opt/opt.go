package opt

// Val is an optional value
type Val[T any] struct {
	val T
	ok bool
}

// New creates an optional value
func New[T any](v T, ok bool) Val[T] {
	if ok {
		return Some(v)
	}
	return None[T]()
}

// None creates an empty optional value
func None[T any]() Val[T] { return Val[T]{} }

// NoneOf is the same as none, but it takes a parameter that it throws away.
// This allows the type T to be inferred.
func NoneOf[T any](v T) Val[T] { return Val[T]{} }

// Some creates a filled optional value
func Some[T any](v T) Val[T] {
	return Val[T]{
		val: v,
		ok: true,
	}
}

// Open retrives the contents of our optional value
func (v Val[T]) Open() (T, bool) {
	return v.val, v.ok
}

// Bind takes a function that doesn't understand optionals and gives you
// another function that does
func Bind[X, Y any](f func(X) Y) func(Val[X]) Val[Y] {
	return func(mx Val[X]) Val[Y] {
		if x, ok := mx.Open(); ok {
			return Some(f(x))
		}
		return None[Y]()
	}
}
