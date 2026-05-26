package functions

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (ss SliceType) DropTop(n int) (drop SliceType) {
	_ = "STUB: not implemented"
	return *new(SliceType)
}

// Copy ss, to make sure no memory is overlapping between input and
// output. See issue #145.
