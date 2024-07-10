package collision

import (
	"testing"
)

func TestResolveCollisions_OverlapsWithExistingIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(50, 150),
		*new(Interval).SetInt(190, 290),
	}
	expected := new(Interval).SetInt(151, 189)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 151, 189 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_ExtendsToNextAvailableSpace(t *testing.T) {
	interval1 := new(Interval).SetInt(120, 300)
	intervals := []Interval{
		*new(Interval).SetInt(50, 150),
		*new(Interval).SetInt(151, 251),
	}
	expected := new(Interval).SetInt(252, 300)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 252, 300 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_NoOverlappingIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 300)
	intervals := []Interval{
		*new(Interval).SetInt(150, 200),
		*new(Interval).SetInt(201, 301),
	}
	expected := new(Interval).SetInt(100, 149)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 100, 149 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_NoAvailableSpaceAfterIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 300)
	intervals := []Interval{
		*new(Interval).SetInt(100, 200),
		*new(Interval).SetInt(201, 301),
		*new(Interval).SetInt(202, 303),
	}

	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 100, 149 with success = true
	if success {
		t.Errorf("expected failure but got %v", result)
	}
}

func TestResolveCollisions_EmptyIntervalArray(t *testing.T) {
	interval1 := new(Interval).SetInt(50, 150)
	intervals := []Interval{}
	expected := new(Interval).SetInt(50, 150)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 50, 150 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_IntervalInsideExistingInterval(t *testing.T) {
	interval1 := new(Interval).SetInt(60, 100)
	intervals := []Interval{
		*new(Interval).SetInt(50, 150),
	}
	expected := new(Interval).SetInt(60, 100)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // success = false
	if result.Equals(*expected) || success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_IntervalExactlyAtBoundary(t *testing.T) {
	interval1 := new(Interval).SetInt(150, 200)
	intervals := []Interval{
		*new(Interval).SetInt(100, 150),
		*new(Interval).SetInt(201, 250),
	}
	expected := new(Interval).SetInt(151, 200)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 151, 200 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_IntervalAtTheStartOfOverlap(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(100, 150),
		*new(Interval).SetInt(151, 200),
	}
	expected := new(Interval).SetInt(151, 200)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 151, 200 with success = true
	if result.Equals(*expected) || success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_LargeGapBetweenIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(50, 300)
	intervals := []Interval{
		*new(Interval).SetInt(50, 100),
		*new(Interval).SetInt(201, 300),
	}
	expected := new(Interval).SetInt(101, 200)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 101, 200 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestResolveCollisions_IntervalExactlyAtEndBoundary(t *testing.T) {
	interval1 := new(Interval).SetInt(250, 300)
	intervals := []Interval{
		*new(Interval).SetInt(200, 250),
	}
	expected := new(Interval).SetInt(251, 300)
	result, success := NewIntervalArray(intervals).ResolveCollisions(*interval1) // Expected: 250, 300 with success = true
	if !result.Equals(*expected) || !success {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
