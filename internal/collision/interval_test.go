package collision

import (
	"math/big"
	"testing"
)

// TestIsPointOverlap

func TestIsPointOverlap_FalseBeforeInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	result := interval.IsPointOverlap(big.NewInt(50)) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestIsPointOverlap_FalseAfterInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	result := interval.IsPointOverlap(big.NewInt(250)) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestIsPointOverlap_TrueInsideInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	result := interval.IsPointOverlap(big.NewInt(150)) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsPointOverlap_TrueAtStartOfInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	result := interval.IsPointOverlap(big.NewInt(100)) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsPointOverlap_TrueAtEndOfInterval(t *testing.T) {
	interval := new(Interval).SetInt(100, 200)
	result := interval.IsPointOverlap(big.NewInt(200)) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

// TestIsPointOverlap

// TestIsOverlap

func TestIsOverlap_NoOverlapBefore(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(0, 99)
	result := interval1.IsOverlap(interval2) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestIsOverlap_NoOverlapAdjacent(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(201, 300)
	result := interval1.IsOverlap(interval2) // expected false
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestIsOverlap_PartialOverlap_StartsInside(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(100, 150)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsOverlap_OverlapAtBoundary(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(200, 250)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsOverlap_PartialOverlap_EndsInside(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(50, 150)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsOverlap_PartialOverlap_Inside(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(150, 250)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsOverlap_CompleteOverlap(t *testing.T) {
	interval1 := new(Interval).SetInt(100, 200)
	interval2 := new(Interval).SetInt(0, 500)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestIsOverlap_ContainedWithin(t *testing.T) {
	interval1 := new(Interval).SetInt(0, 1000)
	interval2 := new(Interval).SetInt(150, 250)
	result := interval1.IsOverlap(interval2) // expected true
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

// TestIsOverlap
