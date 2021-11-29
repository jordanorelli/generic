package merge

import (
	"fmt"
)

type Table[K comparable, V Merges[V]] map[K]V

func (t Table[K, V]) Merge(from Table[K, V]) error {
	for k, v := range from {
		e, ok := t[k]
		if !ok {
			var z V
			z = z.MergeIdentity()
			if err := z.Merge(v); err != nil {
				return fmt.Errorf("tables failed to merge: %w", err)
			}
			t[k] = z
			continue
		}

		if err := e.Merge(v); err != nil {
			return fmt.Errorf("tables failed to merge: %w", err)
		}
	}
	return nil
}
