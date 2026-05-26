package functions

import (
	"github.com/elliotchance/pie/pie"
)

// StringsUsing transforms each element to a string.
func (ss SliceType) StringsUsing(transform func(ElementType) string) pie.Strings {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(pie.Strings)
}
