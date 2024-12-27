package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cord struct {
	x, y int
}

func main() {
	var theMap [][]int
	var trailHeads []cord

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		newRow := make([]int, len(row))
		for col, value := range row {
			newRow[col], _ = strconv.Atoi(value)

			if newRow[col] == 0 {
				trailHeads = append(trailHeads, cord{col, len(theMap)})
			}
		}

		theMap = append(theMap, newRow)
	}

	total := 0
	for _, trailHead := range trailHeads {
		total += len(findAccessibleSummits(theMap, trailHead))
	}

	fmt.Println(total)
}

func findAccessibleSummits(theMap [][]int, position cord) (summits map[string]cord) {
	summits = make(map[string]cord)

	if theMap[position.y][position.x] == 9 {
		summits[fmt.Sprintf("%v,%v", position.x, position.y)] = position
		return
	}

	// left
	if position.x > 0 && theMap[position.y][position.x-1] == theMap[position.y][position.x]+1 {
		for key, value := range findAccessibleSummits(theMap, cord{position.x - 1, position.y}) {
			summits[key] = value
		}
	}
	// right
	if position.x < len(theMap[position.y])-1 && theMap[position.y][position.x+1] == theMap[position.y][position.x]+1 {
		for key, value := range findAccessibleSummits(theMap, cord{position.x + 1, position.y}) {
			summits[key] = value
		}
	}
	// up
	if position.y > 0 && theMap[position.y-1][position.x] == theMap[position.y][position.x]+1 {
		for key, value := range findAccessibleSummits(theMap, cord{position.x, position.y - 1}) {
			summits[key] = value
		}
	}
	// down
	if position.y < len(theMap)-1 && theMap[position.y+1][position.x] == theMap[position.y][position.x]+1 {
		for key, value := range findAccessibleSummits(theMap, cord{position.x, position.y + 1}) {
			summits[key] = value
		}
	}

	return
}
