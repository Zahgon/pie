package pie

import (
	"golang.org/x/exp/constraints"
)

// JSONBytesIndent returns the JSON encoded array as bytes with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func JSONBytesIndent[T constraints.Ordered](ss []T, prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// An error should not be possible.
