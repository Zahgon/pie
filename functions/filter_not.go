package functions

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss SliceType) FilterNot(condition func(ElementType) bool) (ss2 SliceType) {
	_ = "STUB: not implemented"
	return *new(SliceType)
}
