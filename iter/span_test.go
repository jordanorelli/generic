package iter

import (
	"testing"
)

func TestSpan(t *testing.T) {
	var n int
	s := Span(1, 10)

	s.Next(&n)
	if n != 1 {
		t.Errorf("expected n to be 1 but is %d instead", n)
	}
	for s.Next(&n) {
	}
	if n != 9 {
		t.Errorf("expected n to be 9 but is %d instead", n)
	}

	beer := Max(Span(1, 100))
	if beer != 99 {
		t.Errorf("expected 99 beers but saw %d instead", beer)
	}

	old := Min(Span(30, 40))
	if old != 30 {
		t.Errorf("expected 30 to be old but saw %d instead", old)
	}
}

func TestStep(t *testing.T) {
	s := Step(1, 10, 3)
	for n := 0; s.Next(&n); {
		t.Log(n)
	}
}
