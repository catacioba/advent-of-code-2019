package main

import (
	"fmt"
	"strconv"
)

func isValidPartOne(key string) bool {
	if len(key) != 6 {
		return false
	}

	decreases := false
	hasDouble := false

	idx := 1
	for idx < len(key) {
		if key[idx] == key[idx-1] {
			hasDouble = true
		}
		if key[idx] < key[idx-1] {
			decreases = true
		}
		idx++
	}

	return !decreases && hasDouble
}

func PartOne() {

	cnt := 0

	for key := 254032; key < 789860; key++ {
		if isValidPartOne(strconv.Itoa(key)) {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func isValidPartTwo(key string) bool {
	if len(key) != 6 {
		return false
	}

	decreases := false
	hasDouble := false

	idx := 1
	for idx < len(key) {
		if key[idx] == key[idx-1] && !(idx >= 2 && key[idx] == key[idx-2]) && !(idx < len(key)-1 && key[idx] == key[idx+1]) {
			hasDouble = true
		}
		if key[idx] < key[idx-1] {
			decreases = true
		}
		idx++
	}

	return !decreases && hasDouble
}

func PartTwo() {

	cnt := 0

	for key := 254032; key < 789860; key++ {
		if isValidPartTwo(strconv.Itoa(key)) {
			cnt++
		}
	}

	fmt.Println(cnt)

	fmt.Println(isValidPartTwo("112233"))
	fmt.Println(isValidPartTwo("123444"))
	fmt.Println(isValidPartTwo("111122"))
}
