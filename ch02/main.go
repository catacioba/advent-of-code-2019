package main

import (
	"adventofcode/intcode"
	"fmt"
)

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch02/opcode.txt")

	program.UpdateMemory(1, 12)
	program.UpdateMemory(2, 2)

	program.Run()

	fmt.Printf("the result is %d", program.Location(0))
}

func PartTwo() {
	ok := true
	for noun := 0; noun < 100 && ok; noun++ {
		for verb := 0; verb < 100 && ok; verb++ {
			program := intcode.NewIntCodeProgramFromFile("ch02/opcode.txt")

			program.UpdateMemory(1, noun)
			program.UpdateMemory(2, verb)

			program.Run()

			if program.Location(0) == 19690720 {
				fmt.Printf("Found noun=%d verb=%d", noun, verb)
				ok = false
			}
		}
	}
}
