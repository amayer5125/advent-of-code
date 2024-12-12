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
				potentialAntinodes := calculateAntinodes(antennaFrequencies[i], antennaFrequencies[j])
				for _, potentialAntinode := range potentialAntinodes {
					if potentialAntinode.x < 0 || potentialAntinode.x >= width {
						continue
					}
					if potentialAntinode.y < 0 || potentialAntinode.y >= height {
						continue
					}

					antinodes[fmt.Sprintf("%v,%v", potentialAntinode.x, potentialAntinode.y)] = true
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func calculateAntinodes(antennaOne cord, antennaTwo cord) [2]cord {
	xDiff := antennaTwo.x - antennaOne.x
	yDiff := antennaTwo.y - antennaOne.y
	return [2]cord{
		{antennaOne.x - xDiff, antennaOne.y - yDiff},
		{antennaTwo.x + xDiff, antennaTwo.y + yDiff},
	}
}
