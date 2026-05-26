package pie

import "golang.org/x/exp/constraints"

// Product is the product of all of the elements.
func Product[T constraints.Integer | constraints.Float](ss []T) (product T) {
	_ = "STUB: not implemented"
	return *new(T)
}
