package pie

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//	ss.Sort().Reverse()
func Reverse[T any](ss []T) []T {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	return nil
}
