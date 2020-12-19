package util

import (
	"strconv"
)

func ConvertStrArrToIntArr(strArr []string) []int {
	res := make([]int, len(strArr))

	for idx, value := range strArr {
		res[idx], _ = strconv.Atoi(value)
	}

	return res
}

func Contains(array []int, el int) bool {
	for _, num := range array {
		if num == el {
			return true
		}
	}
	return false
}

func permuteAux(array []int, tmp []int, result *[][]int) {
	if len(tmp) == len(array) {
		*result = append(*result, tmp)
		return
	}

	for _, num := range array {
		if !Contains(tmp, num) {
			tmpCopy := tmp
			tmpCopy = append(tmpCopy, num)

			permuteAux(array, tmpCopy, result)
		}
	}
}

func GetPermutations(array []int) [][]int {
	var res [][]int

	permuteAux(array, []int{}, &res)

	return res
}

func MyAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MyMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MyMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

func Lcm3(a, b, c int) int {
	return Lcm(a, Lcm(b, c))
}

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

func Gcd3(a, b, c int) int {
	return Gcd(a, Gcd(b, c))
}
