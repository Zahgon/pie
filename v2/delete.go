package pie

// Removes elements at indices in idx from input slice, returns resulting slice.
// If an index is out of bounds, skip it.
func Delete[T any](ss []T, idx ...int) []T {
	_ = "STUB: not implemented"
	// short path O(n) inplace
	return nil
}

// long path O(mLog(m) + n)

// Copy by consecutive chunks instead of one by one
