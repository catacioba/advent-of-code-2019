package main

import (
	"adventofcode/util"
	"fmt"
)

var keys = []string{}
var doors = []string{}
var visited = map[util.Point]bool{}

var directionsX = []int{1, 0, -1, 0}
var directionsY = []int{0, 1, 0, -1}

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

func dfs(board []string, position util.Point) {
	if !isInRange(board, position) {
		return
	}

	chr := board[position.X][position.Y]

	if isDoor(chr) {
		doors = append(doors, string(chr))
	} else if isKey(chr) {
		keys = append(keys, string(chr))
	}

	for t := 0; t < 4; t++ {
		p := util.Point{
			X: position.X + directionsX[t],
			Y: position.Y + directionsY[t],
		}
		if !visited[p] {
			visited[p] = true

			dfs(board, p)
		}
	}
}

func main() {
	lines := util.ReadLines("ch18/input.txt")

	var start util.Point

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == '@' {
				start = util.Point{
					X: x,
					Y: y,
				}
				break
			}
		}
	}

	dfs(lines, start)

	// for k := range visited {
	// 	fmt.Println(k)
	// }

	fmt.Println(keys)

	fmt.Println()

	fmt.Println(doors)
}