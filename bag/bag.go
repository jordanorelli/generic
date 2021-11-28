package bag

import (
	"errors"
)

var errNotFound = errors.New("not found")
var errTypeError = errors.New("type error")

// Bag is a read-only collection of values. Consumer can add elements to the
// bag if and only no element has been added for that key in the past. Elements
// can be added to the bag either as values or as pointers.
type Bag map[string]bagged

// Add adds a value to a bag. The provided value can be retrieved from the bag
// directly. There's really no reason to call this with a pointer but I don't
// know how to prevent that.
func Add(b Bag, k string, v interface{}) bool {
	if b.Has(k) {
		return false
	}
	b[k] = bagged{val: v}
	return true
}

// Ref adds a reference to a bag. The provided value must be a pointer. Once
// added, the pointer is never retrievable from the bag; reading this key from
// the bag dereferences the pointer at the time of reading.
func Ref[V any](b Bag, k string, v *V) bool {
	if b.Has(k) {
		return false
	}

	if v == nil {
		return false
	}

	b[k] = bagged{val: v, ref: true}
	return true
}

// Get retrieves a value from a bag. Whether a value was added or a ref was
// added, you always get a value out.
func Get[V any](b Bag, k string) (V, error) {
	bv, ok := b[k]
	if !ok {
		var zero V
		return zero, errNotFound
	}

	if bv.ref {
		ptr, ok := bv.val.(*V)
		if !ok {
			var zero V
			return zero, errTypeError
		}
		return *ptr, nil
	}

	v, ok := bv.val.(V)
	if !ok {
		var zero V
		return zero, errTypeError
	}
	return v, nil
}

// Has describes whether or not the bag contains the given key
func (b Bag) Has(k string) bool {
	_, ok := b[k]
	return ok
}

type bagged struct {
	val interface{}
	ref bool
}
