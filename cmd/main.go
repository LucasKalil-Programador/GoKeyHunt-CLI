package main

import (
	"btcgo/internal/core"
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"fmt"
	"math/big"
	"sync"
)

func main() {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)

	run(params, ranges, wallets)
}

func run(params domain.Parameters, ranges *domain.Ranges, wallets *domain.Wallets) {
	inputChannel := make(chan *big.Int, params.WorkerCount*2)
	outputChannel := make(chan *big.Int, 128)
	var wg sync.WaitGroup

	start, _ := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	end, _ := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)

	if params.Rng {
		start, _ = utils.GenerateRandomNumber(start, end)
	}

	printSummary(start, end, params)

	wg.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &wg)
	go core.Scheduler(start, end, params, inputChannel)
	go core.OutputHandler(outputChannel, wallets)
	wg.Wait()

	close(outputChannel)
	core.OutputHandler(outputChannel, wallets)
}

func printSummary(start, end *big.Int, params domain.Parameters) {
	fmt.Printf("\n\n------------- Summary -------------\n")
	fmt.Printf("- Target wallet: %v\n", params.TargetWallet)
	fmt.Printf("- From: %v\n", start)
	fmt.Printf("-   To: %v\n", end)
	fmt.Printf("-\n")
	fmt.Printf("- Workers count: %v\n", params.WorkerCount)
	batchStr := new(big.Int).Sub(core.GetEndValue(start, end, params), start).String()
	fmt.Printf("- Batch size: %v\n", batchStr)
	fmt.Printf("- Use RNG start: %v\n", params.Rng)
	fmt.Printf("- Interval between updates: %v\n", params.UpdateInterval)
	fmt.Printf("------------- Summary -------------\n\n\n")
}
