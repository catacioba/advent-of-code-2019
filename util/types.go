package util

type Point struct {
	X, Y int
}

type Node struct {
	value    string
	children []*Node
}

func (p *Point) Add(other Point) Point {
	return Point {
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p *Point) Distance(other Point) int {
	return MyAbs(p.X-other.X) + MyAbs(p.Y-other.Y)
}
