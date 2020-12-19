package main

import (
	"aoc/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Chemical struct {
	name     string
	quantity int
}

type Reaction struct {
	inputChemicals []Chemical
	outputChemical Chemical
}

func processChemicals(chemicalPart string) []Chemical {
	chemicals := strings.Split(chemicalPart, ",")

	result := make([]Chemical, len(chemicals))

	for idx := range chemicals {
		chem := strings.TrimSpace(chemicals[idx])

		parts := strings.Split(chem, " ")

		chemicalQuantity, _ := strconv.Atoi(parts[0])

		result[idx] = Chemical{
			name:     parts[1],
			quantity: chemicalQuantity,
		}
	}

	return result
}

func readReactions(inputFile string) (map[string]Reaction, map[string]struct{}) {
	lines := util.ReadLines(inputFile)

	reactions := make(map[string]Reaction)
	baseChemicals := make(map[string]struct{})

	for _, line := range lines {
		parts := strings.Split(line, " => ")

		leftPart := parts[0]
		rightPart := parts[1]

		leftChemicals := processChemicals(leftPart)
		rightChemical := processChemicals(rightPart)[0]

		if len(leftChemicals) == 1 && leftChemicals[0].name == "ORE" {
			baseChemicals[rightChemical.name] = struct{}{}
		}

		_, isAlready := reactions[rightChemical.name]
		if isAlready {
			panic("there is already a reaction for this chemical")
		}

		reactions[rightChemical.name] = Reaction{
			inputChemicals: leftChemicals,
			outputChemical: rightChemical,
		}
	}

	baseChemicals["ORE"] = struct{}{}

	return reactions, baseChemicals
}

// func getOre(reactions map[string]Reaction, chem string, quantity int) int {
func getOre(reactions map[string]Reaction, requiredChemicals map[string]Chemical) int {
	oreCount := 0

	for chemName, chem := range requiredChemicals {
		chemReaction := reactions[chemName]

		reactionCount := int(math.Ceil(float64(chem.quantity) / float64(chemReaction.outputChemical.quantity)))

		amount := reactionCount * chemReaction.inputChemicals[0].quantity

		fmt.Printf("%d %v %v %d\n", reactionCount, chem, chemReaction, amount)

		oreCount += amount
	}

	return oreCount
}

func getOrDefault(requiredChemicals map[string]Chemical, chemicalName string) Chemical {
	alreadyHave, ok := requiredChemicals[chemicalName]

	if !ok {
		alreadyHave = Chemical{
			name:     chemicalName,
			quantity: 0,
		}
	}

	return alreadyHave
}

func getChemicals(reactions map[string]Reaction, baseChemicals map[string]struct{}, chemical Chemical) map[string]Chemical {

	requiredChemicals := make(map[string]Chemical)

	chemReaction := reactions[chemical.name]

	reactionCount := int(math.Ceil(float64(chemical.quantity) / float64(chemReaction.outputChemical.quantity)))

	for _, inputChemical := range chemReaction.inputChemicals {
		_, isBaseChemical := baseChemicals[inputChemical.name]

		// if inputChemical.name == "ORE" {
		// 	isBaseChemical = true
		// }

		if isBaseChemical {
			// we have a base chemical here
			alreadyHave := getOrDefault(requiredChemicals, inputChemical.name)

			alreadyHave.quantity += reactionCount * inputChemical.quantity

			requiredChemicals[inputChemical.name] = alreadyHave
		} else {
			// go recursively until we have a base chemical here
			copyChemical := inputChemical
			copyChemical.quantity *= reactionCount

			subRequiredChemicals := getChemicals(reactions, baseChemicals, copyChemical)

			fmt.Printf("*** %v is made with %v\n", copyChemical, subRequiredChemicals)

			for subChemName, subChem := range subRequiredChemicals {
				alreadyHave := getOrDefault(requiredChemicals, subChemName)

				alreadyHave.quantity += subChem.quantity

				requiredChemicals[subChemName] = alreadyHave
			}
		}
	}

	return requiredChemicals
}

func putIfAbsent(m map[string]int, value string) {
	_, ok := m[value]

	if !ok {
		m[value] = 0
	}
}

func topologicalSortSolve(reactions map[string]Reaction, fuelQuantity int) int {
	fmt.Println()

	for _, v := range reactions {
		fmt.Println(v)
	}

	inDegrees := make(map[string]int, len(reactions))

	for chemName, reaction := range reactions {
		// inDegrees[chemName] = 0
		putIfAbsent(inDegrees, chemName)

		for _, inChem := range reaction.inputChemicals {
			putIfAbsent(inDegrees, inChem.name)

			inDegrees[inChem.name]++
		}
	}

	queue := util.NewQueue()

	for chemName, chemInDegree := range inDegrees {
		if chemInDegree == 0 {
			queue.Push(chemName)
		}
	}

	requiredChemicals := make(map[string]Chemical)
	requiredChemicals["FUEL"] = Chemical{
		name:     "FUEL",
		quantity: fuelQuantity,
	}

	for queue.Size() != 0 {
		el := queue.Pop().(string)

		// fmt.Printf(" %s ", el)
		var requiredQuantity int

		requiredQuantity = requiredChemicals[el].quantity
		// requiredChem, ok := requiredChemicals[el]
		// if ok {
		// requiredQuantity = requiredChem.quantity
		// } else {
		// requiredQuantity = 1
		// }

		reaction := reactions[el]

		reactionCount := int(math.Ceil(float64(requiredQuantity) / float64(reaction.outputChemical.quantity)))

		for _, inChem := range reaction.inputChemicals {
			inDegrees[inChem.name]--

			// reqChemName, reqChem := requiredChemicals[inChem.name]

			existingReqChem, ok := requiredChemicals[inChem.name]

			copyChem := inChem

			if ok {
				copyChem.quantity = copyChem.quantity*reactionCount + existingReqChem.quantity
			} else {
				copyChem.quantity = copyChem.quantity * reactionCount
			}

			requiredChemicals[inChem.name] = copyChem

			if inDegrees[inChem.name] == 0 {
				queue.Push(inChem.name)
			}
		}

		if el == "ORE" {
			break
		}

		delete(requiredChemicals, el)
	}

	// fmt.Printf("\n\n@@ %v\n", requiredChemicals)
	return requiredChemicals["ORE"].quantity
}

const oneTrillion = 1000000000000

func findTrillionOreFuelCount(reactions map[string]Reaction) int {
	start := 0
	end := 3 * oneTrillion

	ans := -1

	for start <= end {
		mid := (start + end) / 2

		requiredOre := topologicalSortSolve(reactions, mid)

		if requiredOre > oneTrillion {
			end = mid - 1
		} else {
			ans = mid
			start = mid + 1
		}
	}

	return ans
}

func main() {

	reactions, baseChemicals := readReactions("ch14/input.txt")

	fmt.Println("Reactions:")
	for _, reaction := range reactions {
		fmt.Println(reaction)
	}
	fmt.Println()
	fmt.Println("Base chemicals:")
	fmt.Println(baseChemicals)
	fmt.Println()

	requiredChemicals := getChemicals(reactions, baseChemicals, Chemical{
		name:     "FUEL",
		quantity: 1,
	})
	fmt.Println(requiredChemicals)
	fmt.Println()

	fmt.Println(getOre(reactions, requiredChemicals))

	fmt.Println(topologicalSortSolve(reactions, 1))

	fmt.Println(findTrillionOreFuelCount(reactions))
}
