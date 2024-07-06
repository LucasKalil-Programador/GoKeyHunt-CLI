package main

import (
	"btcgo/internal/core"
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"math/big"
	"sync"
)

func main() {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)

	for i := 0; i < params.BatchCount; i++ {
		run(params, ranges, wallets, i+1)
	}
}

func run(params domain.Parameters, ranges *domain.Ranges, wallets *domain.Wallets, batchCounter int) {
	inputChannel := make(chan *big.Int, params.WorkerCount*2)
	outputChannel := make(chan *big.Int, 128)
	var wg sync.WaitGroup

	start, _ := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	end, _ := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)

	if params.Rng {
		start, _ = utils.GenerateRandomNumber(start, end)
	}

	core.PrintSummary(new(big.Int).Set(start), new(big.Int).Set(end), new(big.Int).Set(start), params, batchCounter)

	wg.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &wg)
	go core.Scheduler(start, end, params, inputChannel)
	go core.OutputHandler(outputChannel, wallets)
	wg.Wait()

	close(outputChannel)
	core.OutputHandler(outputChannel, wallets)
}
