package merge2

import (
	"testing"
	"constraints"
)

type additive int

func (a additive) MergeInto(dest *additive) error {
	*dest += a
	return nil
}

type multiplicative int

func (m multiplicative) MergeInto(dest *multiplicative) error {
	*dest *= m
	return nil
}

func newGCounter[ID comparable]() gcounter[ID] {
	return gcounter[ID]{
		slots: make(map[ID]int),
	}
}

type gcounter[ID comparable] struct {
	slots map[ID]int
}

func (g gcounter[ID]) incr(id ID) {
	g.slots[id]++
}

func (g gcounter[ID]) add(id ID, n int) {
	if n < 0 {
		panic("no")
	}
	g.slots[id] += n
}

func max[N constraints.Ordered](a, b N) N {
	if a >= b {
		return a
	}
	return b
}

func (g gcounter[ID]) MergeInto(dest *gcounter[ID]) error {
	if dest.slots == nil {
		dest.slots = make(map[ID]int)
	}
	for id, count := range g.slots {
		dest.slots[id] = max(count, dest.slots[id])
	}
	return nil
}

// func TestTogether(t *testing.T) {
// 	a, b, c := additive(3), additive(10), additive(17)
// 	total, err := Together[Merges[*additive], additive](a, b, c)
// 	if err != nil {
// 		t.Fatalf("total error: %v", err)
// 	}
// 	if total != 30 {
// 		t.Fatalf("%d != 30", total)
// 	}
// }

func TestInto(t *testing.T) {
	a, b, c := additive(3), additive(10), additive(17)
	if err := Into(&a, b, c); err != nil {
		t.Fatalf("error: %v", err)
	}
	if a != 30 {
		t.Fatalf("%d != 30", a)
	}
}

func TestMaps(t *testing.T) {
	alice := map[string]additive{
		"vanilla": 3,
		"chocolate": 5,
		"strawberry": 2,
	}

	bob := map[string]additive{
		"vanilla": 2,
		"chocolate": 3,
		"pistacchio": 5,
	}

	totals := make(map[string]additive)
	if err := Maps(totals, alice); err != nil {
		t.Fatalf("map error: %v", err)
	}
	Maps(totals, bob)
	t.Log(totals)
}

func TestMap(t *testing.T) {
	alice := Map[string, additive]{
		"vanilla": 3,
		"chocolate": 5,
		"strawberry": 2,
	}
	t.Log(alice)
}

func TestBox(t *testing.T) {
	t.Run("both empty", func (t *testing.T) {
		var a Boxed
		var b Boxed
		if err := a.MergeInto(b); err != nil {
			t.Fatalf("empty boxes failed to merge: %v", err)
		}
	})

	t.Run("empty source", func(t *testing.T) {
		var a Boxed
		b := Box[additive, additive](additive(3))

		if err := a.MergeInto(b); err != nil {
			t.Error(err.Error())
		}
	})

	// t.Run("empty destination", func(t *testing.T) {
	// 	a := Box[additive, additive](additive(3))
	// 	var b Boxed

	// 	if err := a.MergeInto(b); err != nil {
	// 		t.Error(err.Error())
	// 	}
	// })
	
	// an empty box can merge into anything and it should be a no-op
	// t.Run("empty source", func (t *testing.T) {
	// 	var a Boxed
	// 	b := additive(3)
	// 	if err := a.MergeInto(b); err != nil {
	// 		t.Fatalf("empty boxes failed to merge: %v", err)
	// 	}
	// })

	// t.Run("homologous", func(t *testing.T) {
	// 	a, b := Box(additive(3)), Box(additive(8))
	// 	if err := a.MergeInto(b); err != nil {
	// 		t.Fatalf("homologous boxes failed to merge: %v", err)
	// 	}
	// })

	// a, b := Box(additive(3)), Box(additive(5))
	// if err := b.MergeInto(&a); err != nil {
	// 	t.Errorf("merge boxed failed: %v", err)
	// }

	// total, err := Unbox[additive](a)
	// if err != nil {
	// 	t.Errorf("unbox failed: %v", err)
	// }
	// if total != 8 {
	// 	t.Errorf("%d != 8", total)
	// }

	// alice := Map[string, Boxed]{
	// 	"vanilla": Box(additive(3)),
	// 	"chocolate": Box(multiplicative(5)),
	// 	"strawberry": Box(multiplicative(2)),
	// }

	// bob := Map[string, Boxed]{
	// 	"vanilla": Box(additive(3)),
	// 	"chocolate": Box(additive(5)),
	// 	"pistacchio": Box(additive(2)),
	// }

	// Maps(alice, bob)
	// t.Log(alice)
}
