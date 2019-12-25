package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fin, _ := os.Open("ch01/ch01.txt")
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	cnt := 0
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())

		cnt += value/3 - 2
	}

	fmt.Println(cnt)
}
