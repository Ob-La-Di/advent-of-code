package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BoardCell struct {
	Number  int
	IsDrawn bool
}

type Board struct {
	hasWon bool
	cells  [][]BoardCell
}

type Game struct {
	randomNumbers []int
	boards        []Board
}

func main() {

	game, _ := readFile()

	println(game.play())
}

func readFile() (Game, error) {
	data, err := os.ReadFile("./day4/input.txt")
	firstLineBreakIndex := bytes.Index(data, []byte{'\r', '\n'})
	if firstLineBreakIndex == -1 {
		return Game{}, fmt.Errorf("expected line break in input file")
	}
	firstLine := data[:firstLineBreakIndex]
	restOfTheFile := data[firstLineBreakIndex:]
	parsedFirstLine, err := parseRandomNumbers(firstLine)
	if err != nil {
		return Game{}, err
	}

	parsedBoards, err := parseBoards(restOfTheFile)

	game := Game{
		randomNumbers: parsedFirstLine,
		boards:        parsedBoards,
	}

	return game, nil
}

func parseRandomNumbers(input []byte) ([]int, error) {
	if len(input) == 0 {
		return []int{}, nil
	}

	splitNumbers := strings.Split(string(input), ",")
	var result = make([]int, len(splitNumbers))
	for index, item := range splitNumbers {
		convertedNumber, err := strconv.Atoi(item)

		if err != nil {
			return nil, err
		}

		result[index] = convertedNumber
	}

	return result, nil
}

func parseBoards(input []byte) ([]Board, error) {
	numberOfBoards := len(input) / (16*5 + 2)
	boards := make([]Board, numberOfBoards)

	for i := 0; i < numberOfBoards; i++ {
		boards[i] = parseBoard(input[i*82+4 : i*82+4+78])
	}

	return boards, nil
}

func parseBoard(input []byte) Board {
	var cells [][]BoardCell
	for i := 0; i < 5; i++ {
		cells = append(cells, parseLine(input[i*14+i*2:i*14+i*2+14]))
	}

	return Board{
		cells: cells,
	}
}

func parseLine(input []byte) []BoardCell {
	result := make([]BoardCell, 5)
	re := regexp.MustCompile(" +")
	split := re.Split(strings.Trim(string(input), " "), -1)

	for index, item := range split {
		number, _ := strconv.Atoi(item)

		result[index] = BoardCell{Number: number, IsDrawn: false}
	}

	return result
}

func (g *Game) play() int {
	var lastWinnerBoard *Board
	var lastDrawnWinningNumber int
	numberOfWinningBoards := 0
	for _, drawnNumber := range g.randomNumbers {
		for index, board := range g.boards {
			if board.hasWon == false {
				g.boards[index].setNumberDraw(drawnNumber)
			}

			if board.isWinner() && board.hasWon == false {
				numberOfWinningBoards++
				g.boards[index].hasWon = true
				lastWinnerBoard = &board
				lastDrawnWinningNumber = drawnNumber
			}

			if numberOfWinningBoards == len(g.boards) {
				return board.getScore(drawnNumber)
			}
		}
	}

	return (*lastWinnerBoard).getScore(lastDrawnWinningNumber)

}

func (b *Board) setNumberDraw(number int) {
	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j].Number == number {
				b.cells[i][j].IsDrawn = true
			}
		}
	}
}

func (b *Board) isWinner() bool {
	for i := 0; i < 5; i++ {
		allDrawn := true
		for j := 0; j < 5; j++ {
			if b.cells[i][j].IsDrawn == false {
				allDrawn = false
			}
		}
		if allDrawn == true {
			return true
		}

		allDrawn = true
		for j := 0; j < 5; j++ {
			if b.cells[j][i].IsDrawn == false {
				allDrawn = false
			}
		}
		if allDrawn == true {
			return true
		}

	}

	return false
}

func (b *Board) getScore(justCalledNumber int) int {
	total := 0

	for _, cellLines := range b.cells {
		for _, cell := range cellLines {
			if !cell.IsDrawn {
				total += cell.Number
			}
		}
	}

	return justCalledNumber * total
}
