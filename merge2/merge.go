package merge2

import (
	"fmt"
)

type Merges[X any] interface {
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

// Into merges any number of values into some commonly-targetable value.
func Into[X Merges[Y], Y any](dest *Y, sources ...X) error {
	for _, src := range sources {
		if err := src.MergeInto(dest); err != nil {
			return err
		}
	}
	return nil
}

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

type Boxed struct {
	p interface{}
	merge func(interface{}) error
}

func (b Boxed) String() string {
	return fmt.Sprint(b.p)
}

func (b Boxed) IsEmpty() bool {
	return b.p == nil
}

func Box[X Merges[X]](x X) Boxed {
	return Boxed{
		p: &x,
		merge: func(dest interface{}) error {
			return fmt.Errorf("not yet")
			// dp, ok := dest.p.(*X)
			// if !ok {
			// 	return fmt.Errorf("boxed val tried to merge into %T but can only merge into %T", dest, dp)
			// }
			// if dp == nil {
			// 	return fmt.Errorf("that shit is empty fuck you")
			// }
			// fmt.Printf("merging %v into %v\n", x, dp)
			// return nil
		},
	}
}

func (b Boxed) MergeInto(dest interface{}) error {
	if b.IsEmpty() {
		return nil
	}
	return b.merge(dest)
}

func Unbox[X Merges[Y], Y any](b Boxed) (X, error) {
	p, ok := b.p.(*X)
	if !ok {
		var zero X
		return zero, fmt.Errorf("type error")
	}
	return *p, nil
}
