package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Coordinates struct {
	x, y int
}

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}

	return strings.TrimSpace(string(data))
}

func InBounds(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func main() {
	directions := strings.Split(ReadFile("input.txt"), "\n")

	passcode := ""
	keypad := [5][5]string{
		{" ", " ", "1", " ", " "},
		{" ", "2", "3", "4", " "},
		{"5", "6", "7", "8", "9"},
		{" ", "A", "B", "C", " "},
		{" ", " ", "D", " ", " "},
	}
	cords := Coordinates{x: 0, y: 2}
	// totally not robot directions on how to move in a grid
	modifiers := map[string][2]int{
		"U": {0, -1},
		"R": {1, 0},
		"D": {0, 1},
		"L": {-1, 0},
	}

	for _, d := range directions {
		d := strings.Split(d, "")
		for _, cd := range d {
			new_x := InBounds(cords.x+modifiers[cd][0], 0, 4)
			new_y := InBounds(cords.y+modifiers[cd][1], 0, 4)
			if keypad[new_y][new_x] != " " {
				cords.x = new_x
				cords.y = new_y
			}
		}
		passcode += keypad[cords.y][cords.x]
	}
	fmt.Println(passcode)
}
