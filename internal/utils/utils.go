package utils

import (
	"GoKeyHunt/internal/collision"
	"GoKeyHunt/internal/domain"
	"log"
	"math/big"
)

// GetWalletStartAndEnd retrieves the start and end values for a specific wallet from ranges.
// It converts the hex strings from the ranges to big integers.
//
// Parameters:
// - ranges: The domain.Ranges structure containing range information.
// - params: The domain.Parameters structure containing parameters including the target wallet.
//
// Returns:
// - *big.Int: The start value of the wallet range.
// - *big.Int: The end value of the wallet range.
func GetWalletStartAndEnd(ranges domain.Ranges, params domain.Parameters) (*big.Int, *big.Int) {
	start, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	if !ok {
		log.Fatal("Error converting start value")
	}
	end, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)
	if !ok {
		log.Fatal("Error converting end value")
	}
	return start, end
}

// GetStart calculates the start value for the current batch based on the given parameters.
// If RNG is enabled, it generates a random start value. If batch size is defined, it calculates the start based on the batch counter.
//
// Parameters:
// - start: The initial start value as a *big.Int.
// - end: The end value as a *big.Int.
// - params: The domain.Parameters structure containing parameters including batch size and RNG flag.
// - batchCounter: The current batch counter.
//
// Returns:
// - *big.Int: The calculated start value for the current batch.
func GetStart(start, end *big.Int, params domain.Parameters, batchCounter int) *big.Int {
	if params.Rng {
		start, _ = GenerateRandomNumber(start, end)
	} else if params.BatchSize != -1 {
		startAdd := new(big.Int).Mul(
			big.NewInt(params.BatchSize),
			new(big.Int).Sub(big.NewInt(int64(batchCounter)), big.NewInt(1)))
		start = new(big.Int).Add(start, startAdd)
	}
	return start
}

// GetEnd calculates the end value for the current batch based on the given parameters.
// It adjusts the end value if a batch size is defined.
//
// Parameters:
// - start: The start value as a *big.Int.
// - end: The end value as a *big.Int.
// - params: The domain.Parameters structure containing parameters including batch size.
//
// Returns:
// - *big.Int: The calculated end value for the current batch.
func GetEnd(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}

// HandleCollisions checks and resolves any collisions for the interval between start and end values.
// It adjusts the interval if necessary and returns whether a collision was resolved.
//
// Parameters:
// - startOriginal: The original start value as a *big.Int.
// - start: The start value of the current batch as a *big.Int.
// - end: The end value of the current batch as a *big.Int.
// - params: The domain.Parameters structure containing parameters including batch size.
// - intervals: The collision.IntervalArray structure containing existing intervals.
//
// Returns:
// - bool: True if a collision was resolved, false otherwise.
// - collision.Interval: The adjusted interval after handling collisions.
func HandleCollisions(startOriginal, start, end *big.Int, params domain.Parameters, intervals *collision.IntervalArray) (bool, collision.Interval) {
	end = GetEnd(start, end, params)
	interval := new(collision.Interval).Set(start, end)
	hasCollision, newInterval := intervals.HandleIntervalCollision(*interval)

	if hasCollision {
		dif := new(big.Int).Sub(end, start)
		newStart := MaxBigInt(startOriginal, new(big.Int).Sub(start, dif))
		newEnd := new(big.Int).Sub(end, dif)
		interval := new(collision.Interval).Set(newStart, newEnd)
		hasCollision, newInterval = intervals.HandleIntervalCollision(*interval)
	}
	return hasCollision, newInterval
}
