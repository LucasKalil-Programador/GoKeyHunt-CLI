package main

import (
	"GoKeyHunt/internal/app_context"
	"GoKeyHunt/internal/collision"
	"GoKeyHunt/internal/console"
	"GoKeyHunt/internal/core"
	"GoKeyHunt/internal/output_results"
	"GoKeyHunt/internal/utils"
	"fmt"
	"math/big"
	"path/filepath"
	"sync"
	"time"
)

const version = "GoKeyHunt 1.0.0 | Created by Lucas Kalil"

// main is the entry point of the GoKeyHunt application. It initializes the application context,
// starts the main application logic, and prints a summary of the execution.
func main() {
	fmt.Println(version)
	ctx := createAppContext()
	startTime := time.Now()

	runApplication(ctx)

	sizeBeforeOp := ctx.Intervals.Size()
	ctx.Intervals.Save(ctx.CollisionPathFile)
	sizeAfterOp := ctx.Intervals.Size()

	console.PrintEndSummaryIfVerbose(ctx, startTime, sizeBeforeOp, sizeAfterOp)
}

// runApplication orchestrates the execution of the application logic.
// It sets up channels for worker communication, starts worker and output handler goroutines,
// and processes intervals in batches, handling collisions and scheduling tasks.
//
// Parameters:
// - ctx: The application context containing configuration parameters, wallet ranges, intervals, and results.
func runApplication(ctx *app_context.AppCtx) {
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
		start, end := utils.GetWalletStartAndEnd(ranges, params)
		startOriginal := utils.Clone(start)

		start = utils.GetStart(startOriginal, end, params, i+1)

		if start.Cmp(end) > 0 {
			break
		}

		console.PrintSummaryIfVerbose(startOriginal, start, end, params, i+1)

		hasCollision, newInterval := utils.HandleCollisions(startOriginal, start, end, params, intervals)
		if !hasCollision {
			start, end = newInterval.Get()
			core.Scheduler(start, end, params, inputChannel)
			intervals.Append(new(collision.Interval).Set(start, end))
		}
	}

	stopAndWaitWorkers(inputChannel, outputChannel, &workerGroup, &outputGroup)
}

// createAppContext initializes and returns a new application context.
// It loads data, sets parameters, and prepares file paths for collision and result data.
//
// Returns:
// - *app_context.AppCtx: The application context containing parameters, wallet ranges, intervals, and result paths.
func createAppContext() *app_context.AppCtx {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)

	rootDir := utils.GetRootDir()
	collisionPathFile := filepath.Join(rootDir, "data", fmt.Sprintf("wallet-%d-progress.json", params.TargetWallet))
	resultPathFile := filepath.Join(rootDir, "results.json")
	intervals := collision.ReadOrNew(collisionPathFile)
	results := output_results.ReadOrNew(resultPathFile)

	return &app_context.AppCtx{
		Params:            params,
		WalletRanges:      ranges,
		Wallets:           wallets,
		Intervals:         intervals,
		Results:           results,
		CollisionPathFile: collisionPathFile,
		ResultPathFile:    resultPathFile}
}

// stopAndWaitWorkers gracefully shuts down worker and output handler goroutines.
// It closes channels and waits for all goroutines to complete.
//
// Parameters:
// - inputChannel: The channel used to send input data to workers.
// - outputChannel: The channel used to receive output data from workers.
// - workerGroup: The WaitGroup used to synchronize worker goroutines.
// - outputGroup: The WaitGroup used to synchronize output handler goroutines.
func stopAndWaitWorkers(inputChannel, outputChannel chan *big.Int, workerGroup, outputGroup *sync.WaitGroup) {
	close(inputChannel)
	workerGroup.Wait()
	close(outputChannel)
	outputGroup.Wait()
}
