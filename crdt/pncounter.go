package crdt

// PNCounter is a positive-negative counter
type PNCounter[K comparable] struct {
	slots map[K][2]int
}
