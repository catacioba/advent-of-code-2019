package main

import (
	"aoc/ch01"
	"aoc/ch02"
	"aoc/ch03"
	"aoc/ch04"
	"aoc/ch05"
	"aoc/ch06"
	"aoc/ch07"
	"aoc/ch08"
	"aoc/ch09"
	"aoc/ch13"
	"aoc/ch15"
	"aoc/ch17"
	"aoc/ch18"
	"aoc/ch19"
	"aoc/ch21"
	"aoc/ch23"
	"aoc/ch24"
	"aoc/ch25"
	"flag"
	"fmt"
	"log"
)

var challengeMap = map[int]map[int]func(){
	1: {
		1: ch01.PartOne,
		2: ch01.PartTwo,
	},
	2: {
		1: ch02.PartOne,
		2: ch02.PartTwo,
	},
	3: {
		1: ch03.PartOne,
		2: ch03.PartTwo,
	},
	4: {
		1: ch04.PartOne,
		2: ch04.PartTwo,
	},
	5: {
		1: ch05.PartOne,
		2: ch05.PartTwo,
	},
	6: {
		1: ch06.Solve,
		2: ch06.Solve,
	},
	7: {
		1: ch07.PartOne,
		2: ch07.PartTwo,
	},
	8: {
		1: ch08.Solve,
		2: ch08.Solve,
	},
	9: {
		1: ch09.PartOne,
		2: ch09.PartTwo,
	},
	13: {
		1: ch13.PartOne,
		2: ch13.PartTwo,
	},
	15: {
		1: ch15.PartOne,
		2: ch15.PartTwo,
	},
	17: {
		1: ch17.PartOne,
		2: ch17.PartTwo,
	},
	18: {
		1: ch18.PartOne,
		2: ch18.PartTwo,
	},
	19: {
		1: ch19.PartOne,
		2: ch19.PartTwo,
	},
	21: {
		1: ch21.PartOne,
		2: ch21.PartTwo,
	},
	23: {
		1: ch23.PartOne,
		2: ch23.PartTwo,
	},
	24: {
		1: ch24.PartOne,
		2: ch24.PartTwo,
	},
	25: {
		1: ch25.PartOne,
		2: ch25.PartTwo,
	},
}

func main() {
	challengeFlag := flag.Int("ch", 0, "the challenge to run")
	partFlag := flag.Int("p", 0, "the part of the challenge to run")

	flag.Parse()

	if *challengeFlag <= 0 || *challengeFlag > 25 {
		log.Fatal("Invalid challenge value! (1-25)")
	}
	if *partFlag != 1 && *partFlag != 2 {
		log.Fatal("Invalid challenge part value! (1 or 2)")
	}

	challenge, ok := challengeMap[*challengeFlag][*partFlag]
	if ok == false {
		log.Fatal("Challenge not found!")
	}

	fmt.Printf("Running challenge %d part %d\n", *challengeFlag, *partFlag)

	challenge()
}
