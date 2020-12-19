package ch15

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"log"
)

const (
	// direction based constants
	north = 1
	south = 2
	west  = 3
	east  = 4

	// tile based constants
	wall  = 0
	empty = 1
	final = 2
)

var directions = map[int]util.Point{
	north: {
		X: 0,
		Y: 1,
	},
	south: {

		X: 0,
		Y: -1,
	},
	west: {

		X: -1,
		Y: 0,
	},
	east: {

		X: 1,
		Y: 0,
	},
}
var inverseDirections = map[int]int{
	north: south,
	south: north,
	west:  east,
	east:  west,
}

type robot struct {
	current util.Point
	grid    map[util.Point]int
	prev    map[util.Point]int

	input  chan int
	output chan int
}

func newRobot(programInput chan int, programOutput chan int) *robot {
	current := util.Point{
		X: 0, Y: 0,
	}
	grid := make(map[util.Point]int)
	grid[current] = 1

	return &robot{
		current: current,
		grid:    grid,
		prev:    make(map[util.Point]int),
		input:   programInput,
		output:  programOutput,
	}
}

func (r *robot) nextUnvisitedPosition() (int, util.Point) {
	for dir, point := range directions {
		newDirection := r.current.Add(point)

		_, visited := r.grid[newDirection]

		if !visited {
			return dir, newDirection
		}
	}
	return 0, util.Point{}
}

func (r *robot) goBack() bool {
	prevDirection, found := r.prev[r.current]
	if !found {
		return false
	}

	backDirection := inverseDirections[prevDirection]
	backPosition := r.current.Add(directions[backDirection])

	r.input <- backDirection
	backTile := <-r.output

	knownBackTile := r.grid[backPosition]
	if backTile != knownBackTile {
		log.Fatal("Different known positions when going back!")
	}

	r.current = backPosition

	return true
}

func (r *robot) explore() util.Point {
	var targetPosition util.Point

	steps := 0
	for true {
		steps++

		newDirection, newPosition := r.nextUnvisitedPosition()

		if newDirection == 0 {
			if !r.goBack() {
				return targetPosition
			}
		} else {
			r.input <- newDirection

			tile := <-r.output
			switch tile {
			case wall:
				r.grid[newPosition] = wall
			case empty:
				r.grid[newPosition] = empty
				r.prev[newPosition] = newDirection
				r.current = newPosition
			case final:
				r.grid[newPosition] = final
				r.prev[newPosition] = newDirection
				r.current = newPosition
				targetPosition = newPosition
			}
		}
	}
	return targetPosition
}

const red = "\033[0;31m"
const noColor = "\033[0m"

func drawGrid(grid map[util.Point]int, currentX, currentY int) {
	// find grid bounds.
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for p := range grid {
		minX = util.MyMin(minX, p.X)
		maxX = util.MyMax(maxX, p.X)
		minY = util.MyMin(minY, p.Y)
		maxY = util.MyMax(maxY, p.Y)
	}

	fmt.Printf("bounds: X: (%d, %d) Y: (%d, %d)\n", minX, maxX, minY, maxY)

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			tile, found := grid[util.Point{X: x, Y: y}]

			if x == currentX && y == currentY {
				fmt.Printf("%s", red)
			}
			if x == 0 && y == 0 {
				fmt.Printf("%s%c%s", red, 'X', noColor)
			} else {
				if found {
					switch tile {
					case wall:
						fmt.Printf("%c", '#')
					case empty:
						fmt.Printf("%c", '.')
					case final:
						fmt.Printf("%s%c%s", red, 'o', noColor)
					}
				} else {
					fmt.Printf("%c", ' ')
				}
			}

			if x == currentX && y == currentY {
				fmt.Printf("%s", noColor)
			}
		}
		fmt.Println()
	}
}

// Does a BFS from source and returns the distance to all tiles.
func bfs(grid map[util.Point]int, source util.Point) map[util.Point]int {
	distance := make(map[util.Point]int)
	distance[source] = 0

	queue := util.NewQueue()
	queue.Push(source)

	for !queue.IsEmpty() {
		current := queue.Pop().(util.Point)
		dist := distance[current]

		for _, p := range directions {
			nextPosition := current.Add(p)

			t, found := grid[nextPosition]
			if !found || t == wall {
				continue
			}

			_, visited := distance[nextPosition]
			if !visited {
				queue.Push(nextPosition)
				distance[nextPosition] = dist + 1
			}
		}
	}

	return distance
}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch15/input.txt")

	go program.Run()

	r := newRobot(program.Input, program.Output)
	target := r.explore()

	distances := bfs(r.grid, util.Point{
		X: 0,
		Y: 0,
	})

	dist := distances[target]

	//drawGrid(r.grid, 0, 0)

	fmt.Println(dist)

	//return dist
}

//func TestPartOne(t *testing.T) {
//	got := PartOne()
//	want := 214
//	if got != want {
//		t.Errorf("PartOne() = %d; want %d", got, want)
//	}
//}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch15/input.txt")

	go program.Run()

	r := newRobot(program.Input, program.Output)
	target := r.explore()

	distances := bfs(r.grid, target)

	maxDistance := 0
	for _, dist := range distances {
		if dist > maxDistance {
			maxDistance = dist
		}
	}

	fmt.Println(maxDistance)

	//return maxDistance
}

//func TestPartTwo(t *testing.T) {
//	got := PartTwo()
//	want := 344
//	if got != want {
//		t.Errorf("PartTwo() = %d; want %d", got, want)
//	}
//}
