package functions

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss SliceType) Diff(against SliceType) (added, removed SliceType) {
	_ = "STUB: not implemented"
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.
	return *new(SliceType), *new(SliceType)
}
