package core

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"math/big"
	"time"
)

func Scheduler(start, end *big.Int, params domain.Parameters, inputChannel chan<- *big.Int) {
	privKey, increment := new(big.Int).Set(start), big.NewInt(1)
	end = GetEndValue(start, end, params)

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
			PrintProgressString(start, end, privKey, startTime)
		}
	}
}

func GetEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = utils.MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}
