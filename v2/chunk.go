package pie

// Chunk splits the input and returns multi slices whose length equals chunkLength,
// except for the last slice which may contain fewer elements.
//
// Examples:
//
//	Chunk([1, 2, 3], 4) => [ [1, 2, 3] ]
//	Chunk([1, 2, 3], 3) => [ [1, 2, 3] ]
//	Chunk([1, 2, 3], 2) => [ [1, 2], [3] ]
//	Chunk([1, 2, 3], 1) => [ [1], [2], [3] ]
//	Chunk([], 1)        => [ [] ]
//	Chunk([1, 2, 3], 0) => panic: chunkLength should be greater than 0
func Chunk[T any](ss []T, chunkLength int) [][]T { _ = "STUB: not implemented"; return nil }
