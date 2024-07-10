package utils

import (
	"btcgo/internal/domain"
	"log"
	"math/big"
)

func GetStartAndEnd(ranges *domain.Ranges, params domain.Parameters) (*big.Int, *big.Int) {
	start, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Min[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de início")
	}
	end, ok := new(big.Int).SetString(ranges.Ranges[params.TargetWallet].Max[2:], 16)
	if !ok {
		log.Fatal("Erro ao converter o valor de fim")
	}
	return start, end
}

func GetEndValue(start, end *big.Int, params domain.Parameters) *big.Int {
	if params.BatchSize != -1 {
		end = MinBigInt(new(big.Int).Add(start, big.NewInt(params.BatchSize-1)), end)
	}
	return end
}
