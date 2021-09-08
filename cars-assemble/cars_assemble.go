package cars

import (
	"fmt"
	"strconv"
)

const defaultCarNumber float64 = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	// calculate rate per hour as raw float
	rawRatePerHour := successRate(speed) * defaultCarNumber * float64(speed)

	// convert raw float to string in order to get precision by 2
	rawRatePerHourAsString := fmt.Sprintf("%.2f", rawRatePerHour)

	// convert string to float.
	ratePerHour, err := strconv.ParseFloat(rawRatePerHourAsString, 64)

	if err != nil {
		panic(err)
	}

	return ratePerHour
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}

	if speed < 5 {
		return 1.0
	}

	if speed >= 9 {
		return 0.77
	}

	return 0.9
}
