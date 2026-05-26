package pie

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func SortStableUsing[T comparable](ss []T, less func(a, b T) bool) []T {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return nil
}
