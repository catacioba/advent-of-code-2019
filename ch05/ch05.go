package ch05

import (
	"aoc/intcode"
	"fmt"
)

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch05/input.txt")

	program.Input <- 1

	go program.Run()

	for c := range program.Output {
		fmt.Println(c)
	}
}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch05/input.txt")

	program.Input <- 5

	go program.Run()

	for c := range program.Output {
		fmt.Println(c)
	}
}
