package functions

// JSONBytes returns the JSON encoded array as bytes.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss SliceType) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.
