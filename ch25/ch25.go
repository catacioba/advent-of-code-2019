package ch25

import (
	"aoc/intcode"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readUntilTimeout(ch chan int) []byte {
	buf := make([]byte, 0)
	for {
		select {
		case x := <-ch:
			buf = append(buf, byte(x))
		case <-time.After(10 * time.Millisecond):
			return buf
		}
	}
}

func readInstructionFromStdin(bin *bufio.Reader) string {
	b, _, _ := bin.ReadLine()
	return string(b)
}

func giveInstruction(ch chan int, in string) {
	for _, c := range in {
		ch <- int(c)
	}
	ch <- 10
}

const (
	south = "south"
	north = "north"
	west  = "west"
	east  = "east"
)

func parseList(s string) []string {
	parts := strings.Split(s, "\n")

	result := make([]string, 0)
	for _, p := range parts {
		result = append(result, strings.TrimPrefix(p, "- "))
	}

	return result
}

func parseDoors(s string) []string {
	pattern := "Doors here lead:"
	startIdx := strings.Index(s, pattern) + len(pattern) + 1
	endIdx := strings.Index(s[startIdx:], "\n\n")

	block := s[startIdx : startIdx+endIdx]
	//fmt.Printf("(%s)\n", s[startIdx:startIdx+endIdx])

	return parseList(block)
}

func parseItems(s string) []string {
	pattern := "Items here:"

	startIdx := strings.Index(s, pattern)
	if startIdx == -1 {
		return make([]string, 0)
	}
	startIdx += len(pattern) + 1
	endIdx := strings.Index(s[startIdx:], "\n\n")

	block := s[startIdx : startIdx+endIdx]
	return parseList(block)
}

func parseBlock(s string) {
	doors := parseDoors(s)
	items := parseItems(s)

	fmt.Printf("doors: %v\nitems: %v\n\n", doors, items)
}

func play(p *intcode.Program) {

}

func PartOne() {
	program := intcode.NewIntCodeProgramFromFile("ch25/input.txt")
	bin := bufio.NewReader(os.Stdin)

	go program.Run()

	for {
		txt := readUntilTimeout(program.Output)
		fmt.Printf("[%s]\n", txt)

		parseBlock(string(txt))

		cmd := readInstructionFromStdin(bin)
		giveInstruction(program.Input, cmd)
	}
}

func PartTwo() {

}
