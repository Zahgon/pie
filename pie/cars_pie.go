package pie

import (
	"context"
	"math/rand"
)

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss cars) All(fn func(value car) bool) bool { _ = "STUB: not implemented"; return false }

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss cars) Any(fn func(value car) bool) bool { _ = "STUB: not implemented"; return false }

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss cars) Append(elements ...car) cars {
	_ = "STUB: not implemented"
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	return *new(cars)
}

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss cars) Bottom(n int) (top cars) { _ = "STUB: not implemented"; return *new(cars) }

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss cars) Contains(lookingFor car) bool { _ = "STUB: not implemented"; return false }

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss cars) Diff(against cars) (added, removed cars) {
	_ = "STUB: not implemented"
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.
	return *new(cars), *new(cars)
}

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (ss cars) DropTop(n int) (drop cars) { _ = "STUB: not implemented"; return *new(cars) }

// Copy ss, to make sure no memory is overlapping between input and
// output. See issue #145.

// Drop items from the slice while f(item) is true.
// Afterwards, return every element until the slice is empty. It follows the same logic as the dropwhile() function from itertools in Python.
func (ss cars) DropWhile(f func(s car) bool) (ss2 cars) {
	_ = "STUB: not implemented"
	return *new(cars)
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
func (ss cars) Each(fn func(car)) cars { _ = "STUB: not implemented"; return *new(cars) }

// Equals compare elements from the start to the end,
//
// if they are the same is considered the slices are equal if all elements are the same is considered the slices are equal
// if each slice == nil is considered that they're equal
//
// if element realizes Equals interface it uses that method, in other way uses default compare
func (ss cars) Equals(rhs cars) bool { _ = "STUB: not implemented"; return false }

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss cars) Extend(slices ...cars) (ss2 cars) { _ = "STUB: not implemented"; return *new(cars) }

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss cars) Filter(condition func(car) bool) (ss2 cars) {
	_ = "STUB: not implemented"
	return *new(cars)
}

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss cars) FilterNot(condition func(car) bool) (ss2 cars) {
	_ = "STUB: not implemented"
	return *new(cars)
}

// FindFirstUsing will return the index of the first element when the callback returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (ss cars) FindFirstUsing(fn func(value car) bool) int { _ = "STUB: not implemented"; return 0 }

// First returns the first element, or zero. Also see FirstOr().
func (ss cars) First() car {
	_ = "STUB: not implemented"
	return *

	// FirstOr returns the first element or a default value if there are no
	// elements.
	new(car)
}

func (ss cars) FirstOr(defaultValue car) car { _ = "STUB: not implemented"; return *new(car) }

// Float64s transforms each element to a float64.
func (ss cars) Float64s() Float64s {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Float64s)
}

// Insert a value at an index
func (ss cars) Insert(index int, values ...car) cars { _ = "STUB: not implemented"; return *new(cars) }

// Ints transforms each element to an integer.
func (ss cars) Ints() Ints {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Ints)
}

// Join returns a string from joining each of the elements.
func (ss cars) Join(glue string) (s string) { _ = "STUB: not implemented"; return "" }

// The stdlib is efficient for type []string

// General case

// JSONBytes returns the JSON encoded array as bytes.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss cars) JSONBytes() []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.

// JSONBytesIndent returns the JSON encoded array as bytes with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss cars) JSONBytesIndent(prefix, indent string) []byte { _ = "STUB: not implemented"; return nil }

// An error should not be possible.

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss cars) JSONString() string { _ = "STUB: not implemented"; return "" }

// An error should not be possible.

// JSONStringIndent returns the JSON encoded array as a string with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func (ss cars) JSONStringIndent(prefix, indent string) string { _ = "STUB: not implemented"; return "" }

// An error should not be possible.

// Last returns the last element, or zero. Also see LastOr().
func (ss cars) Last() car {
	_ = "STUB: not implemented"
	return *

	// LastOr returns the last element or a default value if there are no elements.
	new(car)
}

func (ss cars) LastOr(defaultValue car) car { _ = "STUB: not implemented"; return *new(car) }

// Len returns the number of elements.
func (ss cars) Len() int {
	_ = "STUB: not implemented"

	// Map will return a new slice where each element has been mapped (transformed).
	// The number of elements returned will always be the same as the input.
	//
	// Be careful when using this with slices of pointers. If you modify the input
	// value it will affect the original slice. Be sure to return a new allocated
	// object or deep copy the existing one.
	return 0
}

func (ss cars) Map(fn func(car) car) (ss2 cars) { _ = "STUB: not implemented"; return *new(cars) }

// Mode returns a new slice containing the most frequently occuring values.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless the input slice has zero items.
func (ss cars) Mode() cars { _ = "STUB: not implemented"; return *new(cars) }

// Pop the first element of the slice
//
// Usage Example:
//
//	type knownGreetings []string
//	greetings := knownGreetings{"ciao", "hello", "hola"}
//	for greeting := greetings.Pop(); greeting != nil; greeting = greetings.Pop() {
//	    fmt.Println(*greeting)
//	}
func (ss *cars) Pop() (popped *car) { _ = "STUB: not implemented"; return nil }

// Random returns a random element by your rand.Source, or zero
func (ss cars) Random(source rand.Source) car {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(car)
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//	ss.Sort().Reverse()
func (ss cars) Reverse() cars {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	return *new(cars)
}

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sended elements if len(this) != len(old) considered func was canceled
func (ss cars) Send(ctx context.Context, ch chan<- car) cars {
	_ = "STUB: not implemented"
	return *new(cars)
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
func (ss cars) SequenceUsing(creator func(int) car, params ...int) cars {
	_ = "STUB: not implemented"
	return *new(cars)
}

// Shift will return two values: the shifted value and the rest slice.
func (ss cars) Shift() (car, cars) { _ = "STUB: not implemented"; return *new(car), *new(cars) }

// Shuffle returns shuffled slice by your rand.Source
func (ss cars) Shuffle(source rand.Source) cars {
	_ = "STUB: not implemented"

	// Avoid the extra allocation.
	return *new(cars)
}

// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
// the algorithm directly from the go source: src/math/rand/rand.go below,
// with some adjustments:

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func (ss cars) SortStableUsing(less func(a, b car) bool) cars {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(cars)
}

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (ss cars) SortUsing(less func(a, b car) bool) cars {
	_ = "STUB: not implemented"
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	return *new(cars)
}

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//	fmt.Sprintf("%v")
func (ss cars) Strings() Strings {
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
func (ss cars) SubSlice(start int, end int) (subSlice cars) {
	_ = "STUB: not implemented"
	return *new(cars)
}

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss cars) Top(n int) (top cars) { _ = "STUB: not implemented"; return *new(cars) }

// StringsUsing transforms each element to a string.
func (ss cars) StringsUsing(transform func(car) string) Strings {
	_ = "STUB: not implemented"

	// Avoid the allocation.
	return *new(Strings)
}

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func (ss cars) Unshift(elements ...car) (unshift cars) {
	_ = "STUB: not implemented"
	return *new(cars)
}
