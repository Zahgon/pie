package pie

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func DropTop[T any](ss []T, n int) (drop []T) { _ = "STUB: not implemented"; return nil }

// Copy ss, to make sure no memory is overlapping between input and
// output. See issue #145.
