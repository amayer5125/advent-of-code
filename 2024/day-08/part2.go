package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cord struct {
	x, y int
}

func main() {
	antennaLocations := make(map[string][]cord)
	var width int
	var height int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "")
		for width = 0; width < len(values); width++ {
			if values[width] != "." {
				antennaLocations[values[width]] = append(antennaLocations[values[width]], cord{width, height})
			}
		}
		height++
	}

	antinodes := make(map[string]bool)
	for _, antennaFrequencies := range antennaLocations {
		for i := 0; i < len(antennaFrequencies)-1; i++ {
			for j := i + 1; j < len(antennaFrequencies); j++ {
				antinodes[fmt.Sprintf("%v,%v", antennaFrequencies[i].x, antennaFrequencies[i].y)] = true
				antinodes[fmt.Sprintf("%v,%v", antennaFrequencies[j].x, antennaFrequencies[j].y)] = true

				potentialAntinodes := calculateAntinodes(antennaFrequencies[i], antennaFrequencies[j], width, height)
				for _, potentialAntinode := range potentialAntinodes {
					antinodes[fmt.Sprintf("%v,%v", potentialAntinode.x, potentialAntinode.y)] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func calculateAntinodes(antennaOne cord, antennaTwo cord, width int, height int) (antinodes []cord) {
	xDiff := antennaTwo.x - antennaOne.x
	yDiff := antennaTwo.y - antennaOne.y

	currentPos := antennaOne
	for {
		currentPos.x -= xDiff
		currentPos.y -= yDiff
		if currentPos.x < 0 || currentPos.x >= width || currentPos.y < 0 || currentPos.y >= height {
			break
		}

		antinodes = append(antinodes, currentPos)
	}
	currentPos = antennaTwo
	for {
		currentPos.x += xDiff
		currentPos.y += yDiff
		if currentPos.x < 0 || currentPos.x >= width || currentPos.y < 0 || currentPos.y >= height {
			break
		}

		antinodes = append(antinodes, currentPos)
	}

	return
}
