package core

import (
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/utils"
	"math/big"
	"sync"
)

// Worker is a function that searches for a private key that matches a wallet address.
//
// This function listens on the privKeyChan for big.Int private keys. For each key, it generates the corresponding
// wallet address and checks if it exists in the provided wallets. If a match is found, the private key is sent
// to the resultChan.
//
// Parameters:
// - wallets: A domain.Wallets instance containing wallet addresses.
// - privKeyChan: A receive-only channel from which big.Int private keys are received.
// - resultChan: A send-only channel to which matching big.Int private keys are sent.
// - wg: A pointer to a sync.WaitGroup that is decremented when the function completes.
func Worker(wallets domain.Wallets, privKeyChan <-chan *big.Int, resultChan chan<- *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	for privKeyInt := range privKeyChan {
		address := utils.CreatePublicHash160(privKeyInt)
		if utils.Contains(wallets.Addresses, address) {
			resultChan <- privKeyInt
		}
	}
}

// WorkersStartUp initializes and starts multiple Worker goroutines.
//
// This function spawns a number of Worker goroutines specified by params.WorkerCount. Each Worker listens on the
// inputChannel for big.Int private keys and sends matching keys to the outputChannel. The function ensures that
// the sync.WaitGroup is properly managed by incrementing the counter before starting the Workers and decrementing
// it after they are done.
//
// Parameters:
// - params: A domain.Parameters instance containing configuration parameters, including WorkerCount.
// - wallets: A domain.Wallets instance containing wallet addresses.
// - inputChannel: A channel from which big.Int private keys are received by Workers.
// - outputChannel: A channel to which matching big.Int private keys are sent by Workers.
// - wg: A pointer to a sync.WaitGroup that tracks the completion of Worker goroutines.
func WorkersStartUp(params domain.Parameters, wallets domain.Wallets, inputChannel chan *big.Int, outputChannel chan *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(params.WorkerCount)

	for i := 0; i < params.WorkerCount; i++ {
		go Worker(wallets, inputChannel, outputChannel, wg)
	}
}
