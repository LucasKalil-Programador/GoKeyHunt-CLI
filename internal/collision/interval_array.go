package collision

import (
	"fmt"
	"math/big"
)

// IntervalArray represents a collection of intervals.
type IntervalArray struct {
	data []Interval
}

// String returns a string representation of the IntervalArray.
// It lists all intervals in the array as a formatted string.
func (ia *IntervalArray) String() string {
	intervalStrings := make([]string, len(ia.data))
	for i, interval := range ia.data {
		intervalStrings[i] = interval.String()
	}

	return fmt.Sprintf("[%s]", intervalStrings)
}

// Size returns the number of intervals in the IntervalArray.
//
// Returns:
// - int: The size of the IntervalArray.
func (ia *IntervalArray) Size() int {
	return len(ia.data)
}

// CalculateTotalProgress computes the total progress represented by the intervals.
// It sums up the difference between the end and start of each interval, adding one to each interval.
//
// Returns:
// - *big.Int: The total progress.
func (ia *IntervalArray) CalculateTotalProgress() *big.Int {
	total, one := new(big.Int), big.NewInt(1)
	for _, interval := range ia.data {
		total.Add(total, interval.b).Sub(total, interval.a).Add(total, one)
	}
	return total
}

// NewIntervalArray creates a new IntervalArray from a slice of intervals.
// It sorts the intervals by their start values.
//
// Parameters:
// - intervals: A slice of Interval to initialize the IntervalArray.
//
// Returns:
// - *IntervalArray: The newly created IntervalArray.
func NewIntervalArray(intervals []Interval) *IntervalArray {
	newIntervalsArr := make([]Interval, len(intervals))
	copy(newIntervalsArr, intervals)
	SortByStart(newIntervalsArr)
	return &IntervalArray{data: newIntervalsArr}
}

// NewEmptyIntervalArray creates a new empty IntervalArray.
//
// Returns:
// - *IntervalArray: The newly created empty IntervalArray.
func NewEmptyIntervalArray() *IntervalArray {
	return &IntervalArray{}
}

// Append adds a new interval to the IntervalArray, inserting it in sorted order.
// It uses the InsertSorted function to ensure the intervals remain ordered.
//
// Parameters:
// - interval: The Interval to be added.
func (interArray *IntervalArray) Append(interval *Interval) {
	interArray.data = InsertSorted(interArray.data, *interval)
}

// ResolveCollisions resolves collisions for a given target interval by adjusting its start and end values.
// It attempts to find a non-overlapping interval and returns whether the adjustment was valid.
//
// Parameters:
// - target: The Interval to resolve collisions for.
//
// Returns:
// - *Interval: The adjusted interval.
// - bool: True if the adjustment was valid, false otherwise.
func (interArray *IntervalArray) ResolveCollisions(target Interval) (*Interval, bool) {
	newInterval := target.Clone()
	for _, interval := range interArray.data {
		if interval.IsPointOverlap(newInterval.a) {
			newInterval.a = new(big.Int).Add(interval.b, big.NewInt(1))
		} else {
			newInterval.b = new(big.Int).Sub(interval.a, big.NewInt(1))
			break
		}
	}
	isValid := newInterval.a.Cmp(newInterval.b) <= 0
	return newInterval, isValid
}

// HandleIntervalCollision checks for collisions with a given interval and resolves them if necessary.
// It attempts to resolve the collision by adjusting the interval and returns whether a collision was resolved.
//
// Parameters:
// - interval: The Interval to check and handle collisions for.
//
// Returns:
// - bool: True if a collision was resolved, false otherwise.
// - Interval: The resulting interval after handling collisions.
func (interArray *IntervalArray) HandleIntervalCollision(interval Interval) (bool, Interval) {
	hasCollision := HasOverlap(&interval, interArray.data)
	if hasCollision {
		intervals := GetIntervalsBetween(&interval, interArray.data)
		newInterval, valid := NewIntervalArray(intervals).ResolveCollisions(interval)
		if valid {
			newHasCollision := HasOverlap(newInterval, interArray.data)
			if !newHasCollision {
				return newHasCollision, *newInterval
			}
		}
	}
	return hasCollision, interval
}

// Optimize merges overlapping intervals to reduce the total number of intervals.
// It returns the number of intervals removed during the optimization.
//
// Returns:
// - int: The number of intervals removed.
func (interArray *IntervalArray) Optimize() int {
	length := len(interArray.data)
	intervals := interArray.data
	rmCount := 0

	i := 0
	for i < length-1 {
		interval1, interval2 := intervals[i], intervals[i+1]
		if new(big.Int).Sub(interval2.a, interval1.b).Cmp(big.NewInt(1)) <= 0 {
			intervals[i] = *new(Interval).Set(interval1.a, maxBigInt(interval1.b, interval2.b))
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			rmCount++
			length--
		} else {
			i++
		}
	}

	interArray.data = intervals
	return rmCount
}

// maxBigInt returns the maximum of two big integers.
//
// Parameters:
// - a: The first big integer.
// - b: The second big integer.
//
// Returns:
// - *big.Int: The maximum of the two big integers.
func maxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}
