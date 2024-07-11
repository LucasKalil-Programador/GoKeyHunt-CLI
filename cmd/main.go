package main

import (
	"btcgo/internal/app_context"
	"btcgo/internal/collision"
	"btcgo/internal/console"
	"btcgo/internal/core"
	"btcgo/internal/domain"
	"btcgo/internal/output_results"
	"btcgo/internal/utils"
	"fmt"
	"math/big"
	"path/filepath"
	"sync"
)

func main() {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)

	collisionPathFile := filepath.Join(utils.GetRootDir(), "data", fmt.Sprintf("wallet-%d.json", params.TargetWallet))
	resultPathFile := filepath.Join(utils.GetRootDir(), "results.json")
	intervals := collision.ReadOrNew(collisionPathFile)
	results := output_results.ReadOrNew(resultPathFile)

	ctx := app_context.AppCtx{
		Params:            params,
		WalletRanges:      ranges,
		Wallets:           wallets,
		Intervals:         intervals,
		Results:           results,
		CollisionPathFile: collisionPathFile,
		ResultPathFile:    resultPathFile}

	run(&ctx)

	ctx.Intervals.Save(collisionPathFile)
}

func run(ctx *app_context.AppCtx) {
	// unpack parameters from ctx
	params, ranges, wallets, resultsJsonPath := *ctx.Params, *ctx.WalletRanges, *ctx.Wallets, ctx.ResultPathFile
	intervals, results := ctx.Intervals, ctx.Results

	inputChannel := make(chan *big.Int, params.WorkerCount*2)
	outputChannel := make(chan *big.Int, params.WorkerCount)
	var workerGroup, outputGroup sync.WaitGroup

	workerGroup.Add(1)
	outputGroup.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &workerGroup)
	go output_results.OutputHandler(params, wallets, results, resultsJsonPath, outputChannel, &outputGroup)

	for i := 0; i < params.BatchCount || params.BatchCount == -1; i++ {
		start, end := utils.GetStartAndEnd(ranges, params)
		startOriginal := utils.Clone(start)

		start = getStart(startOriginal, end, params, i+1)

		if start.Cmp(end) > 0 {
			break
		}

		printSummaryIfVerbose(startOriginal, start, end, params, i+1)

		hasCollision, newInterval := handleCollisions(startOriginal, start, end, params, intervals)
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

func getStart(start, end *big.Int, params domain.Parameters, batchCounter int) *big.Int {
	if params.Rng {
		start, _ = utils.GenerateRandomNumber(start, end)
	} else if params.BatchSize != -1 {
		startAdd := new(big.Int).Mul(
			big.NewInt(params.BatchSize),
			new(big.Int).Sub(big.NewInt(int64(batchCounter)), big.NewInt(1)))
		start = new(big.Int).Add(start, startAdd)
	}
	return start
}

func handleCollisions(startOriginal, start, end *big.Int, params domain.Parameters, intervals *collision.IntervalArray) (bool, collision.Interval) {
	end = utils.GetEndValue(start, end, params)
	interval := new(collision.Interval).Set(start, end)
	hasCollision, newInterval := intervals.HandleIntervalCollision(*interval)

	if hasCollision {
		dif := new(big.Int).Sub(end, start)
		newStart := utils.MaxBigInt(startOriginal, new(big.Int).Sub(start, dif))
		newEnd := new(big.Int).Sub(end, dif)
		interval := new(collision.Interval).Set(newStart, newEnd)
		hasCollision, newInterval = intervals.HandleIntervalCollision(*interval)
	}
	return hasCollision, newInterval
}

func printSummaryIfVerbose(startOriginal, start, end *big.Int, params domain.Parameters, batchCounter int) {
	if params.VerboseSummary {
		if batchCounter <= 1 {
			console.PrintSummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		} else {
			console.PrintTinySummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		}
	}
}
