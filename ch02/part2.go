package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fin, _ := os.Open("ch02/opcode.txt")

	scanner := bufio.NewScanner(fin)

	scanner.Scan()
	line := scanner.Text()

	numbersStr := strings.Split(line, ",")
	numbersInitial := make([]int, len(numbersStr))

	for idx, numStr := range numbersStr {
		numbersInitial[idx], _ = strconv.Atoi(numStr)
	}
	numbersStr = nil
	fmt.Println(numbersInitial)

	ok := true
	for noun := 0; noun < 100 && ok; noun++ {
		for verb := 0; verb < 100 && ok; verb++ {

			numbers := make([]int, len(numbersInitial))
			copy(numbers, numbersInitial)

			numbers[1] = noun
			numbers[2] = verb

			idx := 0
			for {
				op := numbers[idx]

				if op == 99 {
					break
				}

				idx1 := numbers[idx+1]
				idx2 := numbers[idx+2]
				idx3 := numbers[idx+3]

				if op == 1 {
					numbers[idx3] = numbers[idx1] + numbers[idx2]
				} else if op == 2 {
					numbers[idx3] = numbers[idx1] * numbers[idx2]
				} else {
					fmt.Println("error")
				}

				idx += 4
			}

			if numbers[0] == 19690720 {
				fmt.Printf("Found noun=%d verb=%d", noun, verb)
				ok = false
			}
		}
	}
}
