package functions

// JSONStringIndent returns the JSON encoded array as a string with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss SliceType) JSONStringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// An error should not be possible.
