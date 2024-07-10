package collision

import (
	"reflect"
	"testing"
)

// TestHasOverlap

func TestHasOverlap_NoOverlap(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(0, 50),
		*new(Interval).SetInt(50, 99),
		*new(Interval).SetInt(201, 300),
	}
	result := HasOverlap(interval1, intervals) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestHasOverlap_OneOverlap(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(0, 50),
		*new(Interval).SetInt(150, 250),
		*new(Interval).SetInt(300, 400),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_MultipleOverlaps(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(50, 150),
		*new(Interval).SetInt(150, 250),
		*new(Interval).SetInt(300, 400),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_ExactMatch(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(100, 200),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_ContainedInterval(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(120, 180),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_EmptyIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{}
	result := HasOverlap(interval1, intervals) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestHasOverlap_SinglePointOverlapStart(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(200, 300),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_SinglePointOverlapEnd(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(50, 100),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_FullyWithinAnother(t *testing.T) {
	interval1 := new(Interval).SetInt(150, 170)
	intervals := []Interval{
		*new(Interval).SetInt(100, 200),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_EmptyInterval(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(0, 0)
	intervals := []Interval{*interval2}
	result := HasOverlap(interval1, intervals) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestHasOverlap_LargeIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(1000000, 2000000)
	intervals := []Interval{
		*new(Interval).SetInt(1500000, 2500000),
		*new(Interval).SetInt(2500000, 3000000),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_NegativeIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(-200, -100)
	intervals := []Interval{
		*new(Interval).SetInt(-300, -250),
		*new(Interval).SetInt(-150, -50),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_MixedPositiveNegativeIntervals(t *testing.T) {
	interval1 := new(Interval).SetInt(-100, 100)
	intervals := []Interval{
		*new(Interval).SetInt(-200, -150),
		*new(Interval).SetInt(-50, 50),
	}
	result := HasOverlap(interval1, intervals) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestHasOverlap_NoOverlapWithinRange(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(0, 50),
		*new(Interval).SetInt(250, 300),
	}
	result := HasOverlap(interval1, intervals) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

// TestHasOverlap

// TestGetIntervalsBetween

func TestGetIntervalsBetween_NoOverlap(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(0, 50),
		*new(Interval).SetInt(201, 300),
	}
	expected := []Interval{}
	result := GetIntervalsBetween(interval, intervals)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_OneOverlap(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(0, 50),
		*new(Interval).SetInt(150, 250),
		*new(Interval).SetInt(300, 400),
	}
	expected := []Interval{
		*new(Interval).SetInt(150, 250),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_MultipleOverlaps(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(50, 150),
		*new(Interval).SetInt(150, 250),
		*new(Interval).SetInt(300, 400),
	}
	expected := []Interval{
		*new(Interval).SetInt(50, 150),
		*new(Interval).SetInt(150, 250),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_ExactMatch(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(100, 200),
	}
	expected := []Interval{
		*new(Interval).SetInt(100, 200),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_ContainedInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{
		*new(Interval).SetInt(120, 180),
	}
	expected := []Interval{
		*new(Interval).SetInt(120, 180),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_EmptyIntervals(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	intervals := []Interval{}
	expected := []Interval{}
	result := GetIntervalsBetween(interval, intervals)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_MixedPositiveNegativeIntervals(t *testing.T) {
	interval := new(Interval).SetInt(-100, 100)
	intervals := []Interval{
		*new(Interval).SetInt(-200, -150),
		*new(Interval).SetInt(-50, 50),
	}
	expected := []Interval{
		*new(Interval).SetInt(-50, 50),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetIntervalsBetween_SortedResults(t *testing.T) {
	interval := new(Interval).SetInt(100, 300)
	intervals := []Interval{
		*new(Interval).SetInt(250, 350),
		*new(Interval).SetInt(150, 200),
		*new(Interval).SetInt(200, 250),
	}
	expected := []Interval{
		*new(Interval).SetInt(150, 200),
		*new(Interval).SetInt(200, 250),
		*new(Interval).SetInt(250, 350),
	}
	result := GetIntervalsBetween(interval, intervals)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestGetIntervalsBetween
