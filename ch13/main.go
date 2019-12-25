package main

import (
	"adventofcode/intcode"
	"adventofcode/util"
	"fmt"
	"strings"
	"time"
)

type Board struct {
	maxHeight int
	minHeight int
	maxWidth  int
	minWidth  int

	tiles  map[util.Point]int
	blocks int
	Score  int

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
		blocks:    0,
		Score:     0,
	}
}

func (board *Board) setTile(x, y, tileId int) {
	point := util.Point{X: x, Y: y}

	existingId, found := board.tiles[point]

	if found {
		//fmt.Printf("Tile %v changed from %d to %d\n", point, existingId, tileId)
		if existingId == 2 && tileId != 2 {
			board.blocks--
		}
		if existingId != 2 && tileId == 2 {
			board.blocks++
		}
	} else {
		//fmt.Printf("Tile %v set to %d\n", point, tileId)
		board.maxHeight = util.MyMax(board.maxHeight, y)
		board.minHeight = util.MyMin(board.minHeight, y)

		board.maxWidth = util.MyMax(board.maxWidth, x)
		board.minWidth = util.MyMin(board.minWidth, x)

		if tileId == 2 {
			board.blocks++
		}
	}
	if tileId == 4 {
		board.BallPosition = point
	} else if tileId == 3 {
		board.PaddlePosition = point
	}
	board.tiles[point] = tileId
}

func (board *Board) getBlocks() int {
	return board.blocks
}

func (board *Board) draw() {
	fmt.Println("\033[2J")
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()

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

///*
//	(0,0) 	=> 	stays
//	(1,1) 	=> 	left to right down
//	(-1,-1) => 	left to right up
//	(-1,1)	=> 	right to left down
//	(1,-1)	=>	right to left up
//	-1 		=> 	invalid
//*/
//func getBallDirection(pos1, pos2 util.Point) {
//	if pos1 == pos2 {
//		return 0
//	}
//	if pos1.X == pos
//}

func (board *Board) getBallX() int {
	return board.BallPosition.X
}

func main() {
	line := util.ReadLines("ch13/input.txt")[0]
	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))
	numbersBigger := make([]int, 100000)
	copy(numbersBigger, numbers)

	numbersBigger[0] = 2

	program := intcode.NewIntCodeProgram(numbersBigger)

	go program.Run()
	//program.IsRunning = true
	//program.RunUntilIO()

	board := newBoard()

	steps := 0
	//previousBallPosition := board.BallPosition

	for {
		steps++

		if steps >= 1051 {
			board.draw()
			time.Sleep(1 * time.Second)
		}

		select {
		case x := <-program.Output:
			y := <-program.Output
			tile := <-program.Output

			if x == -1 && y == 0 {
				board.Score = tile
			} else {
				board.setTile(x, y, tile)
			}
		//fmt.Printf("%d %d %d\n", x, y, tile)
		default:
			//ballPosition := board.BallPosition
			xBall := board.getBallX()
			xPaddle := board.PaddlePosition.X

			if xPaddle < xBall {
				program.Input <- 1
			} else if xPaddle == xBall {
				program.Input <- 0
			} else {
				program.Input <- -1
			}

		}

		if !program.IsRunning {
			break
		}
	}

	fmt.Println(board.getBlocks())

}
