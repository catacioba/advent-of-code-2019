package ch07

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"math"
	"strings"
	"sync"
)

func computePartOne(program []int, phaseSettings []int) int {
	inputA := make(chan int, 100)
	ab := make(chan int, 100)
	bc := make(chan int, 100)
	cd := make(chan int, 100)
	de := make(chan int, 100)
	outputE := make(chan int, 100)

	programA := intcode.NewIntCodeProgramWithChannels(program, inputA, ab)
	programB := intcode.NewIntCodeProgramWithChannels(program, ab, bc)
	programC := intcode.NewIntCodeProgramWithChannels(program, bc, cd)
	programD := intcode.NewIntCodeProgramWithChannels(program, cd, de)
	programE := intcode.NewIntCodeProgramWithChannels(program, de, outputE)

	inputA <- phaseSettings[0]
	inputA <- 0
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

	return <-outputE
}

func computePartTwo(program []int, phaseSettings []int) int {
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

	return <-ea
}

func PartOne() {
	line := util.ReadLines("ch07/input.txt")[0]
	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	maxValue := math.MinInt32

	for _, permutation := range util.GetPermutations([]int{0, 1, 2, 3, 4}) {
		value := computePartOne(numbers, permutation)

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println(maxValue)
}

func PartTwo() {
	line := util.ReadLines("ch07/input.txt")[0]
	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	maxValue := math.MinInt32

	for _, permutation := range util.GetPermutations([]int{5, 6, 7, 8, 9}) {

		value := computePartTwo(numbers, permutation)

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println(maxValue)
}
