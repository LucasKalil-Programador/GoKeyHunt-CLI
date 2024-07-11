package app_context

import (
	"btcgo/internal/collision"
	"btcgo/internal/domain"
	"btcgo/internal/output_results"
)

type AppCtx struct {
	Params       *domain.Parameters
	WalletRanges *domain.Ranges
	Wallets      *domain.Wallets
	Intervals    *collision.IntervalArray
	Results      *output_results.ResultArray

	CollisionPathFile string
	ResultPathFile    string
}
