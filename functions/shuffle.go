package functions

import (
	"math/rand"
)

// Shuffle returns shuffled slice by your rand.Source
func (ss SliceType) Shuffle(source rand.Source) SliceType {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(SliceType)
}

// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
// the algorithm directly from the go source: src/math/rand/rand.go below,
// with some adjustments:
