package main

import (
	tools "Tools/Tools"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Error in arguments")
	}
	filename := os.Args[1]
	dataSet := tools.ReadFile(filename)
	if len(dataSet) == 0 {
		log.Fatalf("error: file is empty or contains no valid data")
	}

	avg := tools.Avg(dataSet)
	med := tools.Median(dataSet)
	vrn := tools.Variance(dataSet, avg)
	stdDev := math.Sqrt(float64(vrn))

	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med)))
	fmt.Printf("Variance: %d\n", int(math.Round(vrn)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
