package main

import (
	"btcgo/internal/core"
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"log"
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
	outputChannel := make(chan *big.Int, params.WorkerCount)
	var workerGroup, outputGroup sync.WaitGroup

	start, end := getStartAndEnd(ranges, params)
	startClone := utils.Clone(start)

	if params.Rng {
		start, _ = utils.GenerateRandomNumber(start, end)
	}

	core.PrintSummary(startClone, utils.Clone(end), utils.Clone(start), params, batchCounter)

	workerGroup.Add(1)
	outputGroup.Add(1)
	go core.WorkersStartUp(params, wallets, inputChannel, outputChannel, &workerGroup)
	go core.OutputHandler(outputChannel, wallets, &outputGroup)
	core.Scheduler(start, end, params, inputChannel)

	close(inputChannel)
	workerGroup.Wait()
	close(outputChannel)
	outputGroup.Wait()
}

func getStartAndEnd(ranges *domain.Ranges, params domain.Parameters) (*big.Int, *big.Int) {
	start, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de início")
	}
	end, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de fim")
	}
	return start, end
}
