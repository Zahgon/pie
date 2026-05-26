package functions

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss SliceType) Filter(condition func(ElementType) bool) (ss2 SliceType) {
	_ = "STUB: not implemented"
	return *new(SliceType)
}
