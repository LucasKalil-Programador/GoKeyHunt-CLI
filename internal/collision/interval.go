package collision

import (
	"fmt"
	"math/big"
)

// Interval represents a range with a start and end value.
type Interval struct {
	a, b *big.Int
}

// Set initializes the interval with two big integers, ensuring that the start is less than or equal to the end.
// If `a` is greater than `b`, their values are swapped.
//
// Parameters:
// - a: The start value of the interval.
// - b: The end value of the interval.
//
// Returns:
// - *Interval: The updated interval.
func (interval *Interval) Set(a, b *big.Int) *Interval {
	if a.Cmp(b) > 0 {
		a, b = b, a
	}
	interval.a, interval.b = new(big.Int).Set(a), new(big.Int).Set(b)
	return interval
}

// SetInt initializes the interval with two integers, ensuring that the start is less than or equal to the end.
// It converts the integers to big integers before setting the interval.
//
// Parameters:
// - a: The start value of the interval (as an int).
// - b: The end value of the interval (as an int).
//
// Returns:
// - *Interval: The updated interval.
func (interval *Interval) SetInt(a, b int) *Interval {
	return interval.SetInt64(int64(a), int64(b))
}

// SetInt64 initializes the interval with two int64 values, ensuring that the start is less than or equal to the end.
// It converts the int64 values to big integers before setting the interval.
//
// Parameters:
// - a: The start value of the interval (as an int64).
// - b: The end value of the interval (as an int64).
//
// Returns:
// - *Interval: The updated interval.
func (interval *Interval) SetInt64(a, b int64) *Interval {
	if a > b {
		a, b = b, a
	}
	interval.a = big.NewInt(a)
	interval.b = big.NewInt(b)
	return interval
}

// SetString initializes the interval with two string values, interpreting them as integers in the specified base.
// It ensures that the start is less than or equal to the end. If the conversion fails, it returns false.
//
// Parameters:
// - a: The start value of the interval (as a string).
// - b: The end value of the interval (as a string).
// - base: The numeric base for interpreting the string values.
//
// Returns:
// - *Interval: The updated interval.
// - bool: True if the conversion was successful for both values, false otherwise.
func (interval *Interval) SetString(a, b string, base int) (*Interval, bool) {
	newA, successA := new(big.Int).SetString(a, base)
	newB, successB := new(big.Int).SetString(b, base)
	if newA.Cmp(newB) > 0 {
		newA, newB = newB, newA
	}
	interval.a, interval.b = newA, newB
	return interval, successA && successB
}

// Equals checks if the interval is equal to another interval.
//
// Parameters:
// - other: The interval to compare with.
//
// Returns:
// - bool: True if the intervals are equal, false otherwise.
func (i Interval) Equals(other Interval) bool {
	return i.a.Cmp(other.a) == 0 && i.b.Cmp(other.b) == 0
}

// String returns a string representation of the interval in the format "Interval[Start: <start>, End: <end>]".
//
// Returns:
// - string: The string representation of the interval.
func (i Interval) String() string {
	return fmt.Sprintf("Interval[Start: %s, End: %s]", i.a.String(), i.b.String())
}

// Clone creates a new interval with the same start and end values as the current interval.
//
// Returns:
// - *Interval: The cloned interval.
func (i *Interval) Clone() *Interval {
	clone := &Interval{
		a: new(big.Int).Set(i.a),
		b: new(big.Int).Set(i.b),
	}
	return clone
}

// IsPointOverlap checks if a given point is within the interval (inclusive).
//
// Parameters:
// - point: The point to check.
//
// Returns:
// - bool: True if the point is within the interval, false otherwise.
func (i *Interval) IsPointOverlap(point *big.Int) bool {
	return i.a.Cmp(point) <= 0 && i.b.Cmp(point) >= 0
}

// IsOverlap checks if the interval overlaps with another interval.
//
// Parameters:
// - i2: The interval to check for overlap with.
//
// Returns:
// - bool: True if there is an overlap, false otherwise.
func (i1 *Interval) IsOverlap(i2 *Interval) bool {
	return i1.IsPointOverlap(i2.a) || i1.IsPointOverlap(i2.b) || i2.IsPointOverlap(i1.a) || i2.IsPointOverlap(i1.b)
}

// Get returns the start and end values of the interval.
//
// Returns:
// - *big.Int: The start value of the interval.
// - *big.Int: The end value of the interval.
func (i1 *Interval) Get() (*big.Int, *big.Int) {
	return i1.a, i1.b
}
