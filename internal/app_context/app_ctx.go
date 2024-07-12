package app_context

import (
	"GoKeyHunt/internal/collision"
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/output_results"
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
