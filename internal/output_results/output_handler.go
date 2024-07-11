package output_results

import (
	"btcgo/internal/domain"
	"fmt"
	"math/big"
	"sync"
)

func OutputHandler(params domain.Parameters, wallets domain.Wallets, resultArray *ResultArray, jsonPath string, outputChannel <-chan *big.Int, externalWg *sync.WaitGroup) {
	defer externalWg.Done()
	for key := range outputChannel {
		result := NewResult(key, wallets)
		added := resultArray.AppendIfNotExist(*result)
		handlerResult(added, result)
		if added {
			resultArray.Save(jsonPath)
		}
	}
}

func handlerResult(exist bool, result *Result) {
	var addedResultStr string
	if exist {
		addedResultStr = "Added to results.json"
	} else {
		addedResultStr = "Already exist on results.json"
	}
	fmt.Printf("\nFound key for the wallet: %d, %s\n", result.WalletIndex, addedResultStr)
}
