package core

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"math/big"
	"time"
)

func Scheduler(start, end *big.Int, params domain.Parameters, inputChannel chan<- *big.Int) {
	defer close(inputChannel)

	privKey, increment := new(big.Int).Set(start), big.NewInt(1)
	counter := params.BatchSize

	end = GetEndValue(start, end, params)

	ticker := time.NewTicker(time.Duration(params.UpdateInterval) * time.Second)
	go updater(start, end, privKey, ticker)
	defer ticker.Stop()

	for privKey.Cmp(end) <= 0 {
		inputChannel <- new(big.Int).Set(privKey)
		privKey.Add(privKey, increment)

		if counter != -1 {
			if counter <= 0 {
				break
			} else {
				counter--
			}
		}
	}
}

func GetEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = utils.MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize)), end)
	}
	return end
}
