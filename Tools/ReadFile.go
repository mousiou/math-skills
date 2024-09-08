package tools

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(fileName string) []float64 {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var dataSet []float64

	// Create a new scanner to read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatalf("failed to parse line: %s", err)
		}
		dataSet = append(dataSet, value)
	}

	// Check if there was an error during scanning.
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	// Return the slice containing all the numbers read from the file.
	return dataSet
}
