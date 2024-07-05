package utils

import (
	"btcgo/internal/domain"
	"flag"
	"fmt"
	"log"
	"runtime"
)

func GetParameters(wallets domain.Wallets) domain.Parameters {
	// Criando as variáveis para armazenar os valores das flags
	var workerCount int
	var targetWallet int
	var updateInterval int

	// Definindo as flags
	flag.IntVar(&workerCount, "t", 2, fmt.Sprintf("Worker thread count (available CPUs: %d).", runtime.NumCPU()))
	flag.IntVar(&targetWallet, "w", 30, fmt.Sprintf("Target wallet (range: 0 to %d). Use 0 to search all wallets.", len(wallets.Addresses)))
	flag.IntVar(&updateInterval, "u", 1, "Progress update interval in seconds.")

	// Parseando as flags
	flag.Parse()

	// Verificando se o targetWallet está no intervalo permitido
	if targetWallet < 0 || targetWallet > 161 {
		flag.Usage()
		log.Fatalf("\nError: Target wallet must be between 1 and 161.")
	}

	// Retornando os parâmetros
	return domain.Parameters{
		WorkerCount:    workerCount,
		TargetWallet:   targetWallet,
		UpdateInterval: updateInterval,
	}
}
