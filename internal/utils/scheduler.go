package utils

import (
	"btcgo/internal/domain"
	"fmt"
	"math/big"

	"github.com/cheggaaa/pb/v3"
)

func Scheduler(ranges *domain.Ranges, params domain.Parameters, inputChannel chan *big.Int) {
	defer close(inputChannel)
	privKeyMin := new(big.Int)
	privKeyMin.SetString(ranges.Ranges[params.TargetWallet-1].Min[2:], 16)
	privKeyMax := new(big.Int)
	privKeyMax.SetString(ranges.Ranges[params.TargetWallet-1].Max[2:], 16)

	run_counter := new(int64)
	fmt.Printf("\nRun %d\n", *run_counter)
	progressBar := startPB(privKeyMax, privKeyMin)

	privKey := new(big.Int).Set(privKeyMin)
	for privKey.Cmp(privKeyMax) <= 0 {
		inputChannel <- new(big.Int).Set(privKey)
		privKey.Add(privKey, big.NewInt(1))

		progressBar = updatePB(progressBar, run_counter, privKeyMax, privKeyMin, privKey)
	}

	progressBar.Finish()
}

func updatePB(progressBar *pb.ProgressBar, run_counter *int64, privKeyMax *big.Int, privKeyMin *big.Int, currentPrivKey *big.Int) *pb.ProgressBar {
	if progressBar.Current() >= progressBar.Total() {
		progressBar.Finish()
		*run_counter++
		fmt.Printf("\nRun %d | Geral Progress: %0.2f%%\n", *run_counter, calcGeralProgress(privKeyMax, privKeyMin, currentPrivKey))
		progressBar = startPB(privKeyMax, privKeyMin)
	} else {
		progressBar.Increment()
	}
	return progressBar
}

func calcGeralProgress(privKeyMax *big.Int, privKeyMin *big.Int, currentPrivKey *big.Int) float64 {
	totalRange := new(big.Int).Sub(privKeyMax, privKeyMin)
	currentRange := new(big.Int).Sub(currentPrivKey, privKeyMin)
	proportion := new(big.Float).Quo(new(big.Float).SetInt(currentRange), new(big.Float).SetInt(totalRange))
	percentage, _ := proportion.Float64()
	percentageInt64 := percentage * 100
	return percentageInt64
}

func startPB(privKeyMax *big.Int, privKeyMin *big.Int) *pb.ProgressBar {
	var progressBar *pb.ProgressBar
	max_value_pb := big.NewInt(1_000_000_000)
	value_required := new(big.Int).Sub(privKeyMax, privKeyMin)

	if value_required.Cmp(max_value_pb) <= 0 {
		progressBar = pb.Start64(value_required.Int64())
	} else {
		progressBar = pb.Start64(max_value_pb.Int64())
	}
	return progressBar
}
