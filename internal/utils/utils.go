package utils

import (
	"GoKeyHunt/internal/collision"
	"GoKeyHunt/internal/domain"
	"log"
	"math/big"
)

func GetStartAndEnd(ranges domain.Ranges, params domain.Parameters) (*big.Int, *big.Int) {
	start, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de início")
	}
	end, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de fim")
	}
	return start, end
}

func GetEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}

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

func HandleCollisions(startOriginal, start, end *big.Int, params domain.Parameters, intervals *collision.IntervalArray) (bool, collision.Interval) {
	end = GetEndValue(start, end, params)
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
