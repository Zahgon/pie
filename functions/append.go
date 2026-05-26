package functions

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) SliceType {
	_ = "STUB: not implemented"
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	return *new(SliceType)
}
