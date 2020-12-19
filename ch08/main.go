package ch08

import (
	"aoc/util"
	"fmt"
	"github.com/fatih/color"
	"math"
)

func decodeSpaceImage(img []int, height, width int) [][]int {
	pixelsPerLayer := height * width

	layers := len(img) / pixelsPerLayer

	decodedImage := make([][]int, height)

	for i := 0; i < height; i++ {
		decodedImage[i] = make([]int, width)
	}

	//for layer := 0; layer < layers; layer++ {
	//	it := layer * pixelsPerLayer
	//
	//	layerPixels := make()
	//}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			//fmt.Printf("### x=%d y=%d ###\n", x, y)
			for layer := 0; layer < layers; layer++ {
				//start := layer * pixelsPerLayer
				//end := start + pixelsPerLayer

				pixel := img[layer*pixelsPerLayer+x*width+y]
				//fmt.Println(pixel)
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
	red := color.New(color.FgRed).PrintfFunc()
	black := color.New(color.FgBlack).PrintfFunc()

	for _, row := range img {
		for _, pixel := range row {
			if pixel == 1 {
				red("%d ", pixel)
			} else {
				black("%d ", pixel)
			}
		}
		fmt.Println()
	}
}

func main() {

	line := util.ReadLines("ch08/input.txt")[0]

	image := convertStringToArray(line)

	//fmt.Println(image)

	//fmt.Println(imageCheckSum(image, 6, 25))

	img := decodeSpaceImage(image, 6, 25)

	printImage(img)
}
