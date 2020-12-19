package ch12

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Moon struct {
	position Vector3
	velocity Vector3
}

type Vector3 struct {
	x, y, z int
}

func readInput(lines []string) []Moon {
	res := make([]Moon, len(lines))

	for idx := range lines {
		line := lines[idx]

		substr := line[1 : len(line)-1]

		tokens := strings.Split(substr, ",")

		x, _ := strconv.Atoi(strings.TrimSpace(tokens[0])[2:])
		y, _ := strconv.Atoi(strings.TrimSpace(tokens[1])[2:])
		z, _ := strconv.Atoi(strings.TrimSpace(tokens[2])[2:])

		res[idx] = Moon{
			position: Vector3{
				x: x,
				y: y,
				z: z,
			},
			velocity: Vector3{
				x: 0,
				y: 0,
				z: 0,
			},
		}
	}

	return res
}

func printMoons(step int, moons []Moon) {
	fmt.Printf("Step %d\n", step)
	for _, v := range moons {
		fmt.Printf("Coordinates: %v	Velocity: %v	Pot: %d		Kin: %d		Total: %d\n",
			v.position, v.velocity, getKineticEnergy(v), getPotentialEnergy(v), getTotalEnergy(v))
	}
	fmt.Printf("Total System Energy: %d", getSystemTotalEnergy(moons))
	fmt.Println()
}

func adjustGravity(moon1, moon2 *Moon) {
	if moon1.position.x > moon2.position.x {
		moon1.velocity.x -= 1
		moon2.velocity.x += 1
	} else if moon1.position.x < moon2.position.x {
		moon2.velocity.x -= 1
		moon1.velocity.x += 1
	}

	if moon1.position.y > moon2.position.y {
		moon1.velocity.y -= 1
		moon2.velocity.y += 1
	} else if moon1.position.y < moon2.position.y {
		moon2.velocity.y -= 1
		moon1.velocity.y += 1
	}

	if moon1.position.z > moon2.position.z {
		moon1.velocity.z -= 1
		moon2.velocity.z += 1
	} else if moon1.position.z < moon2.position.z {
		moon2.velocity.z -= 1
		moon1.velocity.z += 1
	}
}

func moveMoon(moon *Moon) {
	moon.position.x += moon.velocity.x
	moon.position.y += moon.velocity.y
	moon.position.z += moon.velocity.z
}

func getPotentialEnergy(moon Moon) int {
	return util.MyAbs(moon.position.x) + util.MyAbs(moon.position.y) + util.MyAbs(moon.position.z)
}

func getKineticEnergy(moon Moon) int {
	return util.MyAbs(moon.velocity.x) + util.MyAbs(moon.velocity.y) + util.MyAbs(moon.velocity.z)
}

func getTotalEnergy(moon Moon) int {
	return getPotentialEnergy(moon) * getKineticEnergy(moon)
}

func getSystemTotalEnergy(moons []Moon) int {
	cnt := 0
	for _, moon := range moons {
		cnt += getTotalEnergy(moon)
	}
	return cnt
}

type System struct {
}

func has0XVelocity(moons []Moon) bool {
	for i := 0; i < 4; i++ {
		if moons[i].velocity.x != 0 {
			return false
		}
	}
	return true
}

func has0YVelocity(moons []Moon) bool {
	for i := 0; i < 4; i++ {
		if moons[i].velocity.y != 0 {
			return false
		}
	}
	return true
}

func has0ZVelocity(moons []Moon) bool {
	for i := 0; i < 4; i++ {
		if moons[i].velocity.z != 0 {
			return false
		}
	}
	return true
}

func main() {
	lines := util.ReadLines("ch12/input.txt")

	moons := readInput(lines)

	//printMoons(0, moons)

	xOk := false
	yOk := false
	zOk := false
	var xPeriod, yPeriod, zPeriod int

	//for t := 1; t <= 3000; t++ {
	//cnt := 3
	//for t := 1; cnt > 0; t++ {
	for t := 1; !xOk || !yOk || !zOk; t++ {
		for moon1 := 0; moon1 < 3; moon1++ {
			for moon2 := moon1 + 1; moon2 < 4; moon2++ {
				adjustGravity(&moons[moon2], &moons[moon1])
			}
		}

		for i := 0; i < 4; i++ {
			moveMoon(&moons[i])
		}

		//if getSystemTotalEnergy(moons) == 0 {
		//	break
		//}

		if !xOk && has0XVelocity(moons) {
			xOk = true
			xPeriod = t
			fmt.Printf("x 0 at time %d\n", t)
		}
		if !yOk && has0YVelocity(moons) {
			yOk = true
			yPeriod = t
			fmt.Printf("y 0 at time %d\n", t)
		}
		if !zOk && has0ZVelocity(moons) {
			zOk = true
			zPeriod = t
			fmt.Printf("z 0 at time %d\n", t)
		}
	}

	lcm := util.Lcm3(xPeriod, yPeriod, zPeriod)

	fmt.Println(lcm * 2)

	fmt.Println("done")
}
