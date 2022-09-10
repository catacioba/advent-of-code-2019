package ch24

import (
	"fmt"
	"strings"
)

const (
	bug   = '#'
	empty = '.'
)

type grid [5][5]uint8

func newGrid(g string) grid {
	var grid [5][5]uint8

	lines := strings.Split(g, "\n")

	for i := range lines {
		for j := range lines[i] {
			grid[i][j] = lines[i][j]
		}
	}

	return grid
}

func (g *grid) print() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%c", g[i][j])
		}
		fmt.Println()
	}
}

var dx = []int8{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int8{-1, 0, 1, -1, 1, -1, 0, 1}

func (g *grid) adjacentBugs(i, j int8) uint8 {
	cnt := uint8(0)
	for k := range dx {
		x := i + dx[k]
		y := j + dy[k]

		if x >= 0 && x < 5 && y >= 0 && y < 5 {
			if g[x][y] == bug {
				cnt++
			}
		}
	}
	return cnt
}

type game struct {
	
}

func (g *grid) iterate() [5][5]uint8 {
	var ng [5][5]uint8

	return ng
}

func PartOne() {
	g := newGrid(
		`....#
#..#.
#..##
..#..
#....`)

	g.print()

	// for {
	// 	var ng []
	// }
}

func PartTwo() {}
