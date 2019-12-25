package util

type Point struct {
	X, Y int
}

type Node struct {
	value    string
	children []*Node
}
