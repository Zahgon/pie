package pie

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func Filter[T any](ss []T, condition func(T) bool) (ss2 []T) { _ = "STUB: not implemented"; return nil }
