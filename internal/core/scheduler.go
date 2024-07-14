package core

import (
	"GoKeyHunt/internal/console"
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/utils"
	"math/big"
	"time"
)

// Scheduler generates private keys within a specified range and sends them to an input channel.
//
// This function iterates from the start key to the end key, incrementing by 1, and sends each private key to the
// inputChannel. If VerboseProgress is enabled, it periodically prints the progress. The function also tracks the
// time elapsed and prints the final progress when done.
//
// Parameters:
// - start: A *big.Int representing the starting private key.
// - end: A *big.Int representing the ending private key.
// - params: A domain.Parameters instance containing configuration parameters, including UpdateInterval and VerboseProgress.
// - inputChannel: A send-only channel to which private keys are sent.
func Scheduler(start, end *big.Int, params domain.Parameters, inputChannel chan<- *big.Int) {
	privKey, increment := new(big.Int).Set(start), big.NewInt(1)

	ticker := time.NewTicker(time.Duration(params.UpdateInterval) * time.Second)
	startTime := time.Now()
	if params.VerboseProgress {
		defer ticker.Stop()
	} else {
		ticker.Stop()
	}

	for privKey.Cmp(end) <= 0 {
		select {
		case inputChannel <- utils.Clone(privKey):
			privKey.Add(privKey, increment)
		case <-ticker.C:
			console.PrintProgressString(start, end, privKey, startTime)
		}
	}
	console.PrintProgressString(start, end, privKey, startTime)
}
