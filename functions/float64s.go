package functions

import (
	"github.com/elliotchance/pie/pie"
)

// Float64s transforms each element to a float64.
func (ss SliceType) Float64s() pie.Float64s {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(pie.Float64s)
}
