package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputIntegers, err := readFile()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	println(getNumberOfMeasurementsGreaterThanPrevious(inputIntegers))
	println(getNumberOfMeasurementsGreaterThanPrevious(groupByMeasurementWindow(inputIntegers, 3)))
}

func readFile() ([]int, error) {
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		return nil, err
	}

	result, err := convertArrayToInt(splitFileContent(data))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func splitFileContent(input []byte) []string {
	return strings.Split(string(input), "\r\n")
}

func convertArrayToInt(input []string) ([]int, error) {
	result := make([]int, len(input))

	for index, item := range input {
		var err error
		result[index], err = strconv.Atoi(item)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func getNumberOfMeasurementsGreaterThanPrevious(input []int) int {
	result := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			result += 1
		}
	}

	return result
}

func groupByMeasurementWindow(input []int, windowSize int) []int {
	var result []int

	currentArrayIndex := 0
	for currentArrayIndex < len(input) {
		if currentArrayIndex+windowSize > len(input) {
			break
		}

		sum := 0
		for i := 0; i < windowSize; i++ {
			sum += input[currentArrayIndex+i]
		}
		result = append(result, sum)
		currentArrayIndex += 1
	}

	return result
}
