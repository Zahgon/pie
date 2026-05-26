package pie

// Rotate return slice circularly rotated by a number of positions n.
// If n is positive, the slice is rotated right.
// If n is negative, the slice is rotated left.
func Rotate[T any](ss []T, n int) []T {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	// If there is one element or less, then already rotated.
	return nil
}

// Normalize shift
// no div by 0 since length >= 2

// Avoid the allocation.
// If normalized shift is 0, then already rotated.
