package merge

import (
	"fmt"
	"errors"
)

// Boxed is a boxed value that is able to merge with other values. Boxing a
// value pushes its type-checking from compile time to run time. This can be
// useful in the event that you wish to construct mergeable values with
// heterogeneous members, but should likely be avoided otherwise.
type Boxed struct {
	val   interface{}
	ident func() interface{}
	merge func(interface{}) error
}

func (b Boxed) Merge(from Boxed) error {
	if err := b.merge(from.val); err != nil {
		return fmt.Errorf("boxed merge failed: %w", err)
	}
	return nil
}

func (b Boxed) MergeIdentity() Boxed {
	return Boxed{
		val: b.ident(),
		ident: b.ident,
		merge: b.merge,
	}
}

var typeMismatch = errors.New("mismatched types")

// strip takes a typed function and erases the type information from its
// parameter
func strip[X any](f func(X) error) func(interface{}) error {
	return func(v interface{}) error {
		vv, ok := v.(X)
		if !ok {
			return fmt.Errorf("unexpected %T value, expected %T instead: %w", v, vv, typeMismatch)
		}
		return f(vv)
	}
}

func stripReturn[X any](f func() X) func() interface{} {
	return func() interface {} { return f() }
}

// Box takes a mergeable value and creates a new mergeable value of type Boxed.
// Any two Boxed values can attempt to merge at runtime.
func Box[X Merges[X]](x X) Boxed {
	return Boxed{
		val: x,
		ident: stripReturn(x.MergeIdentity),
		merge: strip(x.Merge),
	}
}

// Unbox removes a value from its box.
func Unbox[X Merges[X]](b Boxed) (X, error) {
	if v, ok := b.val.(X); ok {
		return v, nil
	} else {
		var zero X
		return zero, fmt.Errorf("box contains %T, not %T: %w", b.val, zero, typeMismatch)
	}
}
