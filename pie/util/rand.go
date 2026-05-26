package util

import "math/rand"

// Int31n was copied from src/math/rand/rand.go to support Shuffle in go
// versions before 1.10.
func Int31n(r *rand.Rand, n int32) int32 { _ = "STUB: not implemented"; return 0 }

// Shuffle was copied from src/math/rand/rand.go to support Shuffle in go
// versions before 1.10.
func Shuffle(r *rand.Rand, n int, swap func(i, j int)) { _ = "STUB: not implemented"; return }

// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
// Shuffle really ought not be called with n that doesn't fit in 32 bits.
// Not only will it take a very long time, but with 2³¹! possible permutations,
// there's no way that any PRNG can have a big enough internal state to
// generate even a minuscule percentage of the possible permutations.
// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
