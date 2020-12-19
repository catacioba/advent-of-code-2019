package ch18

import (
	"aoc/util"
	"fmt"
)

var keys = []string{}
var doors = []string{}
var path = []string{}

var directionsX = []int{1, 0, -1, 0}
var directionsY = []int{0, 1, 0, -1}

type Board struct {
	board []string
	start util.Point
}

func isInRange(board []string, pos util.Point) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X > len(board) || pos.Y > len(board[0]) {
		return false
	}
	if board[pos.X][pos.Y] == '#' {
		return false
	}
	return true
}

func isDoor(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isKey(b byte) bool {
	return b >= 'a' && b <= 'z'
}

const lowerToUpper = 'a' - 'A'

func toLower(b byte) byte {
	return b - lowerToUpper
}

func toUpper(b byte) byte {
	return b + lowerToUpper
}

func dfs(board []string, startingPoint util.Point) {
	visited := map[util.Point]bool{}

	dfsUtil(board, startingPoint, visited)
}

func dfsUtil(board []string, position util.Point, visited map[util.Point]bool) {
	if !isInRange(board, position) {
		return
	}

	chr := board[position.X][position.Y]

	if isDoor(chr) {
		doors = append(doors, string(chr))
	} else if isKey(chr) {
		keys = append(keys, string(chr))
	}
	if isDoor(chr) || isKey(chr) {
		path = append(path, string(chr))
	}

	for t := 0; t < 4; t++ {
		p := getNextPosition(position, t)

		if !visited[p] {
			visited[p] = true

			dfsUtil(board, p, visited)
		}
	}
}

func topologicalSort(board []string, position util.Point, stack []byte) {
	if !isInRange(board, position) {
		return
	}
}

// type state struct {
// 	pos   util.Point
// 	keys  []byte
// 	gates []byte
// }

// func newState(p util.Point) *state {
// 	return &state{
// 		pos:   p,
// 		keys:  nil,
// 		gates: nil,
// 	}
// }

func (b *Board) bfs() {
	q := util.NewQueue()

	visited := make(map[util.Point]bool)
	dist := make(map[util.Point]int)
	prev := make(map[util.Point]util.Point)

	dist[b.start] = 0
	q.Push(b.start)
	visited[b.start] = true

	for q.Size() > 0 {
		position := q.Pop().(util.Point)

		chr := b.board[position.X][position.Y]

		if isDoor(chr) || isKey(chr) {
			fmt.Printf("# %c is at distance %d\n", chr, dist[position])
		}

		possibleDirections := 0
		for t := 0; t < 4; t++ {
			p := getNextPosition(position, t)

			if isInRange(b.board, p) && !visited[p] {
				possibleDirections++

				visited[p] = true
				dist[p] = dist[position] + 1
				prev[p] = position

				q.Push(p)
			}
		}
		if possibleDirections > 1 {
			fmt.Printf("split occurred at %v\n", position)
		}
	}
}

func getNextPosition(position util.Point, idx int) util.Point {
	return util.Point{
		X: position.X + directionsX[idx],
		Y: position.Y + directionsY[idx],
	}
}

// func solveSlow(board []string, pos util.Point, visited, keys map[util.Point]bool, totalKeys int) int {

// 	minPath := math.MaxInt32

// 	chr := board[pos.X][pos.Y]

// 	if isKey(chr) {

// 	}

// 	for t := 0; t < 4; t++ {
// 		p := getNextPosition(pos, t)

// 		if isInRange(p) {
// 			d :=
// 		}
// 	}
// }

func getStartPosition(lines []string) util.Point {
	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == '@' {
				return util.Point{
					X: x,
					Y: y,
				}
			}
		}
	}
	panic("Start point not found!")
}

func main() {
	lines := util.ReadLines("ch18/input.txt")
	start := getStartPosition(lines)

	// board := Board{board: lines, start: start}
	// board.bfs()

	dfs(lines, start)

	// // for k := range visited {
	// // 	fmt.Println(k)
	// // }

	fmt.Println(path)
	fmt.Println()
	fmt.Println(keys)
	fmt.Println()
	fmt.Println(doors)

	// s := []int{}

	// topologicalSort(board, start, s)

	// fmt.Println(s)
}
