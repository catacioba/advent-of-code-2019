package ch19

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"log"
)

const (
	stationary = 0
	pulled     = 1
)

func deployDrone(x, y int) int {
	program := intcode.NewIntCodeProgramFromFile("ch19/input.txt")

	go program.Run()

	program.Input <- x
	program.Input <- y

	return <-program.Output
}

func drawGrid(grid [][]byte) {
	width := len(grid)
	height := len(grid[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == 2*y {
				fmt.Printf("%s%c%s", util.Red, grid[x][y], util.NoColor)
			} else {
				fmt.Printf("%c", grid[x][y])
			}
		}
		fmt.Println()
	}
}

func PartOne() {
	grid := make([][]byte, 50)
	for x := 0; x < 50; x++ {
		grid[x] = make([]byte, 50)
	}

	grid[0][0] = '#'
	affected := 1
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			if x == 0 && y == 0 {
				continue
			}

			tile := deployDrone(x, y)

			if tile == stationary {
				grid[x][y] = '.'
			} else if tile == pulled {
				affected++
				grid[x][y] = '#'
			} else {
				log.Fatalf("Unexpected token %d", tile)
			}
		}
	}

	drawGrid(grid)
	fmt.Printf("Affected: %d\n", affected)
}

func hasEnoughCorrectTilesOnX(x, startY, size int) bool {
	cnt := 1

	y := x / 2

	if deployDrone(x, y) != pulled || y < startY {
		return false
	}

	// search up
	newY := y - 1
	for cnt < size && newY >= startY {
		if deployDrone(x, newY) == pulled {
			newY--
			cnt++
		} else {
			break
		}
	}

	// search down
	newY = y + 1
	for cnt < size {
		if deployDrone(x, newY) == pulled {
			newY++
			cnt++
		} else {
			break
		}
	}

	return cnt >= size
}

func hasEnoughCorrectTilesOnY(y, startX, size int) bool {
	cnt := 1

	x := y * 2
	if deployDrone(x, y) != pulled || x < startX {
		return false
	}

	// search left
	newX := x - 1
	for cnt < size && newX >= startX {
		if deployDrone(newX, y) == pulled {
			newX--
			cnt++
		} else {
			break
		}
	}

	// search right
	newX = x + 1
	for cnt < size {
		if deployDrone(newX, y) == pulled {
			newX++
			cnt++
		} else {
			break
		}
	}

	return cnt >= size
}

func binarySearch(valid func(int) bool) int {
	start := 0
	end := 100000

	ans := -1

	for start <= end {
		mid := (start + end) / 2

		if valid(mid) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func findPoint(size int) {
	changed := true
	startX := 0
	startY := 0

	for changed {
		changed = false

		newX := binarySearch(
			func(x int) bool {
				return hasEnoughCorrectTilesOnX(x, startY, size)
			})
		fmt.Printf("First x to have at least %d lines after y=%d is %d\n", size, startY, newX)
		if newX != startX {
			startX = newX
			changed = true
		}

		newY := binarySearch(
			func(y int) bool {
				return hasEnoughCorrectTilesOnY(y, startX, size)
			})
		fmt.Printf("First y to have at least %d lines below x=%d is %d\n", size, startX, newY)
		if newY != startY {
			startY = newY
			changed = true
		}
	}

	fmt.Printf("(%d, %d)\n", startX, startY)
	fmt.Println(startX*10000 + startY)
}

func PartTwo() {
	//findPoint(100)

	xStart := 0
	xEnd := 0

	leftBounds := make(map[int]int)
	rightBounds := make(map[int]int)

	leftBounds[0] = 0
	rightBounds[0] = 0

	l := 800

	for y := 1; y < l; y++ {
		for deployDrone(xStart, y) != pulled {
			xStart++
		}
		leftBounds[y] = xStart

		xEnd = xStart + 1
		for deployDrone(xEnd, y) == pulled {
			xEnd++
		}
		rightBounds[y] = xEnd - 1
	}

	//for y := 0; y < l; y++ {
	//	fmt.Printf("y=%d => %d <= x <= %d\n", y, leftBounds[y], rightBounds[y])
	//}

	size := 100
	found := false
	firstX := -1
	firstY := -1

	for y := 0; y <= l-size && !found; y++ {
		if y == 31 {
			fmt.Println()
		}

		left := leftBounds[y]
		right := rightBounds[y]

		if right-left < size {
			continue
		}

		for z := left; z <= right-size+1 && !found; z++ {
			i := leftBounds[y+size-1]
			if i <= z {
				found = true
				firstX = z
				firstY = y
			}
		}
	}

	fmt.Println(firstX*10000 + firstY)
}
