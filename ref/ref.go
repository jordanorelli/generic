package ref

import (
	"fmt"
)

// New creates a reference for a given pointer
func New[T any](v *T) Ref[T] {
	if v == nil {
		var zero T
		return Ref[T]{ptr: &zero}
	}
	return Ref[T]{ptr: v}
}

// Ref is a read reference to some value T
type Ref[T any] struct { ptr *T }

// Val reads the value for this reference
func (r Ref[T]) Val() T { return *r.ptr }

func (r Ref[T]) String() string { return fmt.Sprintf("ref{%v}", *r.ptr) }
