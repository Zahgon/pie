package pie

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func Diff[T comparable](ss []T, against []T) (added, removed []T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove duplicates
