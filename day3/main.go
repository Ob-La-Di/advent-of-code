package main

import (
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := readFile()

	if err != nil {
		log.Fatal(err.Error())
	}

	gammaRate, epsilonRate := rates(data)
	oxygen, co2 := getOxygenGeneratorAndCO2ScrubberRatings(data)

	println(gammaRate, epsilonRate, gammaRate*epsilonRate)
	println(oxygen, co2, oxygen*co2)
}

func readFile() ([]string, error) {
	data, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		return nil, err
	}

	return splitFileContent(data), nil
}

func splitFileContent(input []byte) []string {
	return strings.Split(string(input), "\r\n")
}

func rates(input []string) (int, int) {
	bytesLength := len(input[0])
	numberOfItemsInArray := len(input)
	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < bytesLength; i++ {
		numberOfOneBit := 0
		for _, item := range input {
			if item[i] == '1' {
				numberOfOneBit += 1
			}
		}

		if numberOfOneBit > numberOfItemsInArray/2 {
			gammaRate += int(math.Pow(2, float64(bytesLength-(i+1))))
		} else {
			epsilonRate += int(math.Pow(2, float64(bytesLength-(i+1))))
		}
	}

	return gammaRate, epsilonRate
}

func getOxygenGeneratorAndCO2ScrubberRatings(input []string) (int, int) {
	bytesLength := len(input[0])
	remainingOxygeneratorItems := make([]string, len(input))
	copy(remainingOxygeneratorItems, input)
	remainingCO2ScrubberItems := make([]string, len(input))
	copy(remainingCO2ScrubberItems, input)

	for i := 0; i < bytesLength; i++ {
		numberOfOneBitOxygen := 0
		for _, item := range remainingOxygeneratorItems {
			if item[i] == '1' {
				numberOfOneBitOxygen += 1
			}
		}
		numberOfOneBitCO2 := 0
		for _, item := range remainingCO2ScrubberItems {
			if item[i] == '0' {
				numberOfOneBitCO2 += 1
			}
		}
		if len(remainingOxygeneratorItems) > 1 {
			if numberOfOneBitOxygen >= len(remainingOxygeneratorItems)/2 {
				removeItemWithBitInPosition(&remainingOxygeneratorItems, i, '0')
			} else {
				removeItemWithBitInPosition(&remainingOxygeneratorItems, i, '1')
			}
			println(remainingOxygeneratorItems)
		}

		if len(remainingCO2ScrubberItems) > 1 {
			if numberOfOneBitCO2 <= len(remainingCO2ScrubberItems)/2 {
				removeItemWithBitInPosition(&remainingCO2ScrubberItems, i, '1')
			} else {
				removeItemWithBitInPosition(&remainingCO2ScrubberItems, i, '0')
			}
			println(remainingCO2ScrubberItems)
		}
	}

	return binaryToInt(remainingCO2ScrubberItems[0]), binaryToInt(remainingOxygeneratorItems[0])
}

func removeItemWithBitInPosition(input *[]string, position int, bit uint8) {
	inputLength := len(*input)

	for i := 0; i < inputLength; i++ {
		if (*input)[i][position] == bit {
			*input = append((*input)[:i], (*input)[i+1:]...)
			inputLength -= 1
			i -= 1
		}
	}
}

func binaryToInt(input string) int {
	result := 0
	inputLength := len(input)

	for index, char := range input {
		if char == '1' {
			result += int(math.Pow(2, float64(inputLength-index-1)))
		}
	}

	return result
}
