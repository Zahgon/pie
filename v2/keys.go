package pie

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func Keys[K comparable, V any](m map[K]V) []K {
	_ = "STUB: not implemented"
	// Avoid allocation
	return nil
}
