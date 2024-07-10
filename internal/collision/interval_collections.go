package collision

import (
	"sort"
)

// Define a type that implements sort.Interface
type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].a.Cmp(a[j].a) < 0 }

// Function to sort an array of Interval by Start field
func SortByStart(intervals []Interval) []Interval {
	sort.Sort(ByStart(intervals))
	return intervals
}

func HasOverlap(interval1 *Interval, intervals []Interval) bool {
	for _, interval2 := range intervals {
		if interval1.IsOverlap(&interval2) {
			// fmt.Printf("interval1: %v, interval2: %v\n", interval1, interval2)
			return true
		}
	}
	return false
}

func GetIntervalsBetween(interval *Interval, intervals []Interval) []Interval {
	var resultIntervals []Interval
	for _, interval2 := range intervals {
		if interval.IsOverlap(&interval2) {
			resultIntervals = append(resultIntervals, interval2)
		}
	}
	return SortByStart(resultIntervals)
}

func InsertSorted(intervals []Interval, newInterval Interval) []Interval {
	index := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].a.Cmp(newInterval.a) >= 0
	})
	intervals = append(intervals, Interval{}) // add an empty interval to extend the slice
	copy(intervals[index+1:], intervals[index:])
	intervals[index] = newInterval
	return intervals
}
