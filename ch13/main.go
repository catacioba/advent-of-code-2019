package ch13

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"time"
)

type Board struct {
	maxHeight int
	minHeight int
	maxWidth  int
	minWidth  int

	tiles map[util.Point]int
	Score int

	BallPosition   util.Point
	PaddlePosition util.Point
}

func newBoard() Board {
	return Board{
		maxHeight: 0,
		minHeight: 0,
		maxWidth:  0,
		minWidth:  0,
		tiles:     make(map[util.Point]int),
		Score:     0,
	}
}

func (board *Board) setTile(x, y, tileId int) {
	point := util.Point{X: x, Y: y}

	board.maxHeight = util.MyMax(board.maxHeight, y)
	board.minHeight = util.MyMin(board.minHeight, y)

	board.maxWidth = util.MyMax(board.maxWidth, x)
	board.minWidth = util.MyMin(board.minWidth, x)

	if tileId == 4 {
		board.BallPosition = point
	} else if tileId == 3 {
		board.PaddlePosition = point
	}
	board.tiles[point] = tileId
}

func (board *Board) draw() {
	fmt.Println("\033[2J")
	fmt.Printf("Score: %d\n", board.Score)
	for y := board.minHeight; y <= board.maxHeight; y++ {
		for x := board.minWidth; x <= board.maxWidth; x++ {
			point := util.Point{
				X: x,
				Y: y,
			}
			tile, _ := board.tiles[point]

			switch tile {
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("*")
			case 3:
				fmt.Print("=")
			case 4:
				fmt.Print("@")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (board *Board) getBallX() int {
	return board.BallPosition.X
}

func (board *Board) readUserInput() int {
	var input string
	fmt.Println("Enter direction:")
	_, _ = fmt.Scanln(&input)

	if input == "h" {
		return -1
	} else if input == "l" {
		return 1
	} else {
		return 0
	}
}

func (board *Board) tryPlay() int {
	xBall := board.getBallX()
	xPaddle := board.PaddlePosition.X

	if xPaddle < xBall {
		return 1
	} else if xPaddle == xBall {
		return 0
	} else {
		return -1
	}
}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch13/input.txt")

	go program.Run()

	count := 0
	for {
		select {
		case _ = <-program.Output:
			_ = <-program.Output
			tile := <-program.Output

			if tile == 2 {
				count += 1
			}
		}

		if !program.IsRunning {
			break
		}
	}

	fmt.Println(count)
}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch13/input.txt")

	program.UpdateMemory(0, 2)

	go program.Run()

	board := newBoard()

	steps := 0
	for {
		steps++

		select {
		case x := <-program.Output:
			y := <-program.Output
			tile := <-program.Output

			if x == -1 && y == 0 {
				board.Score = tile
			} else {
				board.setTile(x, y, tile)
			}
		case <-time.After(2 * time.Millisecond):
			if steps % 1000 == 0 {
				board.draw()
			}
			program.Input <- board.tryPlay()
		}

		if !program.IsRunning {
			fmt.Printf("Final score is: %d\n", board.Score)
			break
		}
	}

	board.draw()
}
