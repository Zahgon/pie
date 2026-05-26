package functions

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (ss SliceType) SortUsing(less func(a, b ElementType) bool) SliceType {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(SliceType)
}
