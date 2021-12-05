package crdt

import (
	"fmt"
	"constraints"
)

func max[N constraints.Ordered](a, b N) N {
	if a >= b {
		return a
	}
	return b
}

func NewGCounter[K comparable]() GCounter[K] {
	return GCounter[K]{slots: make(map[K]int)}
}

type GCounter[K comparable] struct {
	slots map[K]int `json:"slots"`
}

func (g GCounter[K]) Incr(slot K) error {
	var zero K
	if slot == zero {
		return fmt.Errorf("gcounter refuses incr on the zero-value of its key")
	}

	g.slots[slot]++
	return nil
}

func (g GCounter[K]) Add(slot K, delta int) error {
	var zero K
	if slot == zero {
		return fmt.Errorf("gcounter refuses add on the zero-value of its key")
	}

	if delta < 0 {
		return fmt.Errorf("gcounters cannot go down, use a pncounter instead")
	}
	g.slots[slot] += delta
	return nil
}

// Merge into some destination val
func (g *GCounter[K]) MergeInto(dest *GCounter[K]) {
	for slot, count := range g.slots {
		v := max(count, dest.slots[slot])
		dest.slots[slot] = v
	}
}

func (g *GCounter[K]) Total() int {
	var n int
	for _, count := range g.slots {
		n += count
	}
	return n
}
