package ch24

import (
	"aoc/util"
	"fmt"
	"math"
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
var rd = [][2]util.Point{
	{{4, 5}, {0, 5}},
	{{0, 5}, {4, 5}},
	{{0, 1}, {0, 5}},
	{{0, 5}, {0, 1}},
}

func (g *grid) adjacentBugs(i, j int) int {
	cnt := 0
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

func getNewValue(c uint8, cnt int) uint8 {
	if c == bug {
		if cnt != 1 {
			return empty
		} else {
			return bug
		}
	} else {
		if cnt == 1 || cnt == 2 {
			return bug
		} else {
			return empty
		}
	}
}

func (g *grid) iterate() grid {
	var ng [5][5]uint8

	for idx := 0; idx < 5; idx++ {
		for idy := 0; idy < 5; idy++ {
			cnt := g.adjacentBugs(idx, idy)

			ng[idx][idy] = getNewValue(g[idx][idy], cnt)
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

type recursiveGrid struct {
	grids    map[int]grid
	minLevel int
	maxLevel int
}

func newRecursiveGrid(g string) *recursiveGrid {
	grids := make(map[int]grid)
	grids[0] = newGrid(g)

	return &recursiveGrid{
		grids:    grids,
		minLevel: 0,
		maxLevel: 0,
	}
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

func (g *recursiveGrid) get(i, j int, lvl int) uint8 {
	grid, ok := g.grids[lvl]
	if !ok {
		return empty
	}
	if i < 0 || i >= 5 || j < 0 || j >= 5 {
		return empty
	}
	return grid[i][j]
}

func (g *recursiveGrid) adjacentBugs(i, j, lvl int) int {
	cnt := 0
	for k := range dx {
		x := i + int(dx[k])
		y := j + int(dy[k])

		if x < 0 || x > 4 || y < 0 || y > 4 {
			x = 2 + int(dx[k])
			y = 2 + int(dy[k])
			if g.get(x, y, lvl-1) == bug {
				cnt++
			}
		} else if x == 2 && y == 2 {
			l := rd[k]
			for xx := l[0].X; xx < l[0].Y; xx++ {
				for yy := l[1].X; yy < l[1].Y; yy++ {
					//fmt.Printf("get(%d, %d, %d) => %c\n", xx, yy, lvl-1, g.get(xx, yy, lvl-1))
					if g.get(xx, yy, lvl+1) == bug {
						cnt++
					}
				}
			}
		} else {
			if g.get(x, y, lvl) == bug {
				cnt++
			}
		}
	}
	return cnt
}

func (g *recursiveGrid) iterate() {
	newGrids := make(map[int]grid)
	minLevel := math.MaxInt
	maxLevel := math.MinInt

	for lvl := g.minLevel - 1; lvl <= g.maxLevel+1; lvl++ {
		var ng grid

		ok := false
		for idx := 0; idx < 5; idx++ {
			for idy := 0; idy < 5; idy++ {
				if idx == 2 && idy == 2 {
					ng[idx][idy] = '?'
					continue
				}
				cnt := g.adjacentBugs(idx, idy, lvl)
				//fmt.Printf("At lvl %d => (%d, %d) %d adj bugs\n", lvl, idx, idy, cnt)

				newValue := getNewValue(g.get(idx, idy, lvl), cnt)
				if newValue == bug {
					ok = true
				}
				ng[idx][idy] = newValue
			}
		}

		if ok {
			newGrids[lvl] = ng
			minLevel = util.MyMin(minLevel, lvl)
			maxLevel = util.MyMax(maxLevel, lvl)
		}
	}

	g.grids = newGrids
	g.minLevel = minLevel
	g.maxLevel = maxLevel
}

func (g *recursiveGrid) bugs() int {
	cnt := 0
	for _, grid := range g.grids {
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if grid[x][y] == bug {
					cnt++
				}
			}
		}
	}
	return cnt
}

func (g *recursiveGrid) print() {
	for lvl := g.minLevel; lvl <= g.maxLevel; lvl++ {
		fmt.Printf("Depth %d:\n", lvl)

		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				fmt.Printf("%c", g.grids[lvl][i][j])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func PartTwo() {

	//	g := newRecursiveGrid(`....#
	//#..#.
	//#.?##
	//..#..
	//#....`)
	g := newRecursiveGrid(`.##..
##.#.
##.##
.#..#
#.###
`)

	for idx := 0; idx < 200; idx++ {
		g.iterate()
	}

	//g.print()
	fmt.Println(g.bugs())
}
