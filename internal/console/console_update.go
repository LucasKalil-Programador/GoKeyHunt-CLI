package console

import (
	"fmt"
	"math/big"
	"time"

	"github.com/dustin/go-humanize"
)

// PrintProgressString prints the progress of a task to the console.
//
// This function calculates and displays the progress of a task based on the minimum, maximum, and current values
// of a given range. It shows the number of keys processed per second, the percentage of completion, the elapsed time,
// and the estimated time of arrival (ETA) for task completion.
//
// Parameters:
// - minInt: A *big.Int representing the starting value of the range.
// - maxInt: A *big.Int representing the ending value of the range.
// - currentInt: A *big.Int representing the current value in the range.
// - startTime: A time.Time representing the start time of the task.
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

// convertToBigFloat converts *big.Int values to *big.Float values.
//
// This function is used to convert minimum, maximum, and current *big.Int values to *big.Float values for calculations.
//
// Parameters:
// - minInt: A *big.Int representing the minimum value.
// - maxInt: A *big.Int representing the maximum value.
// - currentInt: A *big.Int representing the current value.
//
// Returns:
// - *big.Float: The minimum value as a *big.Float.
// - *big.Float: The maximum value as a *big.Float.
// - *big.Float: The current value as a *big.Float.
func convertToBigFloat(minInt, maxInt, currentInt *big.Int) (*big.Float, *big.Float, *big.Float) {
	min := new(big.Float).SetInt(minInt)
	max := new(big.Float).SetInt(maxInt)
	current := new(big.Float).SetInt(currentInt)
	return min, max, current
}

// getETAStr calculates and returns the estimated time of arrival (ETA) as a string.
//
// This function computes the ETA based on the total value, current value, and the number of keys processed per second.
// If the ETA exceeds a reasonable limit, it returns a message indicating that the ETA is too long to display.
//
// Parameters:
// - totalV: A *big.Float representing the total value of the range.
// - currentV: A *big.Float representing the current value in the range.
// - keysPerSecF: A *big.Float representing the number of keys processed per second.
//
// Returns:
// - string: The estimated time of arrival as a human-readable string.
func getETAStr(totalV, currentV, keysPerSecF *big.Float) string {
	remainF := new(big.Float).Sub(totalV, currentV)
	etaSeconds, _ := new(big.Float).Quo(remainF, keysPerSecF).Int64()
	etaDuration := time.Duration(etaSeconds) * time.Second
	var etaStr string
	if etaSeconds < 315360000 { // Less than 10 years
		currentTime := time.Now()
		etaTime := currentTime.Add(etaDuration)
		etaStr = humanize.Time(etaTime)
	} else {
		etaStr = "ETA is too long to display"
	}
	return etaStr
}

// CalcPercentage calculates the percentage of completion based on current and total values.
//
// This function computes the percentage of progress made towards completion of the task.
//
// Parameters:
// - currentV: A *big.Float representing the current value in the range.
// - totalV: A *big.Float representing the total value of the range.
//
// Returns:
// - float64: The percentage of completion as a float64 value.
func CalcPercentage(currentV, totalV *big.Float) float64 {
	percentage := new(big.Float).Mul(currentV, big.NewFloat(100))
	percentage.Quo(percentage, totalV)
	result, _ := percentage.Float64()
	return result
}
