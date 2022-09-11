package ch01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartOne() {
	fin, _ := os.Open("ch01/input.txt")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	cnt := 0
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())

		cnt += value/3 - 2
	}

	fmt.Println(cnt)
}

func PartTwo() {
	fin, _ := os.Open("ch01/input.txt")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	cnt := 0
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())

		fmt.Printf(">>> %d\n", value)
		for value > 0 {
			value = value/3 - 2

			if value > 0 {
				fmt.Printf("	= %d\n", value)
				cnt += value
			}
		}
	}

	fmt.Println(cnt)
}
