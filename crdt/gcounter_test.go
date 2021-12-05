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

	t.Run("merge", func(t *testing.T) {
		type hostname string

		host1 := NewGCounter[hostname]()
		host2 := NewGCounter[hostname]()

		host1.Add("host1", 3)
		host2.Add("host2", 6)

		host2.MergeInto(&host1)
		if n := host1.Total(); n != 9 {
			t.Errorf("merging once produced %d instead of 9", n)
		}

		host2.MergeInto(&host1)
		if n := host1.Total(); n != 9 {
			t.Errorf("merging a second time changed the target")
		}

		if n := host2.Total(); n != 6 {
			t.Errorf("merging changed the source")
		}
		host1.MergeInto(&host2)
		if host1.Total() != host2.Total() {
			t.Errorf("merging both ways didn't converge")
		}

		host2.Incr("host2")
		if host1.Total() == host2.Total() {
			t.Errorf("improper post-merge mutation")
		}
	})
}
