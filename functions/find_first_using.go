package functions

// FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (ss SliceType) FindFirstUsing(fn func(value ElementType) bool) int {
	_ = "STUB: not implemented"
	return 0
}
