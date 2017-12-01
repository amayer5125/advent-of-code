package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_file(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}

	return strings.TrimSpace(string(data))
}

func main() {
	directions := strings.Split(read_file("input.txt"), ", ")

	cords := [2]int{0, 0}
	// totally not robot directions on how to move in a grid
	modifiers := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	modifier_index := 0
	// cords we visited
	visited := make(map[string]bool)

	for _, dir := range directions {
		turn := dir[:1]
		dist, _ := strconv.Atoi(dir[1:])

		// move 1 pointer forward or backwards wrapping around boundaries
		if turn == "R" {
			modifier_index = (modifier_index + 1) % 4
		} else {
			modifier_index = (modifier_index + 3) % 4
		}

		for i := 0; i < dist; i++ {
			cords[0] += modifiers[modifier_index][0]
			cords[1] += modifiers[modifier_index][1]

			cur_cords := strconv.Itoa(cords[0]) + "," + strconv.Itoa(cords[1])

			if visited[cur_cords] {
				fmt.Println(math.Abs(float64(cords[0])) + math.Abs(float64(cords[1])))
				return
			}
			visited[cur_cords] = true
		}
	}
}
