package ch18

import (
	"aoc/util"
	"fmt"
)

var directionsX = []int{1, 0, -1, 0}
var directionsY = []int{0, 1, 0, -1}

type Board struct {
	board []string
	start util.Point

	keys  []byte
	doors []byte
	path  []byte
}

func (b *Board) initBoard() {
	for x := range b.board {
		for y := range b.board[x] {
			//if b.board[x][y] == '@' {
			//	b.start = util.Point{
			//		X: x,
			//		Y: y,
			//	}
			//	return
			//}
			switch b.board[x][y] {
			case '@':
				b.start = util.Point{X: x, Y: y}
			default:
				chr := b.board[x][y]
				if isDoor(chr) {
					b.doors = append(b.doors, chr)
				} else if isKey(chr) {
					b.keys = append(b.keys, chr)
				}
			}
		}
	}
}

func newBoard(lines []string) *Board {
	b := Board{board: lines}
	b.initBoard()
	return &b
}

func (b *Board) isInRange(pos util.Point) bool {
	if pos.X < 0 || pos.Y < 0 {
		return false
	}
	if pos.X > len(b.board) || pos.Y > len(b.board[0]) {
		return false
	}
	if b.board[pos.X][pos.Y] == '#' {
		return false
	}
	return true
}

func (b *Board) dfs(startingPoint util.Point) {
	visited := map[util.Point]bool{}

	b.dfsUtil(startingPoint, visited)
}

func (b *Board) dfsUtil(position util.Point, visited map[util.Point]bool) {
	if !b.isInRange(position) {
		return
	}

	chr := b.board[position.X][position.Y]

	if isDoor(chr) {
		b.doors = append(b.doors, chr)
	} else if isKey(chr) {
		b.keys = append(b.keys, chr)
	}
	if isDoor(chr) || isKey(chr) {
		b.path = append(b.path, chr)
	}

	for t := 0; t < 4; t++ {
		p := getNextPosition(position, t)

		if !visited[p] {
			visited[p] = true

			b.dfsUtil(p, visited)
		}
	}
}

func (b *Board) detectCycle() {
	visited := make(map[util.Point]struct{})

	hasCycle := b.detectCycleAux(b.start, visited, b.start)

	fmt.Printf("has cycle: %v\n", hasCycle)
}

func (b *Board) detectCycleAux(p util.Point, visited map[util.Point]struct{}, parent util.Point) bool {
	if !b.isInRange(p) {
		return false
	}

	visited[p] = struct{}{}

	for t := 0; t < 4; t++ {
		n := getNextPosition(p, t)

		if parent == n {
			continue
		}

		_, ok := visited[n]
		if !ok {
			//hasCycle = hasCycle || b.detectCycleAux(n, visited, p)
			if b.detectCycleAux(n, visited, p) {
				return true
			}
		} else {
			return true
		}
	}

	return false
}

//func topologicalSort(board []string, position util.Point, stack []byte) {
//	if !isInRange(board, position) {
//		return
//	}
//}

//type state struct {
//	pos   util.Point
//	keys  []byte
//	gates []byte
//}

// func newState(p util.Point) *state {
// 	return &state{
// 		pos:   p,
// 		keys:  nil,
// 		gates: nil,
// 	}
// }

//type searchState struct {
//	pos util.Point
//
//	collectedKeys []byte
//	keysToCollect
//}

type state struct {
	p    util.Point
	keys []byte
}

func copyState(s state) state {
	keysCopy := make([]byte, len(s.keys))
	copy(keysCopy, s.keys)
	return state{
		p:    s.p,
		keys: keysCopy,
	}
}

func topSort(adj map[byte][]byte) []byte {
	result := make([]byte, 0)

	//in := make([]int, n)
	in := make(map[byte]int)
	q := util.NewQueue()

	for k, v := range adj {
		//in[toIntFromLower(k)] = len(v)
		in[k] = len(v)
		if len(v) == 0 {
			q.Push(k)
		}
	}

	//fmt.Printf("==> %v\n", adj)
	////fmt.Printf("==> %v\n", in)
	//fmt.Println("-->")
	//for k, v := range in {
	//	fmt.Printf("%c => %d\n", k, v)
	//}
	next := make(map[byte][]byte)
	for k, _ := range adj {
		next[k] = make([]byte, 0)
	}
	for k, v := range adj {
		for _, n := range v {
			next[n] = append(next[n], k)
		}
	}

	//printMap(adj)
	//fmt.Println("===")
	//printMap(next)

	for q.Size() > 0 {
		el := q.Pop().(byte)
		result = append(result, el)
		for _, x := range next[el] {
			//in[toIntFromLower(x)]--
			in[x] = in[x] - 1
			if in[x] == 0 {
				//q.Push(toIntFromLower(x))
				q.Push(x)
			}
		}
	}

	return result
}

func printMap(m map[byte][]byte) {
	for k, v := range m {
		fmt.Printf("%c => ", k)
		printKeys(v)
	}
}

func (b *Board) bfs() {
	q := util.NewQueue()

	visited := make(map[util.Point]bool)
	dist := make(map[util.Point]int)
	//prev := make(map[util.Point]util.Point)

	requiredKeys := make(map[byte][]byte)
	for _, d := range b.doors {
		requiredKeys[d] = make([]byte, 0)
	}

	s := state{
		p:    b.start,
		keys: make([]byte, 0),
	}
	q.Push(s)
	dist[s.p] = 0
	visited[s.p] = true

	adj := make(map[byte][]byte)

	for q.Size() > 0 {
		//position := q.Pop().(util.Point)
		s = q.Pop().(state)
		//fmt.Printf("=> %v\n", s.p)

		chr := b.board[s.p.X][s.p.Y]

		//if isDoor(chr) || isKey(chr) {
		//	fmt.Printf("# %c is at distance %d\n", chr, dist[s.p])
		//}
		if isDoor(chr) {
			s.keys = append(s.keys, toLower(chr))
		}
		if isKey(chr) {
			fmt.Printf("# %c is at distance %d and requires keys ", chr, dist[s.p])
			printKeys(s.keys)
			adj[chr] = s.keys
			//for _, bb := range s.keys {
			//	fmt.Printf("%c ", bb)
			//}
			//fmt.Println("]")
		}

		//possibleDirections := 0
		for t := 0; t < 4; t++ {
			p := getNextPosition(s.p, t)

			if b.isInRange(p) && !visited[p] {
				//possibleDirections++

				visited[p] = true
				dist[p] = dist[s.p] + 1
				//prev[p] = s.p

				//q.Push(p)
				ns := copyState(s)
				ns.p = p
				q.Push(ns)
			}
		}
		//if possibleDirections > 1 {
		//	fmt.Printf("split occurred at %v\n", s.p)
		//}
	}

	ts := topSort(adj)
	fmt.Print("top sort: ")
	printKeys(ts)
}

func printKeys(arr []byte) {
	if len(arr) == 0 {
		fmt.Println("[]")
		return
	}
	fmt.Print("[")
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%c ", arr[i])
	}
	fmt.Printf("%c]\n", arr[len(arr)-1])
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

func PartOne() {
	lines := util.ReadLines("ch18/input.txt")

	// board := Board{board: lines, start: start}
	// board.bfs()

	//b := Board{board: lines, start: start}
	b := newBoard(lines)
	b.print()

	//b.dfs(b.start)
	b.bfs()

	// // for k := range visited {
	// // 	fmt.Println(k)
	// // }

	//fmt.Printf("path: %v\n", b.path)
	//fmt.Printf("keys: %v\n", b.keys)
	//fmt.Printf("doors: %v\n", b.doors)
	//
	//b.detectCycle()

	// s := []int{}

	// topologicalSort(board, start, s)

	// fmt.Println(s)
}

func PartTwo() {}
