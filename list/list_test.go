package list

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	var l List[int]

	if !l.Empty() {
		t.Errorf("new list is not empty, but should be")
	}

	if l.Head() != 0 {
		t.Errorf("expect head of list to be zero-value for that type but saw %d instead of 0 for int", l.Head())
	}

	tail := l.Tail()
	if !tail.Empty() {
		t.Errorf("empty list should have empty tail but saw %v instead", tail)
	}

	if l.Len() != 0 {
		t.Errorf("empty list should have a length of 0 but has %d instead", l.Len())
	}
}

func TestOne(t *testing.T) {
	var l List[int]

	l.Push(3)

	if l.Empty() {
		t.Errorf("list should have 1 element but is empty")
	}

	if l.Head() != 3 {
		t.Errorf("list's head element should be 3 but saw %d instead", l.Head())
	}

	if l.Len() != 1 {
		t.Errorf("expected a list of size 1 but saw %d instead", l.Len())
	}
}

func TestMake(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		// This doesn't work because the type parameter T cannot be inferred:
		// Make[]()

		// You have to provide T in the case of an empty list
		l := Make[int]()
		if !l.Empty() {
			t.Errorf("make with no params should create empty list but gave %v instead", l)
		}
	})

	t.Run("some strings", func(t *testing.T) {
		// The type parameter T can be inferred based on the arguments passed
		// in
		l := Make("bob", "carol", "dave")

		if l.Len() != 3 {
			t.Errorf("expected length of 3 but saw %d instead", l.Len())
		}

		if l.Head() != "bob" {
			t.Errorf("expected a head element of %q but saw %q instead", "bob", l.Head())
		}

		l.Push("alice")
		if l.Head() != "alice" {
			t.Errorf("expected a head element of %q but saw %q instead", "alice", l.Head())
		}
		if l.Len() != 4 {
			t.Errorf("expected length of 4 but saw %d instead", l.Len())
		}
	})

	t.Run("mixed element types", func(t *testing.T) {
		l := Make[any]("alice", 3, "carol")

		if l.Len() != 3 {
			t.Errorf("expected length of 3 but saw %d instead", l.Len())
		}
	})
}

func mult[T wholeNumber](x T) func(T) T {
	return func(y T) T {
		return x * y
	}
}

type wholeNumber interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
	~int8 | ~int16 | ~int32 | ~int64 | int
}

func eq[T comparable](t *testing.T, expect T, found T) {
	if found != expect {
		t.Errorf("expected %v, found %v", expect, found)
	}
}

func TestMap(t *testing.T) {
	nums := Make(2, 4, 6).Map(mult(5))
	t.Logf("Nums: %v", nums)
	eq(t, 10, nums.At(0))
	eq(t, 20, nums.At(1))
	eq(t, 30, nums.At(2))
	eq(t, 0, nums.At(3))
}
