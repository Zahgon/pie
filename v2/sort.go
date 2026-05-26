package pie

import (
	"golang.org/x/exp/constraints"
)

// Sort works similar to sort.SliceType(). However, unlike sort.SliceType the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func Sort[T constraints.Ordered](ss []T) []T {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return nil
}
