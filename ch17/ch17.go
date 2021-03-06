package ch17

import (
	"aoc/intcode"
	"fmt"
	"time"
)

const (
	scaffold = 35
	open     = 46
	newLine  = 10
)

func drawGrid(grid [][]byte) {
	for _, row := range grid {
		for _, column := range row {
			fmt.Printf("%c", column)
		}
		fmt.Println()
	}
}

func readGrid(channel chan int) [][]byte {
	grid := make([][]byte, 0)
	currentRow := make([]byte, 0)

	for true {
		select {
		case tile := <-channel:
			if tile == 10 {
				grid = append(grid, currentRow)
				currentRow = make([]byte, 0)
			} else {
				currentRow = append(currentRow, byte(tile))
			}
		case <-time.After(100 * time.Millisecond):
			return grid
		}
	}
	panic("")
}

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func isInBounds(x, limit int) bool {
	return x >= 0 && x < limit
}

func markAlignmentParameters(grid [][]byte) int {
	height := len(grid)
	width := len(grid[0])

	result := 0

	for x, row := range grid {
		for y, column := range row {

			ok := column == byte(scaffold)
			for d := 0; d < 4 && ok; d++ {
				newX := x + dx[d]
				newY := y + dy[d]
				if !isInBounds(newX, height) || !isInBounds(newY, width) || grid[newX][newY] != scaffold {
					ok = false
				}
			}

			if ok {
				grid[x][y] = 'O'

				result += x * y
			}
		}
	}

	return result
}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch17/input.txt")

	go program.Run()

	grid := readGrid(program.Output)

	result := markAlignmentParameters(grid)

	drawGrid(grid)

	fmt.Printf("%d", result)
}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch17/input.txt")

	program.UpdateMemory(0, 2)

	go program.Run()

	grid := readGrid(program.Output)

	drawGrid(grid)
}
