package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getOp(op int) int {
	return op % 100
}

func getParamMode(op, paramPosition int) int {
	op /= 100

	for paramPosition != 0 {
		op /= 10
		paramPosition--
	}

	return paramPosition % 10
}

func getParamFromMask(op, mask int) int {
	return op / mask % 10
}

func getMask(pos int) int {
	res := 100

	for pos > 1 {
		res *= 10
		pos--
	}

	return res
}

func getParam(numbers []int, opCode int, idx int, position int) int {
	mask := getMask(position)

	if getParamFromMask(opCode, mask) == 1 {
		return numbers[idx+position]
	} else {
		return numbers[numbers[idx+position]]
	}
}

func main() {
	fin, _ := os.Open("ch05/input.txt")

	scanner := bufio.NewScanner(fin)

	scanner.Scan()
	line := scanner.Text()

	numbersStr := strings.Split(line, ",")
	numbers := make([]int, len(numbersStr))

	for idx, numStr := range numbersStr {
		numbers[idx], _ = strconv.Atoi(numStr)
	}

	numbersStr = nil
	line = ""
	fmt.Println(numbers)

	idx := 0
	shouldRun := true
	for shouldRun {
		opCode := numbers[idx]

		op := getOp(opCode)

		switch op {
		case 99:
			shouldRun = false
			break

		case 1:
			num1 := getParam(numbers, opCode, idx, 1)
			num2 := getParam(numbers, opCode, idx, 2)

			param3 := numbers[idx+3]

			numbers[param3] = num1 + num2
			idx += 4
		case 2:
			num1 := getParam(numbers, opCode, idx, 1)
			num2 := getParam(numbers, opCode, idx, 2)

			param3 := numbers[idx+3]

			numbers[param3] = num1 * num2
			idx += 4
		case 3:
			param1 := numbers[idx+1]

			fmt.Println("Input operation")

			numbers[param1] = 5
			idx += 2
		case 4:
			//param1 := numbers[idx+1]
			param1 := getParam(numbers, opCode, idx, 1)

			fmt.Println(param1)
			idx += 2
		case 5:
			param1 := getParam(numbers, opCode, idx, 1)
			param2 := getParam(numbers, opCode, idx, 2)

			if param1 != 0 {
				idx = param2
			} else {
				idx += 3
			}
		case 6:
			param1 := getParam(numbers, opCode, idx, 1)
			param2 := getParam(numbers, opCode, idx, 2)

			if param1 == 0 {
				idx = param2
			} else {
				idx += 3
			}
		case 7:
			param1 := getParam(numbers, opCode, idx, 1)
			param2 := getParam(numbers, opCode, idx, 2)

			param3 := numbers[idx+3]

			if param1 < param2 {
				numbers[param3] = 1
			} else {
				numbers[param3] = 0
			}
			idx += 4
		case 8:
			param1 := getParam(numbers, opCode, idx, 1)
			param2 := getParam(numbers, opCode, idx, 2)

			param3 := numbers[idx+3]

			if param1 == param2 {
				numbers[param3] = 1
			} else {
				numbers[param3] = 0
			}
			idx += 4
		default:
			fmt.Println("error")
		}
	}
}
