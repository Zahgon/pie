package pie

import (
	"golang.org/x/exp/constraints"
)

// JSONBytes returns the JSON encoded array as bytes.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func JSONBytes[T constraints.Ordered](ss []T) []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.
