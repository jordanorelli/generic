package merge

import (
	"testing"
)

func TestMergeTables(t *testing.T) {
	alice := Table[string, *additive]{
		"vanilla": add(3),
		"chocolate": add(5),
		"strawberry": add(2),
	}

	bob := Table[string, *additive]{
		"vanilla": add(2),
		"chocolate": add(3),
		"pistacchio": add(5),
	}

	votes := make(Table[string, *additive])
	if err := votes.Merge(alice); err != nil {
		t.Fatalf("tables failed to merge: %v", err)
	}
	if err := votes.Merge(bob); err != nil {
		t.Fatalf("tables failed to merge: %v", err)
	}

	check := func(k string, n int) {
		if have := votes[k].total; have != n {
			t.Fatalf("expected %d votes for %s but saw %v instead", n, k, have)
		}
	}
	check("vanilla", 5)
	check("chocolate", 8)
	check("strawberry", 2)
	check("pistacchio", 5)

}
