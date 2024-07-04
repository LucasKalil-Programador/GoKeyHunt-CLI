package main

import (
	"btcgo/internal/core"
	"btcgo/internal/utils"
	"fmt"
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
	go utils.Scheduler(ranges, params, inputChannel)
	wg.Wait()

	close(outputChannel)

	for output := range outputChannel {
		fmt.Printf("\nResultado: %s, PrivKey: %064x", utils.GenerateWif(output), output)
	}
}
