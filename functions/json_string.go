package functions

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss SliceType) JSONString() string { _ = "STUB: not implemented"; return "" }

// An error should not be possible.
