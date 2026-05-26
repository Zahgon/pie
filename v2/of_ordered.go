package pie

import (
	"context"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// OfOrdered encapsulates a slice to be used in multiple chained operations.
// OfOrdered requires that elements be numerical or a string for certain
// operations to be performed.
func OfOrdered[T constraints.Ordered](ss []T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

// OfOrderedSlice provides the proxy methods that operate on slices. If the last
// method in the chain does not return a single value, you can access the Result
// to get final slice.
type OfOrderedSlice[T constraints.Ordered] struct {
	Result []T
}

func (o OfOrderedSlice[T]) All(fn func(value T) bool) bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) Any(fn func(value T) bool) bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) AreSorted() bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) AreUnique() bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) Bottom(n int) OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Contains(lookingFor T) bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) Diff(against []T) ([]T, []T) { _ = "STUB: not implemented"; return nil, nil }

func (o OfOrderedSlice[T]) DropTop(n int) OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) DropWhile(f func(s T) bool) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Each(fn func(T)) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Equals(rhs []T) bool { _ = "STUB: not implemented"; return false }

func (o OfOrderedSlice[T]) Filter(condition func(T) bool) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) FilterNot(condition func(T) bool) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	_ = "STUB: not implemented"
	return 0
}

func (o OfOrderedSlice[T]) First() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) FirstOr(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) Float64s() []float64 { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Group() map[T]int { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Insert(index int, values ...T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Intersect(slices ...[]T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Ints() []int { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Join(glue string) string { _ = "STUB: not implemented"; return "" }

func (o OfOrderedSlice[T]) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) JSONBytesIndent(prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) JSONString() string { _ = "STUB: not implemented"; return "" }

func (o OfOrderedSlice[T]) JSONStringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

func (o OfOrderedSlice[T]) Last() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) LastOr(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) Map(fn func(T) T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Max() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) Min() T { _ = "STUB: not implemented"; return *new(T) }

func (o OfOrderedSlice[T]) Mode() OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Reverse() OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Send(ctx context.Context, ch chan<- T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Shuffle(source rand.Source) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Sort() OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) SortStableUsing(less func(a, b T) bool) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) SortUsing(less func(a, b T) bool) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Strings() []string { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) StringsUsing(transform func(T) string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) SubSlice(start int, end int) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Top(n int) OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Unique() OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) UniqueStable() OfOrderedSlice[T] { _ = "STUB: not implemented"; return nil }

func (o OfOrderedSlice[T]) Unshift(elements ...T) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}

func (o OfOrderedSlice[T]) Delete(idx ...int) OfOrderedSlice[T] {
	_ = "STUB: not implemented"
	return nil
}
