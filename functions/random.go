package functions

import (
	"math/rand"
)

// Random returns a random element by your rand.Source, or zero
func (ss SliceType) Random(source rand.Source) ElementType {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(ElementType)
}
