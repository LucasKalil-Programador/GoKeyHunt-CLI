package utils

import (
	"btcgo/internal/domain"
	"flag"
	"fmt"
	"log"
	"math"
	"runtime"
)

// GetParameters parses command-line flags and returns the parameters for the application.
//
// This function defines and parses various flags related to application configuration,
// such as worker thread count, target wallet, update interval, batch size, and verbosity options.
// It validates the parsed values and returns them as a Parameters structure.
//
// Parameters:
// - wallets: A domain.Wallets structure containing wallet addresses to validate against the target wallet flag.
//
// Returns:
// - domain.Parameters: A Parameters structure containing the parsed and validated flag values.
func GetParameters(wallets domain.Wallets) domain.Parameters {
	var maxInt64 int64 = math.MaxInt64

	// Variables to store flag values
	var workerCount, targetWallet, updateInterval, batchCount int
	var rng, verboseSummary, verboseProgress, verboseKeyFind bool
	var batchSize int64

	// Define flags
	flag.IntVar(&workerCount, "t", 2, fmt.Sprintf("Worker thread count (available CPUs: %d).", runtime.NumCPU()))
	flag.IntVar(&targetWallet, "w", 30, fmt.Sprintf("Target wallet (range: 0 to %d). Use 0 to search all wallets.", len(wallets.Addresses)))
	flag.IntVar(&updateInterval, "u", 1, "Progress update interval in seconds.")
	flag.Int64Var(&batchSize, "bs", -1, fmt.Sprintf("Batch size for execution (range: -1 to %d). If -1, will execute until the end of the wallet.", maxInt64))
	flag.IntVar(&batchCount, "bc", 1, fmt.Sprintf("Number of batches (range: 1 to %d). If -1, will execute until the end of the wallet.", math.MaxInt))
	flag.BoolVar(&rng, "rng", false, "If present, generate random start location.")
	flag.BoolVar(&verboseSummary, "vs", false, "Disable verbose output for summary.")
	flag.BoolVar(&verboseProgress, "vp", false, "Disable verbose output for progress.")
	flag.BoolVar(&verboseKeyFind, "vk", false, "Disable verbose output for key find.")

	// Parse flags
	flag.Parse()

	// Validate targetWallet
	if targetWallet < 0 || targetWallet > len(wallets.Addresses) {
		flag.Usage()
		log.Fatalf("\nError: Target wallet must be between 0 and %d.", len(wallets.Addresses))
	}

	// Validate batchSize
	if batchSize < -1 || batchSize == 0 {
		flag.Usage()
		log.Fatalf("\nError: Batch size must be -1 or greater than 0.")
	}

	// Validate batchCount
	if batchCount < -1 {
		flag.Usage()
		log.Fatalf("\nError: Batch count must be greater than 1.")
	}

	// Return parameters
	return domain.Parameters{
		WorkerCount:     workerCount,
		TargetWallet:    targetWallet,
		UpdateInterval:  updateInterval,
		BatchSize:       batchSize,
		BatchCount:      batchCount,
		Rng:             rng,
		VerboseSummary:  !verboseSummary,
		VerboseProgress: !verboseProgress,
		VerboseKeyFind:  !verboseKeyFind,
	}
}
