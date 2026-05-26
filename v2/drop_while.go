package pie

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the
// same logic as the dropwhile() function from itertools in Python.
func DropWhile[T comparable](ss []T, f func(s T) bool) (ss2 []T) {
	_ = "STUB: not implemented"
	return nil
}
