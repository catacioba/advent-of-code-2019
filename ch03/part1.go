package main

import (
	. "adventofcode/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getPoints(wire string) map[Point]bool {
	points := make(map[Point]bool)

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
				points[Point{x, y}] = true
			}
		case "L":
			for idx := 0; idx < cnt; idx++ {
				x--
				points[Point{x, y}] = true
			}
		case "U":
			for idx := 0; idx < cnt; idx++ {
				y++
				points[Point{x, y}] = true
			}
		case "D":
			for idx := 0; idx < cnt; idx++ {
				y--
				points[Point{x, y}] = true
			}
		}
	}

	return points
}

func main() {

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
