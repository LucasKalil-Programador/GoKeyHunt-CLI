package collision

import (
	"fmt"
	"math/big"
)

type IntervalArray struct {
	intervalsArr []Interval
}

// Método String() para IntervalArray
func (ia *IntervalArray) String() string {
	intervalStrings := make([]string, len(ia.intervalsArr))
	for i, interval := range ia.intervalsArr {
		intervalStrings[i] = interval.String()
	}

	return fmt.Sprintf("[%s]", intervalStrings)
}

func NewIntervalArray(intervals []Interval) *IntervalArray {
	newIntervalsArr := make([]Interval, len(intervals))
	copy(newIntervalsArr, intervals)
	SortByStart(intervals)
	return &IntervalArray{intervalsArr: newIntervalsArr}
}

func NewEmptyIntervalArray() *IntervalArray {
	var newIntervalsArr []Interval
	return &IntervalArray{intervalsArr: newIntervalsArr}
}

func (interArray *IntervalArray) Append(interval *Interval) {
	interArray.intervalsArr = InsertSorted(interArray.intervalsArr, *interval)
}

func (interArray *IntervalArray) ResolveCollisions(target Interval) (*Interval, bool) {
	newInterval := target.Clone()
	for _, interval := range interArray.intervalsArr {
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

func (interArray *IntervalArray) HandleIntervalCollision(interval Interval) (bool, Interval) {
	hasCollision := HasOverlap(&interval, interArray.intervalsArr)
	if hasCollision {
		// fmt.Println("Colisao")
		intervals := GetIntervalsBetween(&interval, interArray.intervalsArr)
		newInterval, valid := NewIntervalArray(intervals).ResolveCollisions(interval)
		if valid {
			newHasCollision := HasOverlap(newInterval, interArray.intervalsArr)
			if !newHasCollision {
				// fmt.Printf("intervals: %v\n", intervals)
				// fmt.Printf("interval: %v, newInterval: %v\n", interval, newInterval)
				// fmt.Println("Resolvido")
				// os.Exit(0)
				return newHasCollision, *newInterval
			}
			// else {
			// 	fmt.Printf("intervals: %v\n", intervals)
			// 	fmt.Printf("interval: %v, newInterval: %v\n", interval, newInterval)
			// 	fmt.Println("Nao resolvido")
			// 	os.Exit(0)
			// }
		}
	}
	return hasCollision, interval
}

func (interArray *IntervalArray) Optimize() int {
	length := len(interArray.intervalsArr)
	intervals := interArray.intervalsArr
	rmCount := 0

	i := 0
	for i < length-1 {
		interval1, interval2 := intervals[i], intervals[i+1]
		if new(big.Int).Sub(interval2.a, interval1.b).Cmp(big.NewInt(1)) <= 0 {
			intervals[i] = *new(Interval).Set(interval1.a, interval2.b)
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			rmCount++
			length--
		} else {
			i++
		}
	}

	interArray.intervalsArr = intervals
	return rmCount
}
