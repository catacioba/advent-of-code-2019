package ch08

import (
	"aoc/util"
	"fmt"
	"math"
)

func decodeSpaceImage(img []int, height, width int) [][]int {
	pixelsPerLayer := height * width

	layers := len(img) / pixelsPerLayer

	decodedImage := make([][]int, height)

	for i := 0; i < height; i++ {
		decodedImage[i] = make([]int, width)
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			for layer := 0; layer < layers; layer++ {

				pixel := img[layer*pixelsPerLayer+x*width+y]
				decodedImage[x][y] = pixel

				if pixel != 2 {
					break
				}
			}
		}
	}

	return decodedImage
}

func imageCheckSum(img []int, height, width int) int {
	pixelsPerLayer := height * width

	layers := len(img) / pixelsPerLayer

	minLayer := math.MaxInt32
	var minLayerCheckSum int

	for layer := 0; layer < layers; layer++ {
		zeroCnt := 0
		oneCnt := 0
		twoCnt := 0

		start := layer * pixelsPerLayer
		end := start + pixelsPerLayer

		for idx := start; idx < end; idx++ {
			switch img[idx] {
			case 0:
				zeroCnt++
			case 1:
				oneCnt++
			case 2:
				twoCnt++
			}
		}

		if zeroCnt < minLayer {
			minLayer = zeroCnt
			minLayerCheckSum = oneCnt * twoCnt
		}
	}

	return minLayerCheckSum
}

func convertStringToArray(str string) []int {
	res := make([]int, len(str))

	for idx, chr := range str {
		res[idx] = int(chr) - '0'
	}

	return res
}

func printImage(img [][]int) {
	for _, row := range img {
		for _, pixel := range row {
			if pixel == 1 {
				fmt.Printf("%s%d%s", util.Red, pixel, util.NoColor)
			} else {
				fmt.Printf("%d", pixel)
			}
		}
		fmt.Println()
	}
}

func Solve() {
	line := util.ReadLines("ch08/input.txt")[0]

	image := convertStringToArray(line)

	//fmt.Println(image)

	//fmt.Println(imageCheckSum(image, 6, 25))

	img := decodeSpaceImage(image, 6, 25)

	printImage(img)
}
