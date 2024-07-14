package app_context

import (
	"GoKeyHunt/internal/collision"
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/output_results"
)

// AppCtx represents the application context containing various components and configuration parameters.
//
// This struct holds references to configuration parameters, wallet ranges, wallets, collision intervals, and result arrays.
// It also includes file paths for saving collision and result data.
//
// Fields:
// - Params: A pointer to domain.Parameters, which contains configuration parameters for the application.
// - WalletRanges: A pointer to domain.Ranges, which defines the ranges of wallet addresses to be processed.
// - Wallets: A pointer to domain.Wallets, which contains the wallet addresses to be searched.
// - Intervals: A pointer to collision.IntervalArray, which manages intervals of collision results.
// - Results: A pointer to output_results.ResultArray, which stores the results of key searches.
//
// - CollisionPathFile: A string representing the file path where collision data is saved.
// - ResultPathFile: A string representing the file path where result data is saved.
type AppCtx struct {
	Params       *domain.Parameters          // Application configuration parameters.
	WalletRanges *domain.Ranges              // Ranges of wallet addresses to be processed.
	Wallets      *domain.Wallets             // Wallet addresses to search against.
	Intervals    *collision.IntervalArray    // Array of collision intervals.
	Results      *output_results.ResultArray // Array of search results.

	CollisionPathFile string // File path for saving collision data.
	ResultPathFile    string // File path for saving result data.
}
