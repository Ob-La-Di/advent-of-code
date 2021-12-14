package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	forward Direction = "forward"
	up                = "up"
	down              = "down"
)

type movement struct {
	dir   Direction
	value int
}

func main() {
	movements, err := readFile()

	if err != nil {
		log.Fatal(err.Error())
	}

	horizontalPosition, depth := getFinalDestination(movements)

	println(horizontalPosition, depth, horizontalPosition*depth)

	horizontalPositionWithAim, depthWithAim := getFinalDestinationWithAim(movements)

	println(horizontalPositionWithAim, depthWithAim, horizontalPositionWithAim*depthWithAim)
}

func getFinalDestination(input []movement) (int, int) {
	finalHorizontalPosition, finalDepth := 0, 0

	for _, movementItem := range input {
		switch {
		case movementItem.dir == forward:
			finalHorizontalPosition += movementItem.value
			break
		case movementItem.dir == up:
			finalDepth -= movementItem.value
			break
		case movementItem.dir == down:
			finalDepth += movementItem.value
			break
		}
	}

	return finalHorizontalPosition, finalDepth
}

func getFinalDestinationWithAim(input []movement) (int, int) {
	finalHorizontalPosition, finalDepth, aim := 0, 0, 0

	for _, movementItem := range input {
		switch {
		case movementItem.dir == forward:
			finalHorizontalPosition += movementItem.value
			finalDepth += aim * movementItem.value
			break
		case movementItem.dir == up:
			aim -= movementItem.value
			break
		case movementItem.dir == down:
			aim += movementItem.value
			break
		}
	}

	return finalHorizontalPosition, finalDepth
}

func readFile() ([]movement, error) {
	data, err := os.ReadFile("./day2/input.txt")
	if err != nil {
		return nil, err
	}

	lines := splitFileContent(data)

	result := make([]movement, len(lines))
	for index, item := range lines {
		movementResult, err := readLine(item)
		if err != nil {
			return nil, err
		}

		result[index] = movementResult
	}

	return result, nil
}

func splitFileContent(input []byte) []string {
	return strings.Split(string(input), "\r\n")
}

func readLine(input string) (movement, error) {
	split := strings.Split(input, " ")

	if len(split) != 2 {
		return movement{}, fmt.Errorf("expected only one space in input string")
	}

	var direction Direction
	switch {
	case split[0] == "down":
		direction = down
		break
	case split[0] == "up":
		direction = up
		break
	case split[0] == "forward":
		direction = forward
		break
	default:
		return movement{}, fmt.Errorf("invalid direction: %s", split[0])
	}

	value, err := strconv.Atoi(split[1])

	if err != nil {
		return movement{}, fmt.Errorf("invalid number in value: %s", split[1])
	}

	return movement{
		dir:   direction,
		value: value,
	}, nil
}
