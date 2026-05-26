package pie

import (
	"context"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// OfNumeric encapsulates a slice to be used in multiple chained operations.
// OfNumeric requires that elements be numerical for certain operations to be
// performed.
func OfNumeric[T constraints.Integer | constraints.Float](ss []T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

// OfNumericSlice provides the proxy methods that operate on slices. If the last
// method in the chain does not return a single value, you can access the Result
// to get final slice.
type OfNumericSlice[T constraints.Integer | constraints.Float] struct {
	Result []T
}

func (o OfNumericSlice[T]) All(fn func(value T) bool) bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) Any(fn func(value T) bool) bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) AreSorted() bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) AreUnique() bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) Average() float64 { _ = "STUB: not implemented"; return 0 }

func (o OfNumericSlice[T]) Bottom(n int) OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Contains(lookingFor T) bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) Diff(against []T) ([]T, []T) { _ = "STUB: not implemented"; return nil, nil }

func (o OfNumericSlice[T]) DropTop(n int) OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) DropWhile(f func(s T) bool) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Each(fn func(T)) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Equals(rhs []T) bool { _ = "STUB: not implemented"; return false }

func (o OfNumericSlice[T]) Filter(condition func(T) bool) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) FilterNot(condition func(T) bool) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (o OfNumericSlice[T]) First() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) FirstOr(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Float64s() []float64 { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Group() map[T]int { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Insert(index int, values ...T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Intersect(slices ...[]T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Ints() []int { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Join(glue string) string { _ = "STUB: not implemented"; return "" }

func (o OfNumericSlice[T]) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) JSONBytesIndent(prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) JSONString() string { _ = "STUB: not implemented"; return "" }

func (o OfNumericSlice[T]) JSONStringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

func (o OfNumericSlice[T]) Last() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) LastOr(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Map(fn func(T) T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Max() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Median() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Min() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Mode() OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Product() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Random(source rand.Source) T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Reduce(reducer func(T, T) T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func (o OfNumericSlice[T]) Reverse() OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Send(ctx context.Context, ch chan<- T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Sequence(params ...int) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Shuffle(source rand.Source) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Sort() OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) SortStableUsing(less func(a, b T) bool) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) SortUsing(less func(a, b T) bool) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Stddev() float64 { _ = "STUB: not implemented"; return 0 }

func (o OfNumericSlice[T]) Strings() []string { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) StringsUsing(transform func(T) string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) SubSlice(start int, end int) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Sum() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfNumericSlice[T]) Top(n int) OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Unique() OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) UniqueStable() OfNumericSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfNumericSlice[T]) Unshift(elements ...T) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfNumericSlice[T]) Delete(idx ...int) OfNumericSlice[T] {
	_ = "STUB: not implemented"
	return nil
}
