package core

import (
	"btcgo/internal/domain"
	"fmt"
	"math/big"

	"github.com/dustin/go-humanize"
)

func PrintSummary(start, end, rng *big.Int, params domain.Parameters, batchCounter int) {

	fmt.Printf("\n\n---------------- Summary ----------------\n")
	fmt.Printf("- Target wallet: %d\n", params.TargetWallet)
	fmt.Printf("-  RNG: %s\n", humanize.BigComma(new(big.Int).Set(rng)))
	fmt.Printf("- From: %s\n", humanize.BigComma(new(big.Int).Set(start)))
	fmt.Printf("-   To: %s\n", humanize.BigComma(new(big.Int).Set(end)))
	fmt.Printf("-\n")
	fmt.Printf("- Workers count: %s\n", humanize.Comma(int64(params.WorkerCount)))
	fmt.Printf("- Batch size: %v\n", humanize.BigComma(new(big.Int).Sub(GetEndValue(rng, end, params), start)))
	fmt.Printf("- Use RNG start: %v\n", params.Rng)
	fmt.Printf("- Interval between updates: %s\n", humanize.Comma(int64(params.UpdateInterval)))
	fmt.Printf("-\n")
	fmt.Printf("- Batch %s/%s\n", humanize.Comma(int64(batchCounter)), humanize.Comma(int64(params.BatchCount)))
	fmt.Printf("---------------- Summary ----------------\n\n\n")

}
