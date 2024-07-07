package utils

import (
	"btcgo/internal/domain"
	"flag"
	"fmt"
	"log"
	"math"
	"runtime"
)

func GetParameters(wallets domain.Wallets) domain.Parameters {
	var maxInt64 int64 = math.MaxInt64

	// Criando as variáveis para armazenar os valores das flags
	var workerCount int
	var targetWallet int
	var updateInterval int
	var batchSize int64
	var batchCount int
	var rng bool

	// Definindo as flags
	flag.IntVar(&workerCount, "t", 2, fmt.Sprintf("Worker thread count (available CPUs: %d).", runtime.NumCPU()))
	flag.IntVar(&targetWallet, "w", 30, fmt.Sprintf("Target wallet (range: 0 to %d). Use 0 to search all wallets.", len(wallets.Addresses)))
	flag.IntVar(&updateInterval, "u", 1, "Progress update interval in seconds.")
	flag.Int64Var(&batchSize, "bs", -1, fmt.Sprintf("Batch size for execution (range: -1 to %d). If -1, will execute until the end of the wallet.", maxInt64))
	flag.IntVar(&batchCount, "bc", 1, fmt.Sprintf("Number of batches (range: 1 to %d). If -1, will execute until the end of the wallet.", math.MaxInt))
	flag.BoolVar(&rng, "rng", false, "If present, generate random start location.")

	// Parseando as flags
	flag.Parse()

	// Verificando se o targetWallet está no intervalo permitido
	if targetWallet < 0 || targetWallet > len(wallets.Addresses) {
		flag.Usage()
		log.Fatalf("\nError: Target wallet must be between 0 and %d.", len(wallets.Addresses))
	}

	// Verificando se o batchSize é válido (não negativo, exceto para -1)
	if batchSize < -1 || batchSize == 0 {
		flag.Usage()
		log.Fatalf("\nError: Batch size must be -1 or greater than 0.")
	}

	// Verificando se o batchCount é válido (não negativo, exceto para -1)
	if batchCount < -1 {
		flag.Usage()
		log.Fatalf("\nError: Batch count must greater than 1.")
	}

	// Retornando os parâmetros
	return domain.Parameters{
		WorkerCount:    workerCount,
		TargetWallet:   targetWallet,
		UpdateInterval: updateInterval,
		BatchSize:      batchSize,
		BatchCount:     batchCount,
		Rng:            rng,
	}
}
