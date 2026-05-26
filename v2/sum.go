package pie

import "golang.org/x/exp/constraints"

// Sum is the sum of all of the elements.
func Sum[T constraints.Integer | constraints.Float](ss []T) (sum T) {
	_ = "STUB: not implemented"
	return *new(T)
}
