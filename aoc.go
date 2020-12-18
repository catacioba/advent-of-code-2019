package main

import (
	"aoc/ch01"
	"aoc/ch02"
	"aoc/ch03"
	"aoc/ch04"
	"aoc/ch13"
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
	13: {
		1: ch13.PartOne,
		2: ch13.PartTwo,
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
