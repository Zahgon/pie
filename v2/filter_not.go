package pie

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func FilterNot[T any](ss []T, condition func(T) bool) (ss2 []T) {
	_ = "STUB: not implemented"
	return nil
}
