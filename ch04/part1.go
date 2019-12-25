package main

import (
	"fmt"
	"strconv"
)

func isValid(key string) bool {
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

func main() {

	cnt := 0

	for key := 254032; key < 789860; key++ {
		if isValid(strconv.Itoa(key)) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
