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

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if words[y][x] != "A" {
				continue
			}

			crossOne := words[y-1][x-1] + words[y][x] + words[y+1][x+1]
			crossTwo := words[y-1][x+1] + words[y][x] + words[y+1][x-1]

			if (crossOne == "MAS" || crossOne == "SAM") && (crossTwo == "MAS" || crossTwo == "SAM") {
				total++
			}
		}
	}

	fmt.Println(total)
}
