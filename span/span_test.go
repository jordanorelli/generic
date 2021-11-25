package span

import (
	"testing"

	"github.com/jordanorelli/generic/iter"
)

func TestSpan(t *testing.T) {
	var n int
	it := New(1, 10).Iter()

	it.Next(&n)
	if n != 1 {
		t.Errorf("expected n to be 1 but is %d instead", n)
	}
	for it.Next(&n) {
	}
	if n != 9 {
		t.Errorf("expected n to be 9 but is %d instead", n)
	}

	// If the function New returns a value of iter.Able[T], this type parameter
	// can be inferred, but when New returns a value of Span[T] (which
	// satisfies iter.Able[T]), the type parameter cannot be inferred. I don't
	// know why this behavior exists or if this is the intended behavior.
	//                |
	//                |
	//                V
	beer := iter.Max[int](New(1, 100))
	if beer != 99 {
		t.Errorf("expected 99 beers but saw %d instead", beer)
	}

	old := iter.Min[int](New(30, 40))
	if old != 30 {
		t.Errorf("expected 30 to be old but saw %d instead", old)
	}

	// honestly a very ridiculous use-case, but I'm looking at some weird
	// inferrence corner cases. It seems the third parameter can't be the
	// literal 2, you have to explicitly type it.
	alpha := Step('a', 'z'+1, rune(2))

	// the rune type can't be inferred here either
	for a, it := iter.Start[rune](alpha); it.Next(&a); {
		t.Logf("%c", a)
	}

	// the rune type can't be inferred here either!
	for a, it := iter.Start[rune](alpha.Iter()); it.Next(&a); {
	}

	t.Logf("%T", iter.Max[int8](New[int8](3, 10)))
}

func TestStep(t *testing.T) {
	it := Step(1, 10, 3).Iter()
	for n := 0; it.Next(&n); {
		t.Log(n)
	}
}
