package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateRandomNumber generates a random big.Int between the specified minimum and maximum values (inclusive).
//
// This function takes two big.Int pointers representing the minimum and maximum bounds,
// and returns a random big.Int within that range. It returns an error if the minimum
// value is greater than the maximum value.
//
// Parameters:
// - minimum: A pointer to a big.Int representing the minimum value.
// - maximum: A pointer to a big.Int representing the maximum value.
//
// Returns:
// - *big.Int: A random big.Int within the specified range.
// - error: An error if the minimum value is greater than the maximum value or if there is an issue generating the random number.
func GenerateRandomNumber(minimum, maximum *big.Int) (*big.Int, error) {
	if minimum.Cmp(maximum) > 0 {
		return nil, fmt.Errorf("the minimum value cannot be greater than the maximum value")
	}

	interval := new(big.Int).Sub(maximum, minimum)
	randNum, err := rand.Int(rand.Reader, interval.Add(interval, big.NewInt(1)))
	if err != nil {
		return nil, err
	}

	return new(big.Int).Add(randNum, minimum), nil
}

// MaxBigInt returns the larger of two big.Int values.
//
// This function compares two big.Int pointers and returns the larger value.
//
// Parameters:
// - a: A pointer to a big.Int.
// - b: A pointer to a big.Int.
//
// Returns:
// - *big.Int: The larger of the two input values.
func MaxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// MinBigInt returns the smaller of two big.Int values.
//
// This function compares two big.Int pointers and returns the smaller value.
//
// Parameters:
// - a: A pointer to a big.Int.
// - b: A pointer to a big.Int.
//
// Returns:
// - *big.Int: The smaller of the two input values.
func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

// Clone creates a copy of the given big.Int value.
//
// This function takes a big.Int pointer and returns a new big.Int pointer
// that is a copy of the input value.
//
// Parameters:
// - input: A pointer to a big.Int to be cloned.
//
// Returns:
// - *big.Int: A new big.Int pointer that is a copy of the input value.
func Clone(input *big.Int) *big.Int {
	return new(big.Int).Set(input)
}
