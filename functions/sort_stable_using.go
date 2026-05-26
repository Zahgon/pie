package functions

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func (ss SliceType) SortStableUsing(less func(a, b ElementType) bool) SliceType {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(SliceType)
}
