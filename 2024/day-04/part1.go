package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var words [][]string
	var total = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words = append(words, strings.Split(scanner.Text(), ""))
	}

	var width = len(words[0])
	var height = len(words)
	var searchPatterns = [8][3][2]int{
		// forward
		{{1, 0}, {2, 0}, {3, 0}},
		// diagonal down right
		{{1, 1}, {2, 2}, {3, 3}},
		// down
		{{0, 1}, {0, 2}, {0, 3}},
		// diagonal down left
		{{1, -1}, {2, -2}, {3, -3}},
		// backward
		{{-1, 0}, {-2, 0}, {-3, 0}},
		// diagonal up left
		{{-1, -1}, {-2, -2}, {-3, -3}},
		// up
		{{0, -1}, {0, -2}, {0, -3}},
		// diagonal up right
		{{-1, 1}, {-2, 2}, {-3, 3}},
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if words[y][x] != "X" {
				continue
			}

			for _, pattern := range searchPatterns {
				// check if full pattern can be searched
				if x+pattern[2][0] < 0 || x+pattern[2][0] >= width || y+pattern[2][1] < 0 || y+pattern[2][1] >= height {
					continue
				}

				if words[y+pattern[0][1]][x+pattern[0][0]]+words[y+pattern[1][1]][x+pattern[1][0]]+words[y+pattern[2][1]][x+pattern[2][0]] == "MAS" {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}
