package tools

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(fileName string) []float64 {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var dataSet []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatalf("failed to parse line: %s", err)
		}
		dataSet = append(dataSet, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	return dataSet
}
