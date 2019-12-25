package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getDirection(dir string) util.Point {
	switch dir {
	case "U":
		return util.Point{Y: 1}
	case "D":
		return util.Point{Y: -1}
	case "R":
		return util.Point{X: 1}
	case "L":
		return util.Point{X: -1}
	}
	panic("unknown direction!")
}

type Path struct {
	direction util.Point
	steps     int
}

func parseWire(wire string) []Path {
	paths := strings.Split(wire, ",")

	var res []Path

	for _, path := range paths {
		dir := path[:1]
		steps, _ := strconv.Atoi(path[1:])

		res = append(res, Path{getDirection(dir), steps})
	}

	return res
}

func move(location *util.Point, direction util.Point) {
	location.X += direction.X
	location.Y += direction.Y
}

func getPointsWithSteps(wire string) map[util.Point]int {
	x := 0
	y := 0
	steps := 0

	points := make(map[util.Point]int)

	for _, path := range strings.Split(wire, ",") {
		dir := path[:1]
		cnt, _ := strconv.Atoi(path[1:])
		path := getDirection(dir)

		for idx := 0; idx < cnt; idx++ {
			steps++

			x += path.X
			y += path.Y

			pos := util.Point{X: x, Y: y}

			_, ok := points[pos]

			if ok == false {
				points[pos] = steps
			}
		}

	}

	return points
}

func getSteps(wire string, otherWire map[util.Point]int) int {
	x := 0
	y := 0
	steps := 0

	for _, path := range strings.Split(wire, ",") {
		dir := path[:1]
		cnt, _ := strconv.Atoi(path[1:])
		path := getDirection(dir)

		for idx := 0; idx < cnt; idx++ {
			steps++

			x += path.X
			y += path.Y

			pos := util.Point{X: x, Y: y}

			otherSteps, ok := otherWire[pos]

			if ok == true {
				return steps + otherSteps
			}
		}

	}

	panic("this should not happen")
}

func main() {

	lines := util.ReadLines("ch03/input.txt")

	wire1 := lines[0]
	wire2 := lines[1]

	//points1 := parseWire(wire1)
	//points2 := parseWire(wire2)
	//
	//wire1Position := util.Point{}
	//wire2Position := util.Point{}
	//
	//steps := 0
	//idx1 := 0
	//path1 := Path{}
	//idx2 := 0
	//path2 := Path{}
	//
	//for wire1Position != wire2Position {
	//	if path1.steps == 0 {
	//		path1 = points1[idx1]
	//		idx1++
	//	}
	//	if path2.steps == 0 {
	//		path2 = points2[idx2]
	//		idx2++
	//	}
	//
	//	move(&wire1Position, path1.direction)
	//	path1.steps--
	//	move(&wire2Position, path2.direction)
	//	path2.steps--
	//}

	points1 := getPointsWithSteps(wire1)
	points2 := getPointsWithSteps(wire2)

	minSteps := math.MaxInt32

	for point, steps1 := range points1 {
		steps2, found := points2[point]

		if found {
			if steps1+steps2 < minSteps {
				minSteps = steps1 + steps2
			}
		}
	}

	fmt.Println(minSteps)
	//fmt.Println(getSteps(wire2, points1))
}
