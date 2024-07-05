package core

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"fmt"
	"math/big"
)

func ProcessReceivedResults(outputChannel <-chan *big.Int, wallets *domain.Wallets) {
	for result := range outputChannel {
		address := utils.CreatePublicHash160(result)
		walletIndex := utils.Find(wallets.Addresses, address) + 1
		wif := utils.GenerateWif(result)

		message := fmt.Sprintf("Wallet: %3d, Key: %064x, WIF %s\n", walletIndex, result, wif)
		utils.WriteInOutput(message)
		fmt.Printf("\r%s", message)
	}
}
