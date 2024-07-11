package console

import (
	"btcgo/internal/app_context"
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"fmt"
	"math/big"
	"time"

	"github.com/dustin/go-humanize"
)

const summaryLabel = "------------------ Summary -------------------"
const tinySummaryLabel = "---------------- Tiny Summary ----------------"
const endSummaryLabel = "---------------- End Summary -----------------"

// sugar print

func PrintSummaryIfVerbose(startOriginal, start, end *big.Int, params domain.Parameters, batchCounter int) {
	if params.VerboseSummary {
		if batchCounter <= 1 {
			PrintSummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		} else {
			PrintTinySummary(startOriginal, utils.Clone(end), utils.Clone(start), params, batchCounter)
		}
	}
}

func PrintEndSummaryIfVerbose(ctx *app_context.AppCtx, startTime time.Time, sizeBeforeOp, sizeAfterOp int) {
	if ctx.Params.VerboseSummary {
		printEnd(ctx, startTime, sizeBeforeOp, sizeAfterOp)
	}
}

func printEnd(ctx *app_context.AppCtx, startTime time.Time, sizeBeforeOp, sizeAfterOp int) {
	intervalProgress := ctx.Intervals.CalculateTotalProgress()
	start, end := utils.GetStartAndEnd(*ctx.WalletRanges, *ctx.Params)
	totalProgress := new(big.Int).Sub(end, start)
	foundTarget := false
	for _, result := range ctx.Results.Resuts {
		if result.WalletIndex == ctx.Params.TargetWallet {

		}
	}
	PrintEndSummary(startTime, sizeBeforeOp, sizeAfterOp, intervalProgress, totalProgress, foundTarget)
}

// print

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

func PrintEndSummary(startTime time.Time, jsonSize, optimizedSize int, intervalProgress, totalProgress *big.Int, foundTarget bool) {
	progressPercent := new(big.Float).Quo(new(big.Float).SetInt(intervalProgress), new(big.Float).SetInt(totalProgress))
	progressPercent.Mul(progressPercent, big.NewFloat(100))

	intervalProgressStr := humanize.BigComma(intervalProgress)
	totalProgressStr := humanize.BigComma(totalProgress.Add(totalProgress, big.NewInt(1)))

	fmt.Printf("\n\n%s\n", endSummaryLabel)
	fmt.Printf("- Elapsed time: %v\n", time.Since(startTime).Truncate(time.Millisecond))
	fmt.Printf("- Wallet JSON size: %d\n", jsonSize)
	fmt.Printf("- After optimization: %d\n", optimizedSize)
	fmt.Printf("- progress: %v%%\n", progressPercent)
	fmt.Printf("- Overall progress: %s/%s\n", intervalProgressStr, totalProgressStr)
	fmt.Printf("- Wallet was found: %v\n", foundTarget)
	fmt.Printf("%s\n\n\n", endSummaryLabel)
}

// Aux functions

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

func getEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = utils.MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}
