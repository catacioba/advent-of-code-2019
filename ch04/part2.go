package main

import (
	"fmt"
	"strconv"
)

func isValid2(key string) bool {
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

func main() {

	cnt := 0

	for key := 254032; key < 789860; key++ {
		if isValid2(strconv.Itoa(key)) {
			cnt++
		}
	}

	fmt.Println(cnt)

	fmt.Println(isValid2("112233"))
	fmt.Println(isValid2("123444"))
	fmt.Println(isValid2("111122"))
}
