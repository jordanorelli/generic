package merge

import (
	"testing"
)

// additive implements an additive merge. Merging the two numbers together adds
// the source value into the receiver without changing the source value.
type additive struct {
	total int
}

func (a *additive) Merge(b *additive) error {
	a.total += b.total
	return nil
}

func add(n int) *additive {
	return &additive{total: n}
}

type multiplicative struct {
	scale int
}

func (m *multiplicative) Merge(v *multiplicative) error {
	m.scale *= v.scale
	return nil
}

func mul(n int) *multiplicative {
	return &multiplicative{scale: n}
}

// exclusive implements an exlsive merge. Merging the two numbers together adds
// the value from the source into the destination, removing it from the source.
type exclusive struct {
	stock int
}

func (e *exclusive) Merge(source *exclusive) error {
	e.stock += source.stock
	source.stock = 0
	return nil
}

func ex(n int) *exclusive {
	return &exclusive{stock: n}
}

func TestMerge(t *testing.T) {
	t.Run("additive", func(t *testing.T) {
		a, b := add(4), add(7)
		if err := Merge(a, b); err != nil {
			t.Errorf("merge error: %v", err)
		}
		if a.total != 11 {
			t.Errorf("merge failed to mutate destination")
		}
		if b.total != 7 {
			t.Errorf("merged caused unexpected mutation")
		}
	})

	t.Run("exclusive", func(t *testing.T) {
		a, b := ex(4), ex(7)
		if err := Merge(a, b); err != nil {
			t.Errorf("merge error: %v", err)
		}
		if a.stock != 11 {
			t.Errorf("merge failed to mutate destination")
		}
		if b.stock != 0 {
			t.Errorf("merge failed to mutate source")
		}
	})
}


