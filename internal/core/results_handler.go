package core

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"fmt"
	"math/big"
	"sync"
)

func OutputHandler(outputChannel <-chan *big.Int, wallets *domain.Wallets, params domain.Parameters, externalWg *sync.WaitGroup) {
	defer externalWg.Done()
	for result := range outputChannel {
		address := utils.CreatePublicHash160(result)
		walletIndex := utils.Find(wallets.Addresses, address) + 1
		wif := utils.GenerateWif(result)

		message := fmt.Sprintf("Wallet: %3d, Key: %064x, WIF %s\n", walletIndex, result, wif)
		utils.WriteInOutput(message)
		if params.VerboseKeyFind {
			fmt.Printf("\r%s", message)
		}
	}
}
