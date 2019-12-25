package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

//func getOrCreate(planets map[string]util.Node, planetName string) {
//	planet, ok := planets[planetName]
//
//	if !ok {
//		planet = util.Node{planetName
//		}
//	}
//	return planet
//}
//
//func createOrbitTree(lines []string) map[string]util.Node {
//	res := make(map[string]util.Node)
//
//	for _, line := range lines {
//		fmt.Println(line)
//
//		objects := strings.Split(line, ")")
//
//		left := objects[0]
//		right := objects[1]
//
//
//	}
//
//	return res
//}

type OrbitInfo struct {
	directOrbit string
	value       string
	orbitCount  int
	visited     bool
}

func goUp(orbitInfoMap map[string]*OrbitInfo, planet string) int {
	planetInfo := orbitInfoMap[planet]

	if !planetInfo.visited {
		cnt := goUp(orbitInfoMap, planetInfo.directOrbit)

		planetInfo.orbitCount = cnt + 1
		planetInfo.visited = true
	}

	return planetInfo.orbitCount
}

func main() {
	lines := util.ReadLines("ch06/input.txt")

	orbitInfoMap := make(map[string]*OrbitInfo)
	orbitInfoMap["COM"] = &OrbitInfo{
		directOrbit: "",
		value:       "COM",
		orbitCount:  0,
		visited:     true,
	}

	for _, line := range lines {
		objects := strings.Split(line, ")")

		left := objects[0]
		right := objects[1]

		orbitInfoMap[right] = &OrbitInfo{left, right, 0, false}
	}

	totalOrbits := 0

	for planet, _ := range orbitInfoMap {
		orbits := goUp(orbitInfoMap, planet)

		totalOrbits += orbits
	}

	//fmt.Println(orbitInfoMap)
	fmt.Println(totalOrbits)

	san := orbitInfoMap["SAN"]
	//fmt.Println(san)

	you := orbitInfoMap["YOU"]
	//fmt.Println(you)

	// get common ancestor
	leftAncestors := make(map[string]bool)

	it := you

	for it.directOrbit != "" {
		leftAncestors[it.directOrbit] = true
		it = orbitInfoMap[it.directOrbit]
	}

	it = san
	var commonAncestor *OrbitInfo

	for it.directOrbit != "" {
		_, ok := leftAncestors[it.value]

		if ok {
			//fmt.Println(it)
			commonAncestor = it
			break
		}
		it = orbitInfoMap[it.directOrbit]
	}

	fmt.Println(san)
	fmt.Println(you)
	fmt.Println(commonAncestor)
	fmt.Println(san.orbitCount + you.orbitCount - 2*commonAncestor.orbitCount - 2)
}
