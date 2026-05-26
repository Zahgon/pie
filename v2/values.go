package pie

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func Values[K comparable, V any](m map[K]V) []V {
	_ = "STUB: not implemented"
	// Avoid allocation
	return nil
}
