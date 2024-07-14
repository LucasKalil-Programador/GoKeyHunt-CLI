package collision

import (
	"sort"
)

// Define a type that implements sort.Interface
type ByStart []Interval

// Len returns the length of the slice of intervals.
func (a ByStart) Len() int { return len(a) }

// Swap exchanges the intervals at positions i and j.
func (a ByStart) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares intervals at positions i and j and returns true if the interval at i starts before the interval at j.
func (a ByStart) Less(i, j int) bool { return a[i].a.Cmp(a[j].a) < 0 }

// SortByStart sorts a slice of intervals by the start field.
//
// Parameters:
// - intervals: A slice of Intervals to be sorted.
//
// Returns:
// - []Interval: The sorted slice of intervals.
func SortByStart(intervals []Interval) []Interval {
	sort.Sort(ByStart(intervals))
	return intervals
}

// HasOverlap checks if an interval overlaps with any interval in a slice of intervals.
//
// Parameters:
// - interval1: The interval to check for overlap.
// - intervals: The slice of intervals to compare against.
//
// Returns:
// - bool: True if there is an overlap, false otherwise.
func HasOverlap(interval1 *Interval, intervals []Interval) bool {
	for _, interval2 := range intervals {
		if interval1.IsOverlap(&interval2) {
			return true
		}
	}
	return false
}

// GetIntervalsBetween returns intervals that overlap with a given interval and sorts them by the start field.
//
// Parameters:
// - interval: The interval to check for overlaps.
// - intervals: The slice of intervals to check against.
//
// Returns:
// - []Interval: The intervals that overlap with the given interval, sorted by start.
func GetIntervalsBetween(interval *Interval, intervals []Interval) []Interval {
	var resultIntervals []Interval
	for _, interval2 := range intervals {
		if interval.IsOverlap(&interval2) {
			resultIntervals = append(resultIntervals, interval2)
		}
	}
	return SortByStart(resultIntervals)
}

// InsertSorted inserts a new interval into a sorted slice of intervals, maintaining the order.
//
// Parameters:
// - intervals: The sorted slice of intervals.
// - newInterval: The new interval to be inserted.
//
// Returns:
// - []Interval: The slice of intervals with the new interval inserted.
func InsertSorted(intervals []Interval, newInterval Interval) []Interval {
	index := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].a.Cmp(newInterval.a) >= 0
	})
	intervals = append(intervals, Interval{}) // add an empty interval to extend the slice
	copy(intervals[index+1:], intervals[index:])
	intervals[index] = newInterval
	return intervals
}
