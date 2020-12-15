package main

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getPoints(wire string) map[util.Point]bool {
	points := make(map[util.Point]bool)

	x := 0
	y := 0

	for _, path := range strings.Split(wire, ",") {
		dir := path[:1]
		cnt, _ := strconv.Atoi(path[1:])

		//fmt.Printf("%s %d \n", dir, cnt)
		switch dir {
		case "R":
			for idx := 0; idx < cnt; idx++ {
				x++
				points[util.Point{X: x, Y: y}] = true
			}
		case "L":
			for idx := 0; idx < cnt; idx++ {
				x--
				points[util.Point{X: x, Y: y}] = true
			}
		case "U":
			for idx := 0; idx < cnt; idx++ {
				y++
				points[util.Point{X: x, Y: y}] = true
			}
		case "D":
			for idx := 0; idx < cnt; idx++ {
				y--
				points[util.Point{X: x, Y: y}] = true
			}
		}
	}

	return points
}

func PartOne() {

	fin, _ := os.Open("ch03/input.txt")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	scanner.Scan()
	wire1 := scanner.Text()
	scanner.Scan()
	wire2 := scanner.Text()

	points1 := getPoints(wire1)
	points2 := getPoints(wire2)

	//closestPoint := Point{0,0}
	closestDist := 10000000.0

	for point, _ := range points1 {
		_, ok := points2[point]

		if ok {
			dist := math.Abs(float64(point.X)) + math.Abs(float64(point.Y))
			if closestDist > dist {
				closestDist = dist
				//closestPoint = point
			}
		}
	}

	fmt.Println(closestDist)
}

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

type path struct {
	direction util.Point
	steps     int
}

func parseWire(wire string) []path {
	paths := strings.Split(wire, ",")

	var res []path

	for _, p := range paths {
		dir := p[:1]
		steps, _ := strconv.Atoi(p[1:])

		res = append(res, path{getDirection(dir), steps})
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

func PartTwo() {

	lines := util.ReadLines("ch03/input.txt")

	wire1 := lines[0]
	wire2 := lines[1]

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
}
