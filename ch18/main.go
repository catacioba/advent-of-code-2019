package ch18

import (
	"aoc/util"
	"fmt"
	"strings"
)

const startSymbol = '@'

var directionsX = []int{1, 0, -1, 0}
var directionsY = []int{0, 1, 0, -1}

type section struct {
	start       util.Point
	keys        []byte
	gates       []byte
	allKeysMask keyState
}

func (s *section) setMaskState() {
	for _, key := range s.keys {
		k := toIntFromLower(key)
		s.allKeysMask |= keyState(1) << k
	}
}

type Board struct {
	board        [][]byte
	sections     []section
	sectionCount int
}

func (b *Board) initBoard() {
	visited := make(map[util.Point]struct{})

	for x := range b.board {
		for y := range b.board[x] {
			p := util.Point{
				X: x,
				Y: y,
			}
			if b.isInRange(p) {
				_, vis := visited[p]
				if !vis {
					sec := section{}
					b.dfs(p, visited, &sec)
					sec.setMaskState()
					b.sections = append(b.sections, sec)
					b.sectionCount++
				}
			}
		}
	}
}

func newBoard(lines []string) *Board {
	bytes := make([][]byte, len(lines))
	for idx := 0; idx < len(lines); idx++ {
		bytes[idx] = []byte(lines[idx])
	}
	b := Board{board: bytes, sectionCount: 0}
	b.initBoard()

	fmt.Printf("Found %d sections\n", b.sectionCount)

	return &b
}

func (b *Board) isInRange(pos util.Point) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X >= len(b.board) || pos.Y >= len(b.board[0]) {
		return false
	}
	if b.board[pos.X][pos.Y] == '#' {
		return false
	}
	return true
}

func (b *Board) position(pos util.Point) byte {
	return b.board[pos.X][pos.Y]
}

func (b *Board) dfs(position util.Point, visited map[util.Point]struct{}, sec *section) {
	if !b.isInRange(position) {
		return
	}

	chr := b.position(position)

	if chr == startSymbol {
		sec.start = position
	} else if isKey(chr) {
		sec.keys = append(sec.keys, chr)
	} else if isDoor(chr) {
		sec.gates = append(sec.gates, toLower(chr))
	}

	for t := 0; t < 4; t++ {
		p := getNextPosition(position, t)

		_, vis := visited[p]
		if !vis {
			visited[p] = struct{}{}

			b.dfs(p, visited, sec)
		}
	}
}

func isPointOfInterest(b byte) bool {
	return isKey(b) || isDoor(b)
}

func (b *Board) bfs(start util.Point) map[util.Point]CompactedDistance {
	res := make(map[util.Point]CompactedDistance)

	dist := make(map[util.Point]CompactedDistance)
	queue := util.NewQueue()

	dist[start] = CompactedDistance{
		dist: 0,
		req:  keyState(0),
	}
	queue.Push(start)

	for !queue.IsEmpty() {
		el := queue.Pop().(util.Point)
		cd := dist[el]

		for t := 0; t < 4; t++ {
			p := getNextPosition(el, t)

			if !b.isInRange(p) {
				continue
			}

			if _, vis := dist[p]; !vis {
				c := b.position(p)

				nextReq := cd.req

				if isDoor(c) {
					nextReq.setKey(toLower(c))
				}

				nextCd := CompactedDistance{
					dist: cd.dist + 1,
					req:  nextReq,
				}

				dist[p] = nextCd
				queue.Push(p)

				if isPointOfInterest(c) {
					res[p] = nextCd
				}
			}
		}
	}

	return res
}

func (b *Board) gateAwareBfs(start util.Point) map[util.Point]int {
	res := make(map[util.Point]int)

	dist := make(map[util.Point]int)
	queue := util.NewQueue()

	dist[start] = 0
	queue.Push(start)

	for !queue.IsEmpty() {
		el := queue.Pop().(util.Point)
		d := dist[el]

		for t := 0; t < 4; t++ {
			p := getNextPosition(el, t)

			c := b.position(p)

			if !b.isInRange(p) {
				continue
			}

			if _, vis := dist[p]; !vis {
				dist[p] = d + 1

				if isPointOfInterest(c) {
					res[p] = d + 1
				} else {
					queue.Push(p)
				}
			}
		}
	}

	return res
}

func (b *Board) detectCycle(start util.Point) {
	visited := make(map[util.Point]struct{})

	hasCycle := b.detectCycleAux(start, visited, nil)

	fmt.Printf("has cycle: %v\n", hasCycle)
}

func (b *Board) detectCycleAux(p util.Point, visited map[util.Point]struct{}, parent *util.Point) bool {
	if !b.isInRange(p) {
		return false
	}

	visited[p] = struct{}{}

	for t := 0; t < 4; t++ {
		n := getNextPosition(p, t)

		if parent != nil && *parent == n {
			continue
		}

		_, ok := visited[n]
		if !ok {
			//hasCycle = hasCycle || b.detectCycleAux(n, visited, p)
			if b.detectCycleAux(n, visited, &p) {
				return true
			}
		} else {
			return true
		}
	}

	return false
}

func getNextPosition(position util.Point, idx int) util.Point {
	return util.Point{
		X: position.X + directionsX[idx],
		Y: position.Y + directionsY[idx],
	}
}

func getNextMultiSectionState(s multiSectionState, secIdx int, dirIdx int) (util.Point, multiSectionState) {
	nextPos := getNextPosition(s.sectionStates[secIdx].pos, dirIdx)

	var nextSectionStates [4]sectionState
	for i := 0; i < 4; i++ {
		if i == secIdx {
			nextSectionStates[i] = sectionState{
				pos:     nextPos,
				section: s.sectionStates[secIdx].section,
			}
		} else {
			nextSectionStates[i] = s.sectionStates[i]
		}
	}
	return nextPos, multiSectionState{
		sectionStates: nextSectionStates,
		keys:          s.keys,
	}
}

func (b *Board) print() {
	fmt.Println()
	for _, r := range b.board {
		for _, c := range r {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

type keyState uint32

type sectionState struct {
	pos     util.Point
	section *section
}

func (s *keyState) setKey(key byte) {
	k := toIntFromLower(key)
	*s = *s | (keyState(1) << k)
}

func (s *keyState) hasKey(key byte) bool {
	k := toIntFromLower(key)
	return (*s & (keyState(1) << k)) != 0
}

func (s *keyState) hasRequiredKeys(other keyState) bool {
	return (*s & other) == other
}

func (s *keyState) getKeys() string {
	var ks []string

	for c := byte('a'); c <= 'z'; c++ {
		if s.hasKey(c) {
			ks = append(ks, string([]byte{c}))
		}
	}

	return strings.Join(ks, ",")
}

func (s *section) hasAllKeys(ks keyState) bool {
	return (ks & s.allKeysMask) == s.allKeysMask
}

type multiSectionState struct {
	sectionStates [4]sectionState
	keys          keyState
}

func (mss *multiSectionState) getState() visitedState {
	vs := visitedState{}
	for idx := 0; idx < 4; idx++ {
		vs.sss[idx] = visitedSectionState{
			p:    mss.sectionStates[idx].pos,
			mask: mss.keys & mss.sectionStates[idx].section.allKeysMask,
		}
	}
	return vs
}

type visitedSectionState struct {
	p    util.Point
	mask keyState
}

type visitedState struct {
	sss [4]visitedSectionState
}

func (b *Board) initialMultiSectionState() multiSectionState {
	var secStates [4]sectionState
	for idx := 0; idx < b.sectionCount; idx++ {
		secStates[idx] = sectionState{
			pos:     b.sections[idx].start,
			section: &b.sections[idx],
		}
	}
	s := multiSectionState{
		sectionStates: secStates,
		keys:          0,
	}
	return s
}

func (b *Board) isDone(s multiSectionState) bool {
	done := true
	for secIdx := 0; secIdx < b.sectionCount; secIdx++ {
		if !b.sections[secIdx].hasAllKeys(s.keys) {
			done = false
		}
	}
	return done
}

func (b *Board) multiSectionBfs() int {
	q := util.NewQueue()

	dist := make(map[visitedState]int)

	s := b.initialMultiSectionState()

	dist[s.getState()] = 0
	q.Push(s)

	steps := 0

	for q.Size() > 0 {
		steps += 1

		s = q.Pop().(multiSectionState)
		ss := s.getState()
		d := dist[ss]

		//if steps%10000 == 0 {
		//	fmt.Println(d)
		//}

		if b.isDone(s) {
			fmt.Printf("Took %d steps\n", steps)
			return d
		}

		for secIdx := 0; secIdx < b.sectionCount; secIdx++ {
			if s.sectionStates[secIdx].section.hasAllKeys(s.keys) {
				continue
			}
			for idx := 0; idx < 4; idx++ {
				nextPos, nextState := getNextMultiSectionState(s, secIdx, idx)

				if !b.isInRange(nextPos) {
					continue
				}

				chr := b.position(nextPos)

				if isDoor(chr) && !nextState.keys.hasKey(toLower(chr)) {
					continue
				}
				if isKey(chr) {
					nextState.keys.setKey(chr)
				}

				_, visited := dist[nextState.getState()]

				if !visited {
					dist[nextState.getState()] = d + 1
					q.Push(nextState)
				}
			}
		}
	}

	panic("no solution found")
}

func (b *Board) dijkstra(compactedGraph map[util.Point]map[util.Point]CompactedDistance) int {
	pq := util.NewArrayPriorityQueue()
	dist := make(map[visitedState]int)

	s := b.initialMultiSectionState()
	dist[s.getState()] = 0
	pq.Push(s, 0)

	for _, ss := range s.sectionStates {
		b.detectCycle(ss.pos)
	}

	for !pq.Empty() {
		s := pq.Pop().(multiSectionState)
		ss := s.getState()
		d := dist[ss]

		if b.isDone(s) {
			return d
		}

		for secIdx := 0; secIdx < b.sectionCount; secIdx++ {
			if s.sectionStates[secIdx].section.hasAllKeys(s.keys) {
				continue
			}

			neighbors := compactedGraph[s.sectionStates[secIdx].pos]
			//neighbors := b.gateAwareBfs(s.sectionStates[secIdx].pos)

			for nextPos, cost := range neighbors {
				if !s.keys.hasRequiredKeys(cost.req) {
					continue
				}
				c := b.position(nextPos)
				if isKey(c) && s.keys.hasKey(c) {
					continue
				}
				if isDoor(c) && !s.keys.hasKey(toLower(c)) {
					continue
				}

				var nextSectionStates [4]sectionState
				for i := 0; i < 4; i++ {
					if i == secIdx {
						nextSectionStates[i] = sectionState{
							pos:     nextPos,
							section: s.sectionStates[secIdx].section,
						}
					} else {
						nextSectionStates[i] = s.sectionStates[i]
					}
				}
				nextState := multiSectionState{
					sectionStates: nextSectionStates,
					keys:          s.keys,
				}

				if isKey(c) {
					nextState.keys.setKey(c)
				}

				dState := nextState.getState()
				currDist, visited := dist[dState]

				if !visited || currDist > d+cost.dist {
					dist[dState] = d + cost.dist
					pq.Push(nextState, d+cost.dist)
				}
			}
		}
	}

	panic("no solution found")
}

type CompactedDistance struct {
	dist int
	req  keyState
}

func (b *Board) compactedDistances() map[util.Point]map[util.Point]CompactedDistance {
	res := make(map[util.Point]map[util.Point]CompactedDistance)

	for i, r := range b.board {
		for j, c := range r {
			if c != startSymbol && !isPointOfInterest(c) {
				continue
			}
			p := util.Point{X: i, Y: j}
			dist := b.bfs(p)
			res[p] = dist
		}
	}

	return res
}

func debugCompactedGraph(b *Board, g map[util.Point]map[util.Point]CompactedDistance) {
	for k, v := range g {
		fmt.Printf("For %c (%v) => {", b.position(k), k)

		for n, d := range v {
			fmt.Printf(" [%c (%v) -> %d] ", b.position(n), n, d)
		}

		fmt.Println("}")
	}
}

func PartOne() {
	lines := util.ReadLines("ch18/input.txt")

	b := newBoard(lines)
	b.print()

	fmt.Println(b.multiSectionBfs())
}

func PartTwo() {
	lines := util.ReadLines("ch18/input.txt")

	b := newBoard(lines)
	b.print()

	compactedDistances := b.compactedDistances()
	//debugCompactedGraph(b, compactedDistances)

	//fmt.Println(b.multiSectionBfs())
	fmt.Println(b.dijkstra(compactedDistances))
}
