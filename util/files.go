package util

import (
	"bufio"
	"os"
)

func ReadLines(filename string) []string {
	fin, err := os.Open(filename)
	Must(err)
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
