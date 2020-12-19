package ch20

import (
	"aoc/util"
	"fmt"
)

type Portal struct {
	code     string
	position util.Point
}

func isLetter(chr byte) bool {
	return chr >= 'A' && chr <= 'Z'
}

func getPortals(board []string) {

	portals := make(map[Portal]bool)

	h := len(board)
	w := len(board[0])

	for x := 2; x < h-2; x++ {
		if isLetter(board[x][0]) {
			portals[Portal{
				code: board[x][:2],
				position: util.Point{
					X: x,
					Y: 3,
				},
			}] = true
		}
		if isLetter(board[x][w-1]) {
			portals[Portal{
				code: board[x][w-2:],
				position: util.Point{
					X: x,
					Y: w - 3,
				},
			}] = true
		}
	}

	fmt.Println(portals)
}

func main() {

	lines := util.ReadLines("ch20/input.txt")

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	getPortals(lines)
}
