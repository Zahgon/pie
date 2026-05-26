package pie

import (
	"context"
	"math/rand"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss Float64s) Abs() Float64s { _ = "STUB: not implemented"; return *new(Float64s) }

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Float64s) All(fn func(value float64) bool) bool { _ = "STUB: not implemented"; return false }

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss Float64s) Any(fn func(value float64) bool) bool { _ = "STUB: not implemented"; return false }

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Append(elements ...float64) Float64s {
	_ = "STUB: not implemented"
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	return *new(Float64s)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.Float64sAreSorted.
func (ss Float64s) AreSorted() bool { _ = "STUB: not implemented"; return false }

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Float64s) AreUnique() bool { _ = "STUB: not implemented"; return false }

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Float64s) Average() float64 { _ = "STUB: not implemented"; return 0 }

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Float64s) Bottom(n int) (top Float64s) { _ = "STUB: not implemented"; return *new(Float64s) }

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss Float64s) Contains(lookingFor float64) bool { _ = "STUB: not implemented"; return false }

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss Float64s) Diff(against Float64s) (added, removed Float64s) {
	_ = "STUB: not implemented"
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.
	return *new(Float64s), *new(Float64s)
}

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (ss Float64s) DropTop(n int) (drop Float64s) { _ = "STUB: not implemented"; return *new(Float64s) }

// Copy ss, to make sure no memory is overlapping between input and
// output. See issue #145.

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the same logic as the dropwhile() function from itertools in Python.
func (ss Float64s) DropWhile(f func(s float64) bool) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
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
func (ss Float64s) Each(fn func(float64)) Float64s {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Equals compare elements from the start to the end,
//
// if they are the same is considered the slices are equal if all elements are the same is considered the slices are equal
// if each slice == nil is considered that they're equal
//
// if element realizes Equals interface it uses that method, in other way uses default compare
func (ss Float64s) Equals(rhs Float64s) bool { _ = "STUB: not implemented"; return false }

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Extend(slices ...Float64s) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss Float64s) Filter(condition func(float64) bool) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Float64s) FilterNot(condition func(float64) bool) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (ss Float64s) FindFirstUsing(fn func(value float64) bool) int {
	_ = "STUB: not implemented"
	return 0
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Float64s) First() float64 { _ = "STUB: not implemented"; return 0 }

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Float64s) FirstOr(defaultValue float64) float64 { _ = "STUB: not implemented"; return 0 }

// Float64s transforms each element to a float64.
func (ss Float64s) Float64s() Float64s {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Float64s)
}

// Group returns a map of the value with an individual count.
func (ss Float64s) Group() map[float64]int { _ = "STUB: not implemented"; return nil }

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss Float64s) Intersect(slices ...Float64s) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Insert a value at an index
func (ss Float64s) Insert(index int, values ...float64) Float64s {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Ints transforms each element to an integer.
func (ss Float64s) Ints() Ints {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Ints)
}

// Join returns a string from joining each of the elements.
func (ss Float64s) Join(glue string) (s string) { _ = "STUB: not implemented"; return "" }

// The stdlib is efficient for type []string

// General case

// JSONBytes returns the JSON encoded array as bytes.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Float64s) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.

// JSONBytesIndent returns the JSON encoded array as bytes with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss Float64s) JSONBytesIndent(prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// An error should not be possible.

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Float64s) JSONString() string { _ = "STUB: not implemented"; return "" }

// An error should not be possible.

// JSONStringIndent returns the JSON encoded array as a string with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss Float64s) JSONStringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// An error should not be possible.

// Last returns the last element, or zero. Also see LastOr().
func (ss Float64s) Last() float64 { _ = "STUB: not implemented"; return 0 }

// LastOr returns the last element or a default value if there are no elements.
func (ss Float64s) LastOr(defaultValue float64) float64 { _ = "STUB: not implemented"; return 0 }

// Len returns the number of elements.
func (ss Float64s) Len() int {
	_ = "STUB: not implemented"

	// Map will return a new slice where each element has been mapped (transformed).
	// The number of elements returned will always be the same as the input.
	//
	// Be careful when using this with slices of pointers. If you modify the input
	// value it will affect the original slice. Be sure to return a new allocated
	// object or deep copy the existing one.
	return 0
}

func (ss Float64s) Map(fn func(float64) float64) (ss2 Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Max is the maximum value, or zero.
func (ss Float64s) Max() (max float64) { _ = "STUB: not implemented"; return 0 }

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
//
// If the number of elements is even, then the float64 mean of the two "median values"
// is returned.
func (ss Float64s) Median() float64 { _ = "STUB: not implemented"; return 0 }

// This implementation aims at linear time O(n) on average.
// It uses the same idea as QuickSort, but makes only 1 recursive
// call instead of 2. See also Quickselect.

// 1 or 0 recursive calls

// Min is the minimum value, or zero.
func (ss Float64s) Min() (min float64) { _ = "STUB: not implemented"; return 0 }

// Mode returns a new slice containing the most frequently occuring values.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless the input slice has zero items.
func (ss Float64s) Mode() Float64s { _ = "STUB: not implemented"; return *new(Float64s) }

// Pop the first element of the slice
//
// Usage Example:
//
//	type knownGreetings []string
//	greetings := knownGreetings{"ciao", "hello", "hola"}
//	for greeting := greetings.Pop(); greeting != nil; greeting = greetings.Pop() {
//	    fmt.Println(*greeting)
//	}
func (ss *Float64s) Pop() (popped *float64) { _ = "STUB: not implemented"; return nil }

// Product is the product of all of the elements.
func (ss Float64s) Product() (product float64) { _ = "STUB: not implemented"; return 0 }

// Random returns a random element by your rand.Source, or zero
func (ss Float64s) Random(source rand.Source) float64 {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return 0
}

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of float64 if there are no elements in the slice. It will panic if the reducer is nil and the slice has more than one element (required to invoke reduce).
// Otherwise returns result of applying reducer from left to right.
func (ss Float64s) Reduce(reducer func(float64, float64) float64) (el float64) {
	_ = "STUB: not implemented"
	return 0
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//	ss.Sort().Reverse()
func (ss Float64s) Reverse() Float64s {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	return *new(Float64s)
}

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sended elements if len(this) != len(old) considered func was canceled
func (ss Float64s) Send(ctx context.Context, ch chan<- float64) Float64s {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Sequence generates all numbers in range or returns nil if params invalid
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
func (ss Float64s) Sequence(params ...int) Float64s {
	_ = "STUB: not implemented"
	return *new(Float64s)
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
func (ss Float64s) SequenceUsing(creator func(int) float64, params ...int) Float64s {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Shift will return two values: the shifted value and the rest slice.
func (ss Float64s) Shift() (float64, Float64s) { _ = "STUB: not implemented"; return 0, *new(Float64s) }

// Shuffle returns shuffled slice by your rand.Source
func (ss Float64s) Shuffle(source rand.Source) Float64s {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(Float64s)
}

// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
// the algorithm directly from the go source: src/math/rand/rand.go below,
// with some adjustments:

// Sort works similar to sort.Float64s(). However, unlike sort.Float64s the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Float64s) Sort() Float64s {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(Float64s)
}

// Stddev is the standard deviation
func (ss Float64s) Stddev() float64 { _ = "STUB: not implemented"; return 0 }

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//	fmt.Sprintf("%v")
func (ss Float64s) Strings() Strings {
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
func (ss Float64s) SubSlice(start int, end int) (subSlice Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}

// Sum is the sum of all of the elements.
func (ss Float64s) Sum() (sum float64) { _ = "STUB: not implemented"; return 0 }

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Float64s) Top(n int) (top Float64s) { _ = "STUB: not implemented"; return *new(Float64s) }

// StringsUsing transforms each element to a string.
func (ss Float64s) StringsUsing(transform func(float64) string) Strings {
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
func (ss Float64s) Unique() Float64s {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	return *new(Float64s)
}

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func (ss Float64s) Unshift(elements ...float64) (unshift Float64s) {
	_ = "STUB: not implemented"
	return *new(Float64s)
}
