package crdt

import (
	"testing"
)

func TestGCounter(t *testing.T) {
	t.Run("incr", func(t *testing.T) {
		g := NewGCounter[string]()
		if n := g.Total(); n != 0 {
			t.Fatalf("new gcounter has count of %d, should be 0", n)
		}

		if err := g.Incr("jordan"); err != nil {
			t.Fatalf("gcounter failed incr: %v", err)
		}

		if n := g.Total(); n != 1 {
			t.Fatalf("new gcounter has count of %d, should be 1", n)
		}

		if err := g.Incr(""); err == nil {
			t.Fatalf("incrementing the zero value succeeded, should have failed")
		}

		if err := g.Incr("jordan"); err != nil {
			t.Fatalf("gcounter failed incr: %v", err)
		}

		if n := g.Total(); n != 2 {
			t.Fatalf("new gcounter has count of %d, should be 2", n)
		}
	})

	t.Run("add", func(t *testing.T) {
		g := NewGCounter[string]()
		if n := g.Total(); n != 0 {
			t.Fatalf("new gcounter has count of %d, should be 0", n)
		}

		if err := g.Add("jordan", 4); err != nil {
			t.Fatalf("gcounter failed incr: %v", err)
		}

		if n := g.Total(); n != 4 {
			t.Fatalf("new gcounter has count of %d, should be 1", n)
		}

		if err := g.Add("", 10); err == nil {
			t.Fatalf("adding to zero key succeeded, should have failed")
		}

		if err := g.Add("jordan", -3); err == nil {
			t.Fatalf("adding negatively to the gcounter succeeded, should have failed")
		}

		if err := g.Add("jordan", 3); err != nil {
			t.Fatalf("gcounter failed add: %v", err)
		}

		if n := g.Total(); n != 7 {
			t.Fatalf("new gcounter has count of %d, should be 7", n)
		}
	})
}
