package main

import (
	"btcgo/internal/collision"
	"btcgo/internal/console"
	"btcgo/internal/core"
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"math/big"
	"sync"
)

func main() {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)
	intervalArray := collision.NewEmptyIntervalArray()
	run(params, ranges, wallets, intervalArray)
	intervalArray.Optimize()
}

func run(params domain.Parameters, ranges *domain.Ranges, wallets *domain.Wallets, intervals *collision.IntervalArray) {
	inputChannel := make(chan *big.Int, params.WorkerCount*2)
	outputChannel := make(chan *big.Int, params.WorkerCount)
	var workerGroup, outputGroup sync.WaitGroup

	workerGroup.Add(1)
	outputGroup.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &workerGroup)
	go core.OutputHandler(outputChannel, wallets, params, &outputGroup)

	for i := 0; i < params.BatchCount || params.BatchCount == -1; i++ {
		batchCounter := i + 1
		start, end := utils.GetStartAndEnd(ranges, params)
		startClone := utils.Clone(start)

		if params.Rng {
			start, _ = utils.GenerateRandomNumber(start, end)
		} else if params.BatchSize != -1 {
			startAdd := new(big.Int).Mul(
				big.NewInt(params.BatchSize),
				new(big.Int).Sub(big.NewInt(int64(batchCounter)), big.NewInt(1)))
			start = new(big.Int).Add(start, startAdd)
		}

		if start.Cmp(end) > 0 {
			break
		}

		if params.VerboseSummary {
			if batchCounter <= 1 {
				console.PrintSummary(startClone, utils.Clone(end), utils.Clone(start), params, batchCounter)
			} else {
				console.PrintTinySummary(startClone, utils.Clone(end), utils.Clone(start), params, batchCounter)
			}
		}

		end = utils.GetEndValue(start, end, params)
		interval := new(collision.Interval).Set(start, end)
		hasCollision, newInterval := intervals.HandleIntervalCollision(*interval)

		if !hasCollision {
			start, end = newInterval.Get()
			core.Scheduler(start, end, params, inputChannel)
			intervals.Append(new(collision.Interval).Set(start, end))
		}
	}

	close(inputChannel)
	workerGroup.Wait()
	close(outputChannel)
	outputGroup.Wait()
}
