package ch09

import (
	"aoc/intcode"
	"aoc/util"
	"fmt"
	"strings"
)

func PartOne() {
	line := util.ReadLines("ch09/input.txt")[0]

	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	numbersBigger := make([]int, 100000)
	copy(numbersBigger, numbers)

	program := intcode.NewIntCodeProgram(numbersBigger)

	program.Input <- 1

	program.Run()

	code := <-program.Output
	fmt.Println(code)
}

func PartTwo() {
	line := util.ReadLines("ch09/input.txt")[0]

	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	numbersBigger := make([]int, 100000)
	copy(numbersBigger, numbers)

	program := intcode.NewIntCodeProgram(numbersBigger)

	program.Input <- 2

	program.Run()

	code := <-program.Output
	fmt.Println(code)
}
