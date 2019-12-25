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
	numbers := make([]int, len(numbersStr))

	for idx, numStr := range numbersStr {
		numbers[idx], _ = strconv.Atoi(numStr)
	}

	fmt.Println(numbers)

	numbers[1] = 12
	numbers[2] = 2

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

	fmt.Printf("the result is %d", numbers[0])
}
