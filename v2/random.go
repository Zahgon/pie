package pie

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// Random returns a random element by your rand.Source, or zero.
func Random[T constraints.Integer | constraints.Float](ss []T, source rand.Source) T {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(T)
}
