package core

import (
	"GoKeyHunt/internal/console"
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/utils"
	"math/big"
	"time"
)

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
