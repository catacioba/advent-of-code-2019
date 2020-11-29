package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var pattern = [4]int{0, 1, 0, -1}

func getPattern(size int) [][]int {
	result := make([][]int, size)

	for row := range result {
		rowPattern := make([]int, size)

		current := 0
		cnt := 1

		for idx := range rowPattern {
			if cnt%(row+1) == 0 {
				cnt = 0
				current++
			}

			rowPattern[idx] = pattern[current%len(pattern)]
			cnt++
		}

		result[row] = rowPattern
	}

	return result
}

func dot(a string, rowCnt int) string {
	// if len(a) != len(b) {
	// 	panic("vectors do not have the same length")
	// }

	accum := 0

	current := 0
	cnt := 1

	for idx := range a {
		if cnt%(rowCnt+1) == 0 {
			cnt = 0
			current++
		}

		leftInt := int(a[idx]) - '0'
		// rightInt, _ := strconv.Atoi(b[idx])
		rightInt := pattern[current%len(pattern)]
		cnt++

		// fmt.Println(leftInt, ",", rightInt)

		accum += leftInt * rightInt
	}

	accum = accum % 10

	res := int(math.Abs(float64(accum))) % 10

	// fmt.Printf("%s %v %d\n", a, b, res)

	finalRes := strconv.Itoa(res)

	// fmt.Println(res, "-", finalRes)

	return finalRes
}

func fft(inputSignal string) string {

	// matrix := getPattern(len(inputSignal))

	parts := make([]string, len(inputSignal))

	for idx := range inputSignal {
		parts[idx] = dot(inputSignal, idx)
	}

	return strings.Join(parts, "")
}

func encodeSlow(inputSignal string) string {
	// matrix := getPattern(len(inputSignal))

	for t := 0; t < 100; t++ {
		outputSignal := fft(inputSignal)

		// fmt.Println(outputSignal)
		// fmt.Printf("Iteration %d done\n", t+1)

		inputSignal = outputSignal
	}

	return inputSignal
}

func getRealSignal(inputSignal string) string {
	sb := strings.Builder{}

	for t := 0; t < 10000; t++ {
		sb.WriteString(inputSignal)
	}

	return sb.String()
}

func encodeReal(inputSignal string) string {

	realInput := getRealSignal(inputSignal)

	fmt.Println(len(realInput))

	// matrix := getPattern(len(realInput))

	// fmt.Println("a")

	messageOffset, _ := strconv.Atoi(realInput[:7])

	// for t := 0; t < 100; t++ {
	// 	outputSignal := fft(realInput, matrix)

	// 	fmt.Println("b")

	// 	realInput = outputSignal
	// }
	realInput = encodeSlow(realInput)

	return realInput[messageOffset : messageOffset+8]
	// return realInput
}

// func encodeFast() string {

// }

func conv(a [][]string) [][]int {
	res := make([][]int, len(a))

	for x := 0; x < len(a); x++ {
		row := make([]int, len(a))

		for y := 0; y < len(a); y++ {
			row[y], _ = strconv.Atoi(a[x][y])
		}

		res[x] = row
	}

	return res
}

func mul(a, b [][]int) [][]int {
	n := len(a)

	res := make([][]int, n)

	for x := 0; x < n; x++ {
		res[x] = make([]int, n)

		for y := 0; y < n; y++ {
			accum := 0

			for z := 0; z < n; z++ {
				accum += a[x][z] * b[z][y]
			}

			// accum %= 10

			res[x][y] = accum
		}
	}

	return res
}

func printMatrix(a [][]int) {
	n := len(a)

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Printf("%d\t", a[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

func testAssumption() {
	signal := "12345678123456781234567812345678"

	offset := 20
	offsetSignal := signal[offset:]

	encodeAll := encodeSlow(signal)
	fmt.Printf("%s => %s\n", signal, encodeAll)
	encodePart := encodeSlow(offsetSignal)
	fmt.Printf("%s => %s\n", offsetSignal, encodePart)

	for idx := offset; idx <= len(signal); idx++ {
		if encodeAll[idx] != encodePart[idx-offset] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func charToInt(ch byte) byte {
	return ch - '0'
}

func intToChar(b byte) byte {
	return '0' + byte(b)
}

func solveSmart(signal []byte) string {
	stringSignal := string(signal)

	offset, _ := strconv.Atoi(stringSignal[0:7])
	fmt.Println(offset)

	realSignal := []byte(getRealSignal(stringSignal))
	if len(realSignal) > 2*offset {
		fmt.Println("this would not work!")
		return ""
	}

	offsetSignal := realSignal[offset:]
	n := len(offsetSignal)
	for iteration := 0; iteration < 100; iteration++ {
		fmt.Printf("iteration %d started\n", iteration)
		accum := byte(0)

		for idx := n - 1; idx >= 0; idx-- {
			accum = (accum + charToInt(offsetSignal[idx])) % 10
			offsetSignal[idx] = intToChar(accum)
		}

		fmt.Printf("iteration %d done\n", iteration)
	}

	return string(offsetSignal[:8])
}

func main() {
	// inputSignal := "19617804207202209144916044189917"
	// inputSignal := "1234567812345678"
	inputSignal := util.ReadLines("ch16/input.txt")[0]
	// fmt.Printf("input len: %d\n", len(inputSignal))

	// realSignal := getRealSignal(inputSignal)
	// fmt.Printf("real input len: %d\n", len(realSignal))

	// for _, v := range getPattern(8) {
	// 	fmt.Println(v)
	// }

	// matrix := getPattern(len(inputSignal))

	// fmt.Println(encodeSlow(inputSignal, matrix))
	// fmt.Println(encodeSlow(inputSignal))

	// inputSignal = util.ReadLines("ch16/input.txt")[0]

	// fmt.Println(encodeReal(inputSignal))

	// intMatrix := conv(matrix)

	// for t := 0; t < 4; t++ {
	// 	printMatrix(intMatrix)

	// 	intMatrix = mul(intMatrix, intMatrix)
	// }

	// printMatrix(intMatrix)

	// testAssumption()

	fmt.Println(solveSmart([]byte(inputSignal)))

	fmt.Println("Done")
}
