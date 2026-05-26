package util

// These constants was copied from src/math/bits.go to support Round in go
// versions before 1.10.
const (
	shift    = 64 - 11 - 1
	mask     = 0x7FF
	bias     = 1023
	signMask = 1 << 63
	uvone    = 0x3FF0000000000000
	fracMask = 1<<shift - 1
)

// Round was copied from src/math/floor.go to support Round in go
// versions before 1.10.
func Round(x float64) float64 {
	_ = "STUB: not implemented"
	// Round is a faster implementation of:
	//
	//	func Round(x float64) float64 {
	//	  t := Trunc(x)
	//	  if Abs(x-t) >= 0.5 {
	//	    return t + Copysign(1, x)
	//	  }
	//	  return t
	//	}
	return 0
}

// Round abs(x) < 1 including denormals.
// +-0

// +-1

// Round any abs(x) >= 1 containing a fractional component [0,1).
//
// Numbers with larger exponents are returned unchanged since they
// must be either an integer, infinity, or NaN.
