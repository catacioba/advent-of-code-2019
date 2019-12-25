package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"sort"
)

type LinePoint struct {
	quadrant int
	slope    float64
}

type LinePointList struct {
	linePoint LinePoint
	points    []util.Point
}

func isBefore(x1, y1, x2, y2 int) bool {
	if x1 < x2 {
		return true
	} else if x1 == x2 {
		return y1 < y2
	}
	return false
}

func getQuadrant(xOrigin, yOrigin, x, y int) int {
	if xOrigin < x {
		if yOrigin < y {
			return 2
		} else {
			return 3
		}
	} else {
		if yOrigin <= y {
			return 1
		} else {
			return 4
		}
	}
}

func findReachableMap(asteroidMap []string, x, y int) map[LinePoint][]util.Point {
	h := len(asteroidMap)
	w := len(asteroidMap[0])

	slopes := make(map[LinePoint][]util.Point)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == x && j == y {
				continue
			}

			if asteroidMap[i][j] == '#' {
				//fmt.Printf("%d %d\n", i, j)
				slope := float64(i-x) / float64(j-y)

				lp := LinePoint{
					//before: isBefore(i, j, x, y),
					quadrant: getQuadrant(x, y, i, j),
					slope:    slope,
				}

				points, ok := slopes[lp]

				if ok {
					points = append(points, util.Point{X: i, Y: j})

					slopes[lp] = points
				} else {
					slopes[lp] = []util.Point{{i, j}}
				}
			}
		}
	}

	////fmt.Println(slopes)
	//for k, v := range slopes {
	//	fmt.Printf("%v %v\n", k, v)
	//}

	return slopes
}

func findReachable(asteroidMap []string, x, y int) int {

	slopes := findReachableMap(asteroidMap, x, y)

	return len(slopes)
}

func main() {
	lines := util.ReadLines("ch10/input.txt")

	fmt.Println(lines)

	fmt.Println()

	maxReachable := math.MinInt64
	var maxPointX, MaxPointY int

	for x := range lines {
		for y := range lines[x] {
			if lines[x][y] == '#' {
				reachable := findReachable(lines, x, y)

				if reachable > maxReachable {
					maxReachable = reachable

					maxPointX = x
					MaxPointY = y
				}

				fmt.Print(reachable)
			} else {
				fmt.Printf("%c", lines[x][y])
			}
		}
		fmt.Println()
	}

	fmt.Println(findReachable(lines, 2, 2))
	fmt.Println()
	fmt.Printf("%d %d => %d\n", maxPointX, MaxPointY, maxReachable)

	reachableMap := findReachableMap(lines, maxPointX, MaxPointY)
	//fmt.Println(reachableMap)

	maxPoint := util.Point{maxPointX, MaxPointY}

	var points []LinePointList
	for k, v := range reachableMap {

		sort.Slice(v, func(i, j int) bool {
			return util.GetDistance(maxPoint, v[i]) < util.GetDistance(maxPoint, v[j])
		})

		points = append(points, LinePointList{
			linePoint: k,
			points:    v,
		})
	}

	sort.Slice(points, func(i, j int) bool {
		pointI := points[i].linePoint
		pointJ := points[j].linePoint

		if pointI.quadrant == pointJ.quadrant {
			return pointI.slope < pointJ.slope
		} else {
			return pointI.quadrant < pointJ.quadrant
		}
	})

	for _, p := range points {
		fmt.Println(p)
	}

	steps := 1
	idx := 0

	for steps <= 200 {
		p := points[idx]

		fmt.Printf("%d => {%d %d}\n", steps, p.points[0].Y, p.points[0].X)

		points[idx].points = points[idx].points[1:]

		idx++
		steps++
	}
}
