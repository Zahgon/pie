package pie

// SubSlice will return the subSlice from start to end(excluded)
//
// Condition 1: If start < 0 or end < 0, nil is returned.
// Condition 2: If start >= end, nil is returned.
// Condition 3: Return all elements that exist in the range provided,
// if start or end is out of bounds, zero items will be placed.
func SubSlice[T any](ss []T, start int, end int) (subSlice []T) {
	_ = "STUB: not implemented"
	return nil
}
