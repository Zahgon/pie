package pie

import "golang.org/x/exp/constraints"

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
//
// If the number of elements is even, then the ElementType mean of the two
// "median values" is returned.
func Median[T constraints.Integer | constraints.Float](ss []T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// This implementation aims at linear time O(n) on average.
// It uses the same idea as QuickSort, but makes only 1 recursive
// call instead of 2. See also Quickselect.

// 1 or 0 recursive calls
