package merge

// Merges defines the ability to merge a value with another value. Note that
// while any type can satisfy Merges[X], practically speaking this interface is
// only useful to form the constraint [X Merges[X]]. E.g., in the following
// example, the type A merges with itself:
//
//     type A struct {}
//     func (a *A) Merge(*A) error { ... }
//
// This is the recommended pattern of implementation for utilizing the Merges
// interface. In this example, the type B merges with a different, type, type C:
//
//      type B struct {}
//      type C struct {}
//      func (b *B) Merge(*C) error { ... }
//
// Although the type B satisfies the interface Merges[*C], it does not satisfy
// the constraint [X Merges[X]], which is what is used throughout this package.
type Merges[X any] interface {
	// MergeIdentity defines the identity for the merge function
	MergeIdentity() X
	Merge(X) error
}

// Merge takes two values of any type that define their own semantics of how to
// merge one value into another value of the same type. X in this case must be
// a mutable type.
func Merge[X Merges[X]](dest, src X) error {
	return dest.Merge(src)
}
