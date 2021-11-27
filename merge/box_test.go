package merge

import (
	"testing"
	"errors"
)

func TestBox(t *testing.T) {
	t.Run("matching", func(t *testing.T) {
		a, b := add(3), add(7)
		dest, src := Box(a), Box(b)
		if err := Merge(dest, src); err != nil {
			t.Fatalf("unexpected error merging boxes: %v", err)
		}

		if a.total != 10 {
			t.Error("box failed to mutate contents in merge")
		}
		if b.total != 7 {
			t.Error("box mutated source contents in merge for some reason")
		}
	})

	t.Run("mismatched", func(t *testing.T) {
		a, b := add(3), mul(7)
		dest, src := Box(a), Box(b)
		err := Merge(dest, src)
		if err == nil {
			t.Fatalf("mismatched merge succeeded but should have failed")
		}
		if !errors.Is(err, typeMismatch) {
			t.Fatalf("merge gave unexpected error value: %v", err)
		}


		if a.total != 3 {
			t.Error("failed merge still mutated values")
		}
	})
}

func TestUnbox(t *testing.T) {
	t.Run("matching", func(t *testing.T) {
		a := add(3)
		b := Box(a)

		v, err := Unbox[*additive](b)
		if err != nil {
			t.Fatalf("unexpected unbox error: %v", err)
		}
		if v.total != 3 {
			t.Fatalf("boxing and unboxing messed up the value somehow")
		}
	})

	t.Run("mismatched", func(t *testing.T) {
		a := add(3)
		b := Box(a)

		v, err := Unbox[*multiplicative](b)
		if err == nil {
			t.Fatalf("mismatched unboxing should have returned an error but succeeded and unboxed %v instead", v)
		}
		if !errors.Is(err, typeMismatch) {
			t.Fatalf("unbox expected to give typeMismach but instead gave unexpected error: %v", err)
		}
	})
}
