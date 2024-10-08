package console

import (
	"GoKeyHunt/internal/app_context"
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/utils"
	"fmt"
	"math/big"
	"time"

	"github.com/dustin/go-humanize"
)

const summaryLabel = "------------------ Summary -------------------"
const tinySummaryLabel = "---------------- Tiny Summary ----------------"
const endSummaryLabel = "---------------- End Summary -----------------"

// PrintSummaryIfVerbose prints a summary of the task if verbosity is enabled.
//
// This function prints a detailed summary or a compact summary depending on the batch counter value and verbosity settings.
//
// Parameters:
// - startOriginal: A *big.Int representing the original start value.
// - start: A *big.Int representing the current start value.
// - end: A *big.Int representing the end value.
// - params: A domain.Parameters instance containing configuration parameters.
// - batchCounter: The current batch count.
func PrintSummaryIfVerbose(startOriginal, start, end *big.Int, params domain.Parameters, batchCounter int) {
	if params.VerboseSummary {
		if batchCounter <= 1 {
			PrintSummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		} else {
			PrintTinySummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		}
	}
}

// PrintEndSummaryIfVerbose prints the end summary if verbosity is enabled.
//
// This function prints a summary of the task completion including elapsed time, JSON size, progress, and whether the target wallet was found.
//
// Parameters:
// - ctx: A pointer to app_context.AppCtx containing application context and configuration.
// - startTime: The time when the task started.
// - sizeBeforeOp: The size of the data before the operation.
// - sizeAfterOp: The size of the data after the operation.
func PrintEndSummaryIfVerbose(ctx *app_context.AppCtx, startTime time.Time, sizeBeforeOp, sizeAfterOp int) {
	if ctx.Params.VerboseSummary {
		printEnd(ctx, startTime, sizeBeforeOp, sizeAfterOp)
	}
}

// printEnd prints the end summary of the task.
//
// This function prints details about the progress and outcome of the task, including interval progress, total progress, and whether the target wallet was found.
//
// Parameters:
// - ctx: A pointer to app_context.AppCtx containing application context and configuration.
// - startTime: The time when the task started.
// - sizeBeforeOp: The size of the data before the operation.
// - sizeAfterOp: The size of the data after the operation.
func printEnd(ctx *app_context.AppCtx, startTime time.Time, sizeBeforeOp, sizeAfterOp int) {
	intervalProgress := ctx.Intervals.CalculateTotalProgress()
	start, end := utils.GetWalletStartAndEnd(*ctx.WalletRanges, *ctx.Params)
	totalProgress := new(big.Int).Sub(end, start)
	foundTarget := false
	for _, result := range ctx.Results.Resuts {
		if result.WalletIndex == ctx.Params.TargetWallet {
			foundTarget = true
			break
		}
	}
	PrintEndSummary(startTime, sizeBeforeOp, sizeAfterOp, intervalProgress, totalProgress, foundTarget)
}

// PrintSummary prints a detailed summary of the task.
//
// This function prints information such as the target wallet, range, worker count, batch size, update interval, and batch count.
//
// Parameters:
// - start: A *big.Int representing the start value.
// - end: A *big.Int representing the end value.
// - rng: A *big.Int representing the range value.
// - params: A domain.Parameters instance containing configuration parameters.
// - batchCounter: The current batch count.
func PrintSummary(start, end, rng *big.Int, params domain.Parameters, batchCounter int) {
	rngStr, startStr, endStr, workerCountStr, batchSizeStr,
		updateIntervalStr, batchCounterStr, maxBatchCounterStr := getStrings(rng, end, start, params, batchCounter)

	fmt.Printf("\n\n%s\n", summaryLabel)
	fmt.Printf("- Target wallet: %d\n", params.TargetWallet)
	if params.Rng {
		fmt.Printf("-  RNG: %s\n", rngStr)
	}
	fmt.Printf("- From: %s\n", startStr)
	fmt.Printf("-   To: %s\n", endStr)
	fmt.Printf("-\n")
	fmt.Printf("- Workers count: %s\n", workerCountStr)
	fmt.Printf("- Batch size: %v\n", batchSizeStr)
	fmt.Printf("- Use RNG start: %v\n", params.Rng)
	fmt.Printf("- Interval between updates: %s\n", updateIntervalStr)
	fmt.Printf("-\n")
	fmt.Printf("- Batch %s/%s\n", batchCounterStr, maxBatchCounterStr)
	fmt.Printf("%s\n\n\n", summaryLabel)
}

// PrintTinySummary prints a compact summary of the task.
//
// This function prints a smaller subset of information including RNG usage, batch size, and current batch count.
//
// Parameters:
// - start: A *big.Int representing the start value.
// - end: A *big.Int representing the end value.
// - rng: A *big.Int representing the range value.
// - params: A domain.Parameters instance containing configuration parameters.
// - batchCounter: The current batch count.
func PrintTinySummary(start, end, rng *big.Int, params domain.Parameters, batchCounter int) {
	rngStr, _, _, _, batchSizeStr, _, batchCounterStr, maxBatchCounterStr := getStrings(rng, end, start, params, batchCounter)

	fmt.Printf("\n\n%s\n", tinySummaryLabel)
	if params.Rng {
		fmt.Printf("-  RNG: %s\n", rngStr)
	}
	fmt.Printf("- Batch size: %v\n", batchSizeStr)
	fmt.Printf("- Batch %s/%s\n", batchCounterStr, maxBatchCounterStr)
	fmt.Printf("%s\n\n\n", tinySummaryLabel)
}

// PrintEndSummary prints the final summary of the task completion.
//
// This function prints details about the elapsed time, JSON size before and after optimization, progress, and whether the target wallet was found.
//
// Parameters:
// - startTime: The time when the task started.
// - jsonSize: The size of the JSON data before optimization.
// - optimizedSize: The size of the data after optimization.
// - intervalProgress: The progress made in the current interval.
// - totalProgress: The total progress of the task.
// - foundTarget: A boolean indicating whether the target wallet was found.
func PrintEndSummary(startTime time.Time, jsonSize, optimizedSize int, intervalProgress, totalProgress *big.Int, foundTarget bool) {
	progressPercent := new(big.Float).Quo(new(big.Float).SetInt(intervalProgress), new(big.Float).SetInt(totalProgress))
	progressPercent.Mul(progressPercent, big.NewFloat(100))

	intervalProgressStr := humanize.BigComma(intervalProgress)
	totalProgressStr := humanize.BigComma(totalProgress.Add(totalProgress, big.NewInt(1)))

	fmt.Printf("\n\n%s\n", endSummaryLabel)
	fmt.Printf("- Elapsed time: %v\n", time.Since(startTime).Truncate(time.Millisecond))
	fmt.Printf("- Wallet JSON size: %d\n", jsonSize)
	fmt.Printf("- After optimization: %d\n", optimizedSize)
	fmt.Printf("- progress: %v%%\n", progressPercent.Text('f', -1))
	fmt.Printf("- Overall progress: %s/%s\n", intervalProgressStr, totalProgressStr)
	fmt.Printf("- Wallet was found: %v\n", foundTarget)
	fmt.Printf("%s\n\n\n", endSummaryLabel)
}

// getStrings returns formatted strings for displaying various parameters and counts.
//
// This function generates formatted strings for the range, start, end, worker count, batch size, update interval, batch counter, and max batch counter.
//
// Parameters:
// - rng: A *big.Int representing the range value.
// - end: A *big.Int representing the end value.
// - start: A *big.Int representing the start value.
// - params: A domain.Parameters instance containing configuration parameters.
// - batchCounter: The current batch count.
//
// Returns:
// - string: Formatted string for the range value.
// - string: Formatted string for the start value.
// - string: Formatted string for the end value.
// - string: Formatted string for the worker count.
// - string: Formatted string for the batch size.
// - string: Formatted string for the update interval.
// - string: Formatted string for the current batch counter.
// - string: Formatted string for the maximum batch counter.
func getStrings(rng, end, start *big.Int, params domain.Parameters, batchCounter int) (string, string, string, string, string, string, string, string) {
	batchSize := utils.MinBigInt(new(big.Int).Sub(getEndValue(rng, end, params), rng), big.NewInt(params.BatchSize))
	if params.BatchSize == -1 {
		batchSize = new(big.Int).Sub(end, start)
	}
	batchSize = batchSize.Add(batchSize, big.NewInt(1))
	batchCount := big.NewInt(int64(params.BatchCount))
	if params.BatchCount == -1 && !params.Rng {
		bs := utils.MinBigInt(new(big.Int).Sub(getEndValue(rng, end, params), start), big.NewInt(params.BatchSize))
		x := new(big.Int).Sub(end, start)
		batchCount = new(big.Int).Add(new(big.Int).Div(x, bs), big.NewInt(1))
	}

	rngStr := humanize.BigComma(new(big.Int).Set(rng))
	startStr := humanize.BigComma(new(big.Int).Set(start))
	endStr := humanize.BigComma(new(big.Int).Set(end))
	workerCountStr := humanize.Comma(int64(params.WorkerCount))
	batchSizeStr := humanize.BigComma(batchSize)
	updateIntervalStr := humanize.Comma(int64(params.UpdateInterval))
	batchCounterStr := humanize.Comma(int64(batchCounter))
	maxBatchCounterStr := humanize.BigComma(batchCount)
	if params.BatchCount == -1 && params.Rng {
		maxBatchCounterStr = "infinity"
	}
	return rngStr, startStr, endStr, workerCountStr, batchSizeStr, updateIntervalStr, batchCounterStr, maxBatchCounterStr
}

// getEndValue calculates the end value for a batch based on the parameters.
//
// This function calculates the end value for the batch considering the batch size configuration.
//
// Parameters:
// - start: A *big.Int representing the start value.
// - end: A *big.Int representing the end value.
// - params: A domain.Parameters instance containing configuration parameters.
//
// Returns:
// - *big.Int: The calculated end value for the batch.
func getEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = utils.MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}
