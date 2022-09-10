package ch17

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"strings"
	"time"
)

const (
	scaffold = 35
	open     = 46
	newLine  = 10
)

var (
	dx          = []int{-1, 0, 1, 0}
	dy          = []int{0, 1, 0, -1}
	orientation = []byte{'^', '>', 'v', '<'}
)

type Grid struct {
	data   [][]byte
	height int
	width  int

	startX         int
	startY         int
	orientationIdx int
}

func (g *Grid) drawGrid() {
	for _, row := range g.data {
		for _, column := range row {
			fmt.Printf("%c", column)
		}
		fmt.Println()
	}
}

func (g *Grid) initRobot() {
	for x, row := range g.data {
		for y, col := range row {
			for idx, c := range orientation {
				if col == c {
					g.startX = x
					g.startY = y
					g.orientationIdx = idx
				}
			}
		}
	}
	if g.startX == -1 {
		panic("invalid robot position")
	}
}

func getSize(grid [][]byte) (int, int) {
	return len(grid), len(grid[0])
}

func readGrid(channel chan int) *Grid {
	grid := make([][]byte, 0)
	currentRow := make([]byte, 0)

	last := 0
	for {
		tile := <-channel
		if tile == newLine && last == newLine {
			height, width := getSize(grid)

			grid := Grid{
				data:           grid,
				height:         height,
				width:          width,
				startX:         0,
				startY:         0,
				orientationIdx: 0,
			}

			grid.initRobot()

			fmt.Printf("heigth=%d width=%d startX=%d startY=%d orientation=%d\n", grid.height, grid.width, grid.startX, grid.startY, grid.orientationIdx)
			return &grid
		}
		if tile == newLine {
			grid = append(grid, currentRow)
			currentRow = make([]byte, 0)
		} else {
			currentRow = append(currentRow, byte(tile))
		}
		last = tile
	}
}

func isInBounds(x, limit int) bool {
	return x >= 0 && x < limit
}

func (g *Grid) isValidScaffold(x int, y int) bool {
	return isInBounds(x, g.height) && isInBounds(y, g.width) && (g.data[x][y] == scaffold || g.data[x][y] == 'R' || g.data[x][y] == 'L')
}

func (g *Grid) getAlignmentParameters() []util.Point {
	result := make([]util.Point, 0)

	for x, row := range g.data {
		for y, column := range row {
			ok := column == byte(scaffold)

			for d := 0; d < 4 && ok; d++ {
				newX := x + dx[d]
				newY := y + dy[d]

				if !g.isValidScaffold(newX, newY) {
					ok = false
				}
			}

			if ok {
				result = append(result, util.Point{X: x, Y: y})
			}
		}
	}

	return result
}

func (g *Grid) markAlignmentParameters() int {
	result := 0
	alignmentParameters := g.getAlignmentParameters()

	for _, p := range alignmentParameters {
		g.data[p.X][p.Y] = 'O'
		result += p.X * p.Y
	}

	return result
}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch17/input.txt")

	go program.Run()

	grid := readGrid(program.Output)

	result := grid.markAlignmentParameters()

	grid.drawGrid()

	fmt.Printf("%d", result)
}

func (g *Grid) getPath() *paths {
	visited := make(map[util.Point]struct{})

	scaffoldMap := g.getAllScaffolds()

	alignmentParams := g.getAlignmentParameters()
	alignmentParamsMap := make(map[util.Point]struct{})
	for _, p := range alignmentParams {
		alignmentParamsMap[p] = struct{}{}
	}

	startPosition := position{x: g.startX, y: g.startY, orientation: g.orientationIdx, char: orientation[g.orientationIdx]}

	steps := newQueue()
	paths := newPaths()
	g.dfs(startPosition, visited, scaffoldMap, alignmentParamsMap, steps, paths)

	fmt.Printf("Found %d paths\n", len(paths.data))

	return paths
}

func (g *Grid) getAllScaffolds() map[util.Point]struct{} {
	scaffoldMap := make(map[util.Point]struct{})

	for x, row := range g.data {
		for y, col := range row {
			if col == scaffold {
				scaffoldMap[util.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	return scaffoldMap
}

type position struct {
	x           int
	y           int
	orientation int
	char        byte
}

func (g *Grid) getValidMoves(p position) []position {
	result := make([]position, 0)

	for d := -1; d <= 1; d++ {
		idx := p.orientation + d
		if idx < 0 {
			idx = 3
		}
		if idx > 3 {
			idx = 0
		}
		newX := p.x + dx[idx]
		newY := p.y + dy[idx]

		char := p.char
		if d < 0 {
			char = 'L'
		}
		if d > 0 {
			char = 'R'
		}

		if g.isValidScaffold(newX, newY) {
			result = append(result, position{
				x:           newX,
				y:           newY,
				orientation: idx,
				char:        char,
			})
		}
	}

	return result
}

// func (g *Grid) markPositions(q *queue) {
// 	for _, p := range q.data {
// 		g.data[p.x][p.y] = p.char
// 	}
// }

type queue struct {
	data []position
}

func newQueue() *queue {
	return &queue{data: make([]position, 0)}
}

func (q *queue) add(p position) {
	q.data = append(q.data, p)
}

func (q *queue) remove() {
	if len(q.data) == 0 {
		panic("Remove on empty queue")
	}
	q.data = q.data[:len(q.data)-1]
}

func (q *queue) Copy() *queue {
	dataCopy := make([]position, len(q.data))
	copy(dataCopy, q.data)
	return &queue{
		data: dataCopy,
	}
}

type paths struct {
	data []*queue
}

func newPaths() *paths {
	return &paths{make([]*queue, 0)}
}

func (p *paths) addPath(q *queue) {
	p.data = append(p.data, q)
}

func (g *Grid) dfs(p position, visited map[util.Point]struct{}, scaffoldMap map[util.Point]struct{}, alignmentParamsMap map[util.Point]struct{}, steps *queue, s *paths) bool {
	currentPoint := util.Point{X: p.x, Y: p.y}
	delete(scaffoldMap, currentPoint)

	if len(scaffoldMap) == 0 {
		stepsCopy := steps.Copy()
		s.addPath(stepsCopy)
	}

	validMoves := g.getValidMoves(p)

	for _, move := range validMoves {
		movePoint := util.Point{X: move.x, Y: move.y}

		_, isVisited := visited[movePoint]
		_, isAlignmentParam := alignmentParamsMap[movePoint]

		if !isVisited || (isVisited && isAlignmentParam) {
			steps.add(move)
			visited[movePoint] = struct{}{}

			g.dfs(move, visited, scaffoldMap, alignmentParamsMap, steps, s)

			steps.remove()
			delete(visited, movePoint)
		}
	}

	// did not reach solution.
	scaffoldMap[currentPoint] = struct{}{}

	return false
}

// func (g *Grid) resetPosition(p position) {
// 	g.data[p.x][p.y] = scaffold
// }

func getInstructionsStr(steps *queue) string {
	instructionStr := strings.Builder{}

	cnt := 0
	char := byte('@')
	orIdx := -1

	for _, s := range steps.data {
		if char == byte('@') {
			char = s.char
			orIdx = s.orientation
			cnt = 1
		} else {
			if s.orientation != orIdx {
				instructionStr.WriteByte(char)
				instructionStr.WriteByte(',')
				instructionStr.WriteString(fmt.Sprint(cnt))
				instructionStr.WriteByte(',')

				char = s.char
				orIdx = s.orientation
				cnt = 1
			} else {
				cnt++
			}
		}
	}
	instructionStr.WriteByte(char)
	instructionStr.WriteByte(',')
	instructionStr.WriteString(fmt.Sprint(cnt))

	return instructionStr.String()
}

func extract(s string, a string) string {
	ss := strings.ReplaceAll(s, a, "")
	ss = strings.ReplaceAll(ss, ",,", ",")
	if ss[0] == ',' {
		ss = ss[1:]
	}
	if len(ss) > 0 && ss[len(ss)-1] == ',' {
		ss = ss[:]
	}
	return ss
}

func verify(s string, a string, b string, c string) (bool, string) {
	if len(a) > 20 || len(b) > 20 || len(c) > 20 {
		return false, ""
	}
	ss := strings.ReplaceAll(s, a, "A")
	ss = strings.ReplaceAll(ss, b, "B")
	ss = strings.ReplaceAll(ss, c, "C")
	ok := len(ss) <= 20
	if ok {
		return true, ss
	} else {
		return false, ""
	}
}

type solution struct {
	a    string
	b    string
	c    string
	code string
}

func split(instructions string) (*solution, error) {
	aStart := 0
	aEnd := 2

	for aEnd < len(instructions) {
		if instructions[aEnd] == ',' {
			a := instructions[aStart:aEnd]

			s := extract(instructions, a)

			bStart := 0
			bEnd := 2

			for bEnd < len(s) {
				if s[bEnd] == ',' {
					b := s[bStart:bEnd]

					ss := extract(s, b)

					cStart := 0
					cEnd := 2

					for cEnd < len(ss) {
						if ss[cEnd] == ',' {
							c := ss[cStart:cEnd]

							sss := extract(ss, c)

							if len(sss) == 0 {
								ok, str := verify(instructions, a, b, c)
								if ok {
									fmt.Printf("a:%s\nb:%s\nc:%s\ncode:%s\n", a, b, c, str)
									return &solution{a: a, b: b, c: c, code: str}, nil
								}
							}
						}
						cEnd++
					}
				}
				bEnd++
			}
		}
		aEnd++
	}
	return nil, fmt.Errorf("no valid split found")
}

func solve(grid *Grid) *solution {
	paths := grid.getPath()

	for idx := range paths.data {
		steps := paths.data[idx]
		instructions := getInstructionsStr(steps)

		sol, err := split(instructions)
		if err == nil {
			return sol
		}
	}
	panic("no split found")
}

func sendInstruction(ch chan int, s string) {
	for _, x := range s {
		y := int(x)
		ch <- y
	}
	ch <- newLine
}

func read(ch chan int, lines int) []byte {
	d := make([]byte, 0)
	last := 0
	c := 0
	for {
		select {
		case x := <-ch:
			if x == newLine && last == newLine {
				return d
			}
			if x == newLine {
				c++
				if lines > 0 && c%lines == 0 {
					time.Sleep(time.Second / 3)
				}
			}
			d = append(d, byte(x))
			fmt.Printf("%c", x)
			last = x
		case <-time.After(100 * time.Millisecond):
			return d
		}
	}
}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch17/input.txt")

	program.UpdateMemory(0, 2)

	go program.Run()

	grid := readGrid(program.Output)

	grid.drawGrid()
	read(program.Output, 0)

	sol := solve(grid)

	code := sol.code
	a := sol.a
	b := sol.b
	c := sol.c
	// code := "A,B,A,B,A,C,B,C,A,C"
	// a := "L,10,L,12,R,6"
	// b := "R,10,L,4,L,4,L,12"
	// c := "L,10,R,10,R,6,L,4"
	fmt.Println()

	sendInstruction(program.Input, code)
	read(program.Output, 0)

	sendInstruction(program.Input, a)
	read(program.Output, 0)

	sendInstruction(program.Input, b)
	read(program.Output, 0)

	sendInstruction(program.Input, c)
	read(program.Output, 0)

	sendInstruction(program.Input, "n")
	read(program.Output, 0)

	output := <-program.Output
	fmt.Println(output)
}
