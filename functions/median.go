package functions

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
//
// If the number of elements is even, then the ElementType mean of the two "median values"
// is returned.
func (ss SliceType) Median() ElementType { _ = "STUB: not implemented"; return *new(ElementType) }

// This implementation aims at linear time O(n) on average.
// It uses the same idea as QuickSort, but makes only 1 recursive
// call instead of 2. See also Quickselect.

// 1 or 0 recursive calls
