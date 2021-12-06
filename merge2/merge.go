package merge2

import (
	"fmt"
	"errors"
)

var typeMismatch = errors.New("type mismatch")

// Merges[X] is any type that can merge itself into some *X
type Merges[X any] interface {
	// MergeInto should merge the receiver into the parameter *X. Callers are
	// expected to mutate the parameter in some way, returning any error
	// encountered when attempting to do so.
	MergeInto(*X) error
}

// func Together[X Merges[Y], Y any](vals ...X) (Y, error) {
// 	var dest Y
// 	for _, v := range vals {
// 		if err := v.MergeInto(&dest); err != nil {
// 			return dest, fmt.Errorf("error merging values: %w", err)
// 		}
// 	}
// 	return dest, nil
// }
//

// Into merges any number of values into some commonly-targetable value.
// func Into[X any](dest *X, sources ...Merges[X]) error {
func Into[X any, Y Merges[X]](dest *X, sources ...Y) error {
	for _, src := range sources {
		if err := src.MergeInto(dest); err != nil {
			return err
		}
	}
	return nil
}

// Maps merges two maps. All keys that appear in source are merged into dest.
// Keys that appear in src but not in dest merge into the zero value of T and
// then store in dest.
func Maps[K comparable, V Merges[T], T any](dest map[K]T, src map[K]V) error {
	for k, v := range src {
		dv := dest[k]
		if err := v.MergeInto(&dv); err != nil {
			return fmt.Errorf("map merge error at key %v: %w", k, err)
		}
		dest[k] = dv
	}
	return nil
}

// hmmmmm is it possible to make a map type where the keys are anythign
// comparable and the values are any mixed set of values that merge into a
// single type?
//
// I sorta want this:
// Map[K comparable, V Merges[T], T any]

type Map[K comparable, V Merges[V]] map[K]V

func (m Map[K, V]) MergeInto(dest map[K]V) error {
	for k, v := range m {
		dv := dest[k]
		if err := v.MergeInto(&dv); err != nil {
			return fmt.Errorf("merge failed on key %v: %w", k, err)
		}
		dest[k] = dv
	}
	return nil
}

// Emerge is an empty merge: given some merge source, emerge merges the merge
// source into the zero value of the merge destination type.
//
// Emerge[Y any](x Merges[Y]) (Y, error)
func Emerge[X Merges[Y], Y any](x X) (Y, error) {
	var y Y
	return y, x.MergeInto(&y)
}

// Boxed is a type-erased container for merge semantics.
//
// Boxed exists in order to facilitate the merging of heterogeneous collections
// of values.
type Boxed struct {
	// the value inside of the box. Since the only way to put a value in a box
	// is through the Box function or by merging from another box, we know that
	// this value defines some merge function
	val interface{}

	// mkdest creates a value of the type to which the val field merges
	mkdest func() interface{}

	// homologous describes whether or not the contained value merges with its
	// own type. Any boxed type that defines merge semantics against its own
	// type can merge into an empty box.
	homologous bool

	// mergeInto is a function that defines how to merge the val field into
	// some destination
	mergeInto func(dest interface{}) error
}

func (b Boxed) String() string {
	return fmt.Sprint(b.val)
}

// IsEmpty is true for boxes that contain nothing
func (b Boxed) IsEmpty() bool {
	return b.val == nil
}

// IsTerminal describes whether or not the box defines any merge semantics. A
// box containing a value that does not define merge semantics is a box that
// terminates a merge chain.
func (b Boxed) IsTerminal() bool {
	return b.mergeInto == nil
}

// erasef1e erases type information from a function of input arity 1 that
// returns an error
func erasef1e[X any](f func(X) error) func(interface{}) error {
	return func(v interface{}) error {
		tv, ok := v.(X)
		if !ok {
			return fmt.Errorf("unexpected %T value, expected %T instead: %w", v, tv, typeMismatch)
		}
		return f(tv)
	}
}

func mergeFn[X Merges[Y], Y any](x X) func(interface{}) error {
	return nil
	// return func(v interface{}) error {
	// 	x.MergeInto
	// }
}

// nunu[X] creates a function for the type X that creates a new zero value
// having the type X, then erases the type information by sticking it in an
// empty interface. This is a constructor-constructor that creates a
// type-erased constructor. Pretty gross!
func nunu[X any]() func() interface {} {
	return func() interface{} {
		var zero X
		return zero
	}
}

// sametype determines whether or not some manually instantiated type
// parameters are or are not the same type.
func sametype[X, Y any]() bool {
	_, ok := interface{}(*new(X)).(Y)
	return ok
}

// Box boxes a value. Only values that define -some- merge semantic can go into
// the box.
//
// Box[Y any](x Merges[Y]) Boxed
func Box[X Merges[Y], Y any](x X) Boxed {
	return Boxed{
		val: x,
		mkdest: nunu[Y](),
		homologous: sametype[X, Y](),
		mergeInto: erasef1e(x.MergeInto),
	}
}

// EndBox creates a terminal box in a merge chain. This boxed value contains a
// value that does not define any merge semantics. It may only be used as a
// merge destination, not a merge source.
func EndBox[X any](x X) Boxed {
	return Boxed{
		val: x,
	}
}

// MergeInto merges the Boxed value into some destination value. If the boxed
// value does not merge into the value supplied, the failure occurs at runtime.
// For merge semantics that are type-checked at compile time... don't box your
// values, I dunno what to tell you.
func (b Boxed) MergeInto(dest interface{}) error {
	if b.IsEmpty() {
		return nil
	}
	if db, ok := dest.(Boxed); ok {
		if db.IsEmpty() {
			v := b.mkdest()
			if b.homologous {
				db.homologous = true

				db.mergeInto = erasef1e(v.MergeInto)
				db.mkdest = b.mkdest
			}
		}
	}
	return b.mergeInto(dest)
}

// func Unbox[X Merges[Y], Y any](b Boxed) (X, error) {
// 	p, ok := b.val.(*X)
// 	if !ok {
// 		var zero X
// 		return zero, fmt.Errorf("type error")
// 	}
// 	return *p, nil
// }
