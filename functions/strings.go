package functions

import (
	"github.com/elliotchance/pie/pie"
)

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//	fmt.Sprintf("%v")
func (ss SliceType) Strings() pie.Strings {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(pie.Strings)
}
