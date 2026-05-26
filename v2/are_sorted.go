package pie

import (
	"golang.org/x/exp/constraints"
)

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.SliceIsSorted.
func AreSorted[T constraints.Ordered](ss []T) bool { _ = "STUB: not implemented"; return false }
