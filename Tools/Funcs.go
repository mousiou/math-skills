package tools

import (
	"math"
	"sort"
)

// Avg calculates the average (mean) of a slice of float64 numbers.
func Avg(dataSet []float64) float64 {
	sum := 0.0
	for _, v := range dataSet {
		sum += v
	}
	return sum / float64(len(dataSet))
}

// Median calculates the median of a slice of float64 numbers.
func Median(dataSet []float64) float64 {
	sort.Float64s(dataSet)
	ln := len(dataSet)
	if ln%2 == 0 {
		return float64(dataSet[ln/2-1]+dataSet[ln/2]) / 2
	}
	return float64(dataSet[ln/2])
}

// Variance calculates the variance of a slice of float64 numbers based on the provided average.
func Variance(dataSet []float64, avg float64) float64 {
	variance := 0.0
	for _, val := range dataSet {
		variance += math.Pow(val-avg, 2)
	}
	// Return the variance by dividing the sum of squared differences by the number of elements in the dataSet.
	return variance / float64(len(dataSet))
}
