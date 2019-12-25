package main

import (
	"adventofcode/intcode"
	"adventofcode/util"
	"fmt"
	"math"
	"strings"
	"sync"
)

func compute(program intcode.Program, phaseSettings []int) int {

	value := 0

	for i := 0; i < len(phaseSettings); i++ {
		programCopy := program

		programCopy.Input <- phaseSettings[i]
		programCopy.Input <- value

		programCopy.Run()

		value = <-programCopy.Output
	}

	return value
}

func compute2(program []int, phaseSettings []int) int {
	value := 0

	//channels := make([]chan int, len(phaseSettings))
	//
	//for i := range phaseSettings {
	//	channels[i] = make(chan int, 10)
	//}

	ea := make(chan int, 100)
	ab := make(chan int, 100)
	bc := make(chan int, 100)
	cd := make(chan int, 100)
	de := make(chan int, 100)

	programA := intcode.NewIntCodeProgramWithChannels(program, ea, ab)
	programB := intcode.NewIntCodeProgramWithChannels(program, ab, bc)
	programC := intcode.NewIntCodeProgramWithChannels(program, bc, cd)
	programD := intcode.NewIntCodeProgramWithChannels(program, cd, de)
	programE := intcode.NewIntCodeProgramWithChannels(program, de, ea)
	//programA.Input = ea
	//programA.Output = ab
	//
	//programB := program
	//programB.Input = ab
	//programB.Output = bc
	//
	//programC := program
	//programC.Input = bc
	//programC.Output = cd
	//
	//programD := program
	//programD.Input = cd
	//programD.Output = de
	//
	//programE := program
	//programE.Input = de
	//programE.Output = ea

	ea <- phaseSettings[0]
	ea <- 0
	ab <- phaseSettings[1]
	bc <- phaseSettings[2]
	cd <- phaseSettings[3]
	de <- phaseSettings[4]

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		programA.Run()
	}()
	go func() {
		defer wg.Done()
		programB.Run()
	}()
	go func() {
		defer wg.Done()
		programC.Run()
	}()
	go func() {
		defer wg.Done()
		programD.Run()
	}()
	go func() {
		defer wg.Done()
		programE.Run()
	}()

	wg.Wait()

	value = <-ea

	return value
}

func main() {

	line := util.ReadLines("ch07/input.txt")[0]
	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	//program := util.NewIntCodeProgram(numbers)
	//
	//program.Input <- 3
	//program.Input <- 0
	//
	//program.Run()
	//
	//fmt.Println(<-program.Output)

	//fmt.Println(compute(numbers, []int{9, 7, 8, 5, 6}))

	//maxValue := math.MinInt32
	//program := util.NewIntCodeProgram(numbers)
	//
	//for _, permutation := range util.GetPermutations([]int{5, 6, 7, 8, 9}) {
	//
	//	value := compute(program, permutation)
	//
	//	if value > maxValue {
	//		maxValue = value
	//	}
	//}
	//
	//fmt.Println(maxValue)

	maxValue := math.MinInt32
	//program := util.NewIntCodeProgram(numbers)

	for _, permutation := range util.GetPermutations([]int{5, 6, 7, 8, 9}) {

		value := compute2(numbers, permutation)

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println(maxValue)
}
