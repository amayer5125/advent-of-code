package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(id int) bool {
	return id >= r.Start && id <= r.End
}

func main() {
	var importingIngredients bool = false
	var freshIngredientRanges []Range
	var ingredients []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if !importingIngredients {
			if line == "" {
				importingIngredients = true
				continue
			}

			idRange := strings.Split(line, "-")
			start, err := strconv.Atoi(idRange[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			end, err := strconv.Atoi(idRange[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			freshIngredientRanges = append(freshIngredientRanges, Range{start, end})

			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		ingredients = append(ingredients, id)
	}

	freshIngredientRanges = consolidateRanges(freshIngredientRanges)

	fmt.Println(part1(freshIngredientRanges, ingredients))
	fmt.Println(part2(freshIngredientRanges))
}

func part1(freshIngredientRanges []Range, ingredients []int) (freshIngredients int) {
	for _, id := range ingredients {
		for _, idRange := range freshIngredientRanges {
			if idRange.Contains(id) {
				freshIngredients++
				break
			}
		}
	}
	return
}

func part2(freshIngredientRanges []Range) (freshIngredientIds int) {
	for _, idRange := range freshIngredientRanges {
		freshIngredientIds += idRange.End - idRange.Start + 1
	}
	return
}

func consolidateRanges(idRanges []Range) (newRanges []Range) {
	var currentRange Range
	var checkRange Range

	for currentRangeIndex := 0; currentRangeIndex < len(idRanges); currentRangeIndex++ {
		currentRange = idRanges[currentRangeIndex]

		for checkRangeIndex := currentRangeIndex + 1; checkRangeIndex < len(idRanges); checkRangeIndex++ {
			checkRange = idRanges[checkRangeIndex]

			if checkRange.Contains(currentRange.Start) && checkRange.Contains(currentRange.End) {
				currentRange.Start = checkRange.Start
				currentRange.End = checkRange.End

				idRanges = append(idRanges[:checkRangeIndex], idRanges[checkRangeIndex+1:]...)
				checkRangeIndex = currentRangeIndex
				continue
			}

			if currentRange.Contains(checkRange.Start) || currentRange.Contains(checkRange.End) {
				if checkRange.Start < currentRange.Start {
					currentRange.Start = checkRange.Start
				}
				if checkRange.End > currentRange.End {
					currentRange.End = checkRange.End
				}

				idRanges = append(idRanges[:checkRangeIndex], idRanges[checkRangeIndex+1:]...)
				checkRangeIndex = currentRangeIndex
				continue
			}
		}

		newRanges = append(newRanges, currentRange)
	}

	return
}
