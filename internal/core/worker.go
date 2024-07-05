package core

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"math/big"
	"sync"
)

// Worker is a function that searches for a private key that matches a wallet address
func Worker(wallets *domain.Wallets, privKeyChan <-chan *big.Int, resultChan chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for privKeyInt := range privKeyChan {
		address := utils.CreatePublicHash160(privKeyInt)
		if utils.Contains(wallets.Addresses, address) {
			resultChan <- privKeyInt
		}
	}
}

func WorkersStartUp(params domain.Parameters, wallets *domain.Wallets, inputChannel chan *big.Int, outputChannel chan *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < params.WorkerCount; i++ {
		wg.Add(1)
		go Worker(wallets, inputChannel, outputChannel, wg)
	}
}
