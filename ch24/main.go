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

var dx = []int8{-1, 0, 1, 0}
var dy = []int8{0, -1, 0, 1}

func (g *grid) adjacentBugs(i, j int) uint8 {
	cnt := uint8(0)
	for k := range dx {
		x := int8(i) + dx[k]
		y := int8(j) + dy[k]

		if x >= 0 && x < 5 && y >= 0 && y < 5 {
			if g[x][y] == bug {
				cnt++
			}
		}
	}
	return cnt
}

func (g *grid) iterate() grid {
	var ng [5][5]uint8

	for idx := 0; idx < 5; idx++ {
		for idy := 0; idy < 5; idy++ {
			cnt := g.adjacentBugs(idx, idy)

			if g[idx][idy] == bug {
				if cnt != 1 {
					ng[idx][idy] = empty
				} else {
					ng[idx][idy] = bug
				}
			} else {
				if cnt == 1 || cnt == 2 {
					ng[idx][idy] = bug
				} else {
					ng[idx][idy] = empty
				}
			}
		}
	}

	return ng
}

func biodiversityRating(g grid) uint64 {
	pow := uint64(1)
	h := uint64(0)

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			var v uint64
			if g[x][y] == bug {
				v = 1
			} else {
				v = 0
			}
			h += pow * v
			pow *= 2
		}
	}

	return h
}

func PartOne() {
	// 	g := newGrid(
	// 		`....#
	// #..#.
	// #..##
	// ..#..
	// #....`)
	g := newGrid(
		`.##..
##.#.
##.##
.#..#
#.###`)

	g.print()
	fmt.Println()

	visited := make(map[grid]struct{})
	visited[g] = struct{}{}

	for {
		g = g.iterate()

		_, ok := visited[g]
		if ok {
			g.print()
			break
		}

		visited[g] = struct{}{}
	}

	fmt.Println(biodiversityRating(g))
}

func PartTwo() {}
