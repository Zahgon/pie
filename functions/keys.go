package functions

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Keys() KeySliceType {
	_ = "STUB: not implemented"
	// Avoid allocation
	return *new(KeySliceType)
}
