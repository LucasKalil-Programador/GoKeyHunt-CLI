package core

import (
	"btcgo/internal/domain"
	"math/big"
	"time"
)

func Scheduler(ranges *domain.Ranges, params domain.Parameters, inputChannel chan<- *big.Int) {
	defer close(inputChannel)
	privKeyMin := new(big.Int)
	privKeyMin.SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	privKeyMax := new(big.Int)
	privKeyMax.SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)

	privKey := new(big.Int).Set(privKeyMin)
	increment := big.NewInt(1)

	ticker := time.NewTicker(time.Duration(params.UpdateInterval) * time.Second)
	go updater(privKeyMin, privKeyMax, privKey, ticker)
	defer ticker.Stop()

	for privKey.Cmp(privKeyMax) <= 0 {
		inputChannel <- new(big.Int).Set(privKey)
		privKey.Add(privKey, increment)
	}
}
