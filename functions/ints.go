package functions

import (
	"github.com/elliotchance/pie/pie"
)

// Ints transforms each element to an integer.
func (ss SliceType) Ints() pie.Ints {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(pie.Ints)
}
