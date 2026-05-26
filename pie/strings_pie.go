package pie

import (
	"context"
	"math/rand"
)

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Strings) All(fn func(value string) bool) bool { _ = "STUB: not implemented"; return false }

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss Strings) Any(fn func(value string) bool) bool { _ = "STUB: not implemented"; return false }

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss Strings) Append(elements ...string) Strings {
	_ = "STUB: not implemented"
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	return *new(Strings)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.StringsAreSorted.
func (ss Strings) AreSorted() bool { _ = "STUB: not implemented"; return false }

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Strings) AreUnique() bool { _ = "STUB: not implemented"; return false }

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Strings) Bottom(n int) (top Strings) { _ = "STUB: not implemented"; return *new(Strings) }

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss Strings) Contains(lookingFor string) bool { _ = "STUB: not implemented"; return false }

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss Strings) Diff(against Strings) (added, removed Strings) {
	_ = "STUB: not implemented"
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.
	return *new(Strings), *new(Strings)
}

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (ss Strings) DropTop(n int) (drop Strings) { _ = "STUB: not implemented"; return *new(Strings) }

// Copy ss, to make sure no memory is overlapping between input and
// output. See issue #145.

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the same logic as the dropwhile() function from itertools in Python.
func (ss Strings) DropWhile(f func(s string) bool) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Each is more condensed version of Transform that allows an action to happen
// on each elements and pass the original slice on.
//
//	cars.Each(func (car *Car) {
//	    fmt.Printf("Car color is: %s\n", car.Color)
//	})
//
// Pie will not ensure immutability on items passed in so they can be
// manipulated, if you choose to do it this way, for example:
//
//	// Set all car colors to Red.
//	cars.Each(func (car *Car) {
//	    car.Color = "Red"
//	})
func (ss Strings) Each(fn func(string)) Strings { _ = "STUB: not implemented"; return *new(Strings) }

// Equals compare elements from the start to the end,
//
// if they are the same is considered the slices are equal if all elements are the same is considered the slices are equal
// if each slice == nil is considered that they're equal
//
// if element realizes Equals interface it uses that method, in other way uses default compare
func (ss Strings) Equals(rhs Strings) bool { _ = "STUB: not implemented"; return false }

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Strings) Extend(slices ...Strings) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss Strings) Filter(condition func(string) bool) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Strings) FilterNot(condition func(string) bool) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (ss Strings) FindFirstUsing(fn func(value string) bool) int {
	_ = "STUB: not implemented"
	return 0
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Strings) First() string { _ = "STUB: not implemented"; return "" }

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Strings) FirstOr(defaultValue string) string { _ = "STUB: not implemented"; return "" }

// Float64s transforms each element to a float64.
func (ss Strings) Float64s() Float64s {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Float64s)
}

// Group returns a map of the value with an individual count.
func (ss Strings) Group() map[string]int { _ = "STUB: not implemented"; return nil }

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss Strings) Intersect(slices ...Strings) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Insert a value at an index
func (ss Strings) Insert(index int, values ...string) Strings {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Ints transforms each element to an integer.
func (ss Strings) Ints() Ints {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Ints)
}

// Join returns a string from joining each of the elements.
func (ss Strings) Join(glue string) (s string) { _ = "STUB: not implemented"; return "" }

// The stdlib is efficient for type []string

// General case

// JSONBytes returns the JSON encoded array as bytes.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Strings) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.

// JSONBytesIndent returns the JSON encoded array as bytes with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss Strings) JSONBytesIndent(prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// An error should not be possible.

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Strings) JSONString() string { _ = "STUB: not implemented"; return "" }

// An error should not be possible.

// JSONStringIndent returns the JSON encoded array as a string with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss Strings) JSONStringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// An error should not be possible.

// Last returns the last element, or zero. Also see LastOr().
func (ss Strings) Last() string { _ = "STUB: not implemented"; return "" }

// LastOr returns the last element or a default value if there are no elements.
func (ss Strings) LastOr(defaultValue string) string { _ = "STUB: not implemented"; return "" }

// Len returns the number of elements.
func (ss Strings) Len() int {
	_ = "STUB: not implemented"

	// Map will return a new slice where each element has been mapped (transformed).
	// The number of elements returned will always be the same as the input.
	//
	// Be careful when using this with slices of pointers. If you modify the input
	// value it will affect the original slice. Be sure to return a new allocated
	// object or deep copy the existing one.
	return 0
}

func (ss Strings) Map(fn func(string) string) (ss2 Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Max is the maximum value, or zero.
func (ss Strings) Max() (max string) { _ = "STUB: not implemented"; return "" }

// Min is the minimum value, or zero.
func (ss Strings) Min() (min string) { _ = "STUB: not implemented"; return "" }

// Mode returns a new slice containing the most frequently occuring values.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless the input slice has zero items.
func (ss Strings) Mode() Strings { _ = "STUB: not implemented"; return *new(Strings) }

// Pop the first element of the slice
//
// Usage Example:
//
//	type knownGreetings []string
//	greetings := knownGreetings{"ciao", "hello", "hola"}
//	for greeting := greetings.Pop(); greeting != nil; greeting = greetings.Pop() {
//	    fmt.Println(*greeting)
//	}
func (ss *Strings) Pop() (popped *string) { _ = "STUB: not implemented"; return nil }

// Random returns a random element by your rand.Source, or zero
func (ss Strings) Random(source rand.Source) string {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return ""
}

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of string if there are no elements in the slice. It will panic if the reducer is nil and the slice has more than one element (required to invoke reduce).
// Otherwise returns result of applying reducer from left to right.
func (ss Strings) Reduce(reducer func(string, string) string) (el string) {
	_ = "STUB: not implemented"
	return ""
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//	ss.Sort().Reverse()
func (ss Strings) Reverse() Strings {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	return *new(Strings)
}

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sended elements if len(this) != len(old) considered func was canceled
func (ss Strings) Send(ctx context.Context, ch chan<- string) Strings {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// SequenceUsing generates slice in range using creator function
//
// There are 3 variations to generate:
//  1. [0, n).
//  2. [min, max).
//  3. [min, max) with step.
//
// if len(params) == 1 considered that will be returned slice between 0 and n,
// where n is the first param, [0, n).
// if len(params) == 2 considered that will be returned slice between min and max,
// where min is the first param, max is the second, [min, max).
// if len(params) > 2 considered that will be returned slice between min and max with step,
// where min is the first param, max is the second, step is the third one, [min, max) with step,
// others params will be ignored
func (ss Strings) SequenceUsing(creator func(int) string, params ...int) Strings {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Shift will return two values: the shifted value and the rest slice.
func (ss Strings) Shift() (string, Strings) { _ = "STUB: not implemented"; return "", *new(Strings) }

// Shuffle returns shuffled slice by your rand.Source
func (ss Strings) Shuffle(source rand.Source) Strings {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(Strings)
}

// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
// the algorithm directly from the go source: src/math/rand/rand.go below,
// with some adjustments:

// Sort works similar to sort.Strings(). However, unlike sort.Strings the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Strings) Sort() Strings {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(Strings)
}

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func (ss Strings) SortStableUsing(less func(a, b string) bool) Strings {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(Strings)
}

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (ss Strings) SortUsing(less func(a, b string) bool) Strings {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(Strings)
}

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//	fmt.Sprintf("%v")
func (ss Strings) Strings() Strings {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Strings)
}

// SubSlice will return the subSlice from start to end(excluded)
//
// Condition 1: If start < 0 or end < 0, nil is returned.
// Condition 2: If start >= end, nil is returned.
// Condition 3: Return all elements that exist in the range provided,
// if start or end is out of bounds, zero items will be placed.
func (ss Strings) SubSlice(start int, end int) (subSlice Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Strings) Top(n int) (top Strings) { _ = "STUB: not implemented"; return *new(Strings) }

// StringsUsing transforms each element to a string.
func (ss Strings) StringsUsing(transform func(string) string) Strings {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Strings)
}

// Unique returns a new slice with all of the unique values.
//
// The items will be returned in a randomized order, even with the same input.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless then input slice has zero items.
//
// A slice with zero elements is considered to be unique.
//
// See AreUnique().
func (ss Strings) Unique() Strings {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	return *new(Strings)
}

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func (ss Strings) Unshift(elements ...string) (unshift Strings) {
	_ = "STUB: not implemented"
	return *new(Strings)
}
