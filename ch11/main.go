package ch11

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"math"
	"strings"
)

/*
	rotate left:
		up		(0, 1) 	=> 	(-1, 0)		left
		right 	(1, 0) 	=> 	(0, 1) 		up
		down 	(0, -1) =>	(1, 0)		right
		left 	(-1, 0) =>	(0, -1)		down

	rotate right
		up		(0, 1)	=> 	(1, 0)		right
		right 	(1, 0) 	=> 	(0, -1)		down
		down	(0, -1)	=>	(-1, 0)		left
		left	(-1, 0)	=>	(0, 1)		up
*/

var rotateLeftMap = map[util.Point]util.Point{
	{0, 1}:  {-1, 0},
	{1, 0}:  {0, 1},
	{0, -1}: {1, 0},
	{-1, 0}: {0, -1},
}

var rotateRightMap = map[util.Point]util.Point{
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
	{0, -1}: {-1, 0},
	{-1, 0}: {0, 1},
}

func main() {

	line := util.ReadLines("ch11/input.txt")[0]
	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))
	numbersBigger := make([]int, 100000)
	copy(numbersBigger, numbers)

	program := intcode.NewIntCodeProgram(numbersBigger)

	go program.Run()

	position := util.Point{
		X: 0,
		Y: 0,
	}

	direction := util.Point{
		X: 0,
		Y: 1,
	}

	white := make(map[util.Point]bool)
	white[position] = true

	for {
		isWhite, ok := white[position]

		if ok == false {
			isWhite = false
		}

		if isWhite {
			program.Input <- 1
		} else {
			program.Input <- 0
		}

		color := <-program.Output
		if color == 1 {
			white[position] = true
		} else {
			white[position] = false
		}

		leftRotate := <-program.Output

		if leftRotate == 0 {
			direction = rotateLeftMap[direction]
		} else {
			direction = rotateRightMap[direction]
		}

		position.X += direction.X
		position.Y += direction.Y

		if !program.IsRunning {
			break
		}
	}

	fmt.Println(len(white))

	minHeight := math.MaxInt32
	minWidth := math.MaxInt32
	maxHeight := math.MinInt32
	maxWidth := math.MinInt32

	for k, v := range white {
		if v == true {
			minHeight = util.MyMin(minHeight, k.Y)
			maxHeight = util.MyMax(maxHeight, k.Y)

			minWidth = util.MyMin(minWidth, k.X)
			maxWidth = util.MyMax(maxWidth, k.X)
		}
	}

	fmt.Printf("X: %d %d\n", minWidth, maxWidth)
	fmt.Printf("Y: %d %d\n", minHeight, maxHeight)

	for idy := 0; idy >= -5; idy-- {
		for idx := 1; idx <= 39; idx++ {
			isWhite, ok := white[util.Point{X: idx, Y: idy}]

			if ok == false {
				isWhite = false
			}

			if isWhite {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
