package console

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"fmt"
	"math/big"

	"github.com/dustin/go-humanize"
)

func PrintSummary(start, end, rng *big.Int, params domain.Parameters, batchCounter int) {
	rngStr, startStr, endStr, workerCountStr, batchSizeStr,
		updateIntervalStr, batchCounterStr, maxBatchCounterStr := getStrings(rng, end, start, params, batchCounter)

	fmt.Printf("\n\n---------------- Summary ----------------\n")
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
	fmt.Printf("---------------- Summary ----------------\n\n\n")
}

func PrintTinySummary(start, end, rng *big.Int, params domain.Parameters, batchCounter int) {
	rngStr, _, _, _, batchSizeStr, _, batchCounterStr, maxBatchCounterStr := getStrings(rng, end, start, params, batchCounter)

	fmt.Printf("\n\n---------------- Tiny Summary ----------------\n")
	if params.Rng {
		fmt.Printf("-  RNG: %s\n", rngStr)
	}
	fmt.Printf("- Batch size: %v\n", batchSizeStr)
	fmt.Printf("- Batch %s/%s\n", batchCounterStr, maxBatchCounterStr)
	fmt.Printf("---------------- Tiny Summary ----------------\n\n\n")
}

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
