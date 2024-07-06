package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomNumber(minimum, maximum *big.Int) (*big.Int, error) {
	if minimum.Cmp(maximum) > 0 {
		return nil, fmt.Errorf("the minimum value cannot be greater than the maximum value")
	}

	// Generate the difference between maximum and minimum.
	interval := new(big.Int).Sub(maximum, minimum)

	// Generate a random number within the interval.
	randNum, err := rand.Int(rand.Reader, interval.Add(interval, big.NewInt(1)))
	if err != nil {
		return nil, err
	}

	// Add the minimum value to the random number.
	randNum.Add(randNum, minimum)

	return randNum, nil
}

func MaxBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

func MinBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

func Clone(input *big.Int) *big.Int {
	return new(big.Int).Set(input)
}
