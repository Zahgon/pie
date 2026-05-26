package pie

import "golang.org/x/exp/constraints"

// Abs returns the absolute value.
func Abs[T constraints.Integer | constraints.Float](val T) T {
	_ = "STUB: not implemented"
	return *new(T)
}
