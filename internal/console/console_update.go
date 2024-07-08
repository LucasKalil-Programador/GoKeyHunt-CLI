package console

import (
	"fmt"
	"math/big"
	"time"

	"github.com/dustin/go-humanize"
)

func PrintProgressString(minInt, maxInt, currentInt *big.Int, startTime time.Time) {
	min, max, current := convertToBigFloat(minInt, maxInt, currentInt)
	currentF := new(big.Float).Sub(current, min)
	totalF := new(big.Float).Sub(max, min)

	et := time.Since(startTime).Truncate(time.Second)
	keysPerSecF := new(big.Float).Quo(currentF, big.NewFloat(et.Seconds()))
	keysPerSec, _ := keysPerSecF.Int64()

	etaStr := getETAStr(totalF, currentF, keysPerSecF)
	percentage := CalcPercentage(currentF, totalF)

	fmt.Printf("\r%10sk/s | %10f%% | ET: %10v | ETA: %22v", humanize.Comma(keysPerSec), percentage, et, etaStr)
}

func convertToBigFloat(minInt *big.Int, maxInt *big.Int, currentInt *big.Int) (*big.Float, *big.Float, *big.Float) {
	min := new(big.Float).SetInt(minInt)
	max := new(big.Float).SetInt(maxInt)
	current := new(big.Float).SetInt(currentInt)
	return min, max, current
}

func getETAStr(totalV *big.Float, currentV *big.Float, keysPerSecF *big.Float) string {
	remainF := new(big.Float).Sub(totalV, currentV)
	etaSeconds, _ := new(big.Float).Quo(remainF, keysPerSecF).Int64()
	etaDuration := time.Duration(etaSeconds) * time.Second
	var etaStr string
	if etaSeconds < 315360000 {
		currentTime := time.Now()
		etaTime := currentTime.Add(etaDuration)
		etaStr = humanize.Time(etaTime)
	} else {
		etaStr = "ETA is too long to display"
	}
	return etaStr
}

func CalcPercentage(currentV *big.Float, totalV *big.Float) float64 {
	percentage := new(big.Float).Mul(currentV, big.NewFloat(100))
	percentage.Quo(percentage, totalV)
	result, _ := percentage.Float64()
	return result
}
