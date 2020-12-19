package ch09

import (
	"aoc/intcode"
	"aoc/util"
	"strings"
)

func main() {
	line := util.ReadLines("ch09/input.txt")[0]

	numbers := util.ConvertStrArrToIntArr(strings.Split(line, ","))

	numbersBigger := make([]int, 100000)
	copy(numbersBigger, numbers)

	program := intcode.NewIntCodeProgram(numbersBigger)

	program.Input <- 2

	program.Run()

}
