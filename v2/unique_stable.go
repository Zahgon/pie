package pie

// UniqueStable works similar to Unique. However, unlike Unique
// the slice returned will be in previous relative order
func UniqueStable[T comparable](ss []T) []T {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	return nil
}
