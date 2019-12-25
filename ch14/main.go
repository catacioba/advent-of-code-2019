package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Chemical struct {
	name     string
	quantity int
	isBase   bool
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

func readReactions(inputFile string) (map[string]Reaction, map[string]bool) {
	lines := util.ReadLines(inputFile)

	reactions := make(map[string]Reaction)
	baseChemicals := make(map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, " => ")

		leftPart := parts[0]
		rightPart := parts[1]

		leftChemicals := processChemicals(leftPart)
		rightChemical := processChemicals(rightPart)[0]

		if len(leftChemicals) == 1 && leftChemicals[0].name == "ORE" {
			baseChemicals[rightChemical.name] = true
		}

		reactions[rightChemical.name] = Reaction{
			inputChemicals: leftChemicals,
			outputChemical: rightChemical,
		}
	}

	return reactions, baseChemicals
}

//func getOre(reactions map[string]Reaction, chem string, quantity int) int {
func getOre(reactions map[string]Reaction, chemical Chemical) int {
	chemReaction := reactions[chemical.name]

	reactionCount := int(math.Ceil(float64(chemical.quantity) / float64(chemReaction.outputChemical.quantity)))

	oreCount := 0

	for _, inputChemical := range chemReaction.inputChemicals {
		if inputChemical.name == "ORE" {
			oreCount += inputChemical.quantity
		} else {
			tmp := getOre(reactions, inputChemical)

			oreCount += tmp
		}
	}

	return reactionCount * oreCount
}

func getChemicals(reactions map[string]Reaction, chemical Chemical) map[string]Chemical {
	requiredChemicals := make(map[string]Chemical)

	chemReaction := reactions[chemical.name]

	reactionCount := int(math.Ceil(float64(chemical.quantity) / float64(chemReaction.outputChemical.quantity)))

	for _, inputChemical := range chemReaction.inputChemicals {
		//if inputChemical.name == "ORE" {
		//	oreCount += inputChemical.quantity
		//} else {
		//	tmp := getOre(reactions, inputChemical)
		//
		//	oreCount += tmp
		//}
		//
		//tmp := getChemicals(reactions, inputChemical)
		//
		//for idx := 0; idx < reactionCount; idx++ {
		//	requiredChemicals
		//}

		inputReaction := reactions[inputChemical.name]

		if !inputReaction.isBase {
			tmpRequired := getChemicals(reactions, inputChemical)

			for k, v := range tmpRequired {
				chem, ok := requiredChemicals[k]

				if ok {
					chem.quantity += reactionCount * v.quantity
					requiredChemicals[k] = chem
				} else {
					requiredChemicals[k] = v
				}
			}
		} else {

		}
	}

	return requiredChemicals
}

func main() {

	reactions, _ := readReactions("ch14/input.txt")

	//for _, reaction := range reactions {
	//	fmt.Println(reaction)
	//}

	fmt.Println(getChemicals(reactions, Chemical{
		name:     "FUEL",
		quantity: 1,
	}))
}
