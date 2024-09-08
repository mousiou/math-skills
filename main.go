package main

import (
	"fmt"
	"log"
	"math"
	"os"

	tools "Tools/Tools"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: go run main.go data.txt")
	}
	filename := os.Args[1]
	dataSet := tools.ReadFile(filename)
	if len(dataSet) == 0 {
		log.Fatalf("Error: file is empty or contains no valid data")
	}

	avg := tools.Avg(dataSet)
	med := tools.Median(dataSet)
	vrn := tools.Variance(dataSet, avg)
	stdDev := math.Sqrt(vrn)

	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med)))
	fmt.Printf("Variance: %d\n", int(math.Round(vrn)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
