package pie

import (
	"math/rand"
)

// Shuffle returns a new shuffled slice by your rand.Source. The original slice
// is not modified.
func Shuffle[T any](ss []T, source rand.Source) []T {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return nil
}
