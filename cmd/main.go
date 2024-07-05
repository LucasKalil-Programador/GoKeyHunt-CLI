package main

import (
	"btcgo/internal/core"
	"btcgo/internal/utils"
	"math/big"
	"sync"
)

func main() {
	ranges, wallets := utils.LoadData()
	params := utils.GetParameters(*wallets)

	inputChannel := make(chan *big.Int, params.WorkerCount*2)
	outputChannel := make(chan *big.Int, 128)
	var wg sync.WaitGroup

	wg.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &wg)
	go core.Scheduler(ranges, params, inputChannel)
	go core.ProcessReceivedResults(outputChannel, wallets)
	wg.Wait()

	close(outputChannel)
}
