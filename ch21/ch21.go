package ch21

import (
	"aoc/intcode"
	"fmt"
)

func read(ch chan int) {
	cnt := 0
	for x := range ch {
		if x == 10 && cnt > 0 {
			return
		}

		if x == 10 {
			cnt++
		} else {
			cnt = 0
		}

		fmt.Printf("%c", x)
	}
}

func readAll(ch chan int) {
	for x := range ch {
		fmt.Printf("%c", x)
	}
}

func sendInstruction(ch chan int, s string) {
	for _, x := range s {
		y := int(x)
		ch <- y
	}
	ch <- 10 // new line ascii
}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch21/input.txt")

	go program.Run()

	instructions := `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
WALK`

	sendInstruction(program.Input, instructions)

	read(program.Output)

	read(program.Output)

	fmt.Println(<-program.Output)
}

func PartTwo() {
	program := intcode.NewIntCodeProgramFromFile("ch21/input.txt")

	go program.Run()

	instructions := `NOT A J
NOT B T
OR T J
NOT C T
OR T J
AND D J
NOT E T
AND H T
OR E T
AND T J
RUN`

	sendInstruction(program.Input, instructions)

	read(program.Output)

	read(program.Output)

	fmt.Println(<-program.Output)
}
