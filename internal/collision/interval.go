package collision

import (
	"fmt"
	"math/big"
)

type Interval struct {
	a, b *big.Int
}

func (interval *Interval) Set(a, b *big.Int) *Interval {
	interval.a, interval.b = new(big.Int).Set(a), new(big.Int).Set(b)
	return interval
}

func (interval *Interval) SetInt(a, b int) *Interval {
	return interval.SetInt64(int64(a), int64(b))
}

func (interval *Interval) SetInt64(a, b int64) *Interval {
	interval.a = big.NewInt(a)
	interval.b = big.NewInt(b)
	return interval
}

func (interval *Interval) SetString(a, b string, base int) (*Interval, bool) {
	newA, successA := new(big.Int).SetString(a, base)
	newB, successB := new(big.Int).SetString(a, base)
	interval.a, interval.b = newA, newB
	return interval, successA && successB
}

func (i Interval) Equals(other Interval) bool {
	return i.a.Cmp(other.a) == 0 && i.b.Cmp(other.b) == 0
}

func (i Interval) String() string {
	return fmt.Sprintf("Interval[Start: %s, End: %s]", i.a.String(), i.b.String())
}

func (i *Interval) Clone() *Interval {
	clone := &Interval{
		a: new(big.Int).Set(i.a),
		b: new(big.Int).Set(i.b),
	}
	return clone
}

func (i *Interval) IsPointOverlap(point *big.Int) bool {
	return i.a.Cmp(point) <= 0 && i.b.Cmp(point) >= 0
}

func (i1 *Interval) IsOverlap(i2 *Interval) bool {
	return i1.IsPointOverlap(i2.a) || i1.IsPointOverlap(i2.b) || i2.IsPointOverlap(i1.a) || i2.IsPointOverlap(i1.b)
}

func (i1 *Interval) Get() (*big.Int, *big.Int) {
	return i1.a, i1.b
}
