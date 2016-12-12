package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_lines(filename string) []string {
	conts, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return []string{}
	}

	data := strings.Split(string(conts), "\n")

	return data[0 : len(data)-1]
}

func main() {
	directions := read_lines("input.txt")

	var screen [6][50]string
	rect_patt, _ := regexp.Compile(`^rect (\d+)x(\d+)$`)
	rotate_col_patt, _ := regexp.Compile(`^rotate column x=(\d+) by (\d+)$`)
	rotate_row_patt, _ := regexp.Compile(`^rotate row y=(\d+) by (\d+)$`)

	for _, direction := range directions {
		if rect_patt.MatchString(direction) {
			cords := rect_patt.FindAllStringSubmatch(direction, -1)
			cord_x, _ := strconv.Atoi(cords[0][1])
			cord_y, _ := strconv.Atoi(cords[0][2])

			for x := 0; x < cord_x; x++ {
				for y := 0; y < cord_y; y++ {
					// use block to represent pixles
					screen[y][x] = "\u2588"
				}
			}
			continue
		}
		if rotate_col_patt.MatchString(direction) {
			cords := rotate_col_patt.FindAllStringSubmatch(direction, -1)
			col_x, _ := strconv.Atoi(cords[0][1])
			shift_by, _ := strconv.Atoi(cords[0][2])

			var old_col [len(screen)]string
			// store old values from col
			for i, row := range screen {
				old_col[i] = row[col_x]
			}

			// shift each col
			for i := 0; i < len(old_col); i++ {
				old_slot := (i - shift_by) % len(old_col)
				if old_slot < 0 {
					old_slot += len(old_col)
				}
				screen[i][col_x] = old_col[old_slot]
			}
			continue
		}
		if rotate_row_patt.MatchString(direction) {
			cords := rotate_row_patt.FindAllStringSubmatch(direction, -1)
			row_y, _ := strconv.Atoi(cords[0][1])
			shift_by, _ := strconv.Atoi(cords[0][2])

			old_row := screen[row_y]

			// shift each col
			for i := 0; i < len(old_row); i++ {
				old_slot := (i - shift_by) % len(old_row)
				if old_slot < 0 {
					old_slot += len(old_row)
				}
				screen[row_y][i] = old_row[old_slot]
			}
			continue
		}
	}

	// output the screen
	for _, row := range screen {
		for x, _ := range row {
			if row[x] == "" {
				fmt.Print(" ")
				continue
			}
			fmt.Print(row[x])
		}
		fmt.Print("\n")
	}
}
