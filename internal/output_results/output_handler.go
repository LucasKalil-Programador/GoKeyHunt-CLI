package output_results

import (
	"GoKeyHunt/internal/domain"
	"fmt"
	"math/big"
	"sync"
)

// OutputHandler processes keys received from an output channel and updates the ResultArray.
//
// This function listens on the provided outputChannel for big.Int keys. For each key, it creates a new Result
// and attempts to append it to the resultArray if it does not already exist. If the Result is new and the VerboseKeyFind
// parameter is set, it prints the result. If a new Result is added, it saves the resultArray to a JSON file.
//
// Parameters:
// - params: A domain.Parameters instance containing configuration parameters.
// - wallets: A domain.Wallets instance containing wallet addresses.
// - resultArray: A pointer to the ResultArray instance to be updated.
// - jsonPath: A string representing the path to the JSON file where results will be saved.
// - outputChannel: A receive-only channel from which big.Int keys are received.
// - externalWg: A pointer to a sync.WaitGroup that is decremented when the function completes.
func OutputHandler(params domain.Parameters, wallets domain.Wallets, resultArray *ResultArray, jsonPath string, outputChannel <-chan *big.Int, externalWg *sync.WaitGroup) {
	defer externalWg.Done()
	for key := range outputChannel {
		result := NewResult(key, wallets)
		added := resultArray.AppendIfNotExist(*result)

		if params.VerboseKeyFind {
			printResult(added, result)
		}

		if added {
			resultArray.Save(jsonPath)
		}
	}
}

// printResult prints the result of processing a key.
//
// This function prints whether the Result was added to the results.json file or if it already existed.
//
// Parameters:
// - added: A boolean indicating whether the Result was added to the results.json file.
// - result: A pointer to the Result instance to be printed.
func printResult(added bool, result *Result) {
	var addedResultStr string
	if added {
		addedResultStr = "Added to results.json"
	} else {
		addedResultStr = "Already exist on results.json"
	}
	fmt.Printf("\nFound key for the wallet: %d, %s\n", result.WalletIndex, addedResultStr)
}
