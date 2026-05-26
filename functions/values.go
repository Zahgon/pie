package functions

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Values() []ElementType {
	_ = "STUB: not implemented"
	// Avoid allocation
	return nil
}
