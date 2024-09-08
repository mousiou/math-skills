package tools

import (
	"math"
	"sort"
)

func Avg(dataSet []float64) float64 {
	sum := 0.0
	for _, v := range dataSet {
		sum += v
	}
	return float64(sum) / float64(len(dataSet))
}

func Median(dataSet []float64) float64 {
	sort.Float64s(dataSet)
	n := len(dataSet)
	if n%2 == 0 {
		return float64(dataSet[n/2-1]+dataSet[n/2]) / 2
	}
	return float64(dataSet[n/2])
}

func Variance(dataSet []float64, avg float64) float64 {
	variance := 0.0
	for _, v := range dataSet {
		variance += math.Pow(float64(v)-avg, 2)
	}
	return variance / float64(len(dataSet))
}
