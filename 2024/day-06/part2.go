package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type guard struct {
	location  location
	direction direction
}

type location struct {
	x, y int
}

type direction struct {
	xMovement, yMovement int
}

func (d direction) turnRight() direction {
	// [0, -1]
	// [1, 0]
	// [0, 1]
	// [-1, 0]
	if d.xMovement == 0 {
		return direction{d.yMovement * -1, 0}
	}

	return direction{0, d.xMovement}
}

func main() {
	var room [][]string
	var guard = guard{location{0, 0}, direction{0, -1}}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		room = append(room, row)

		guardIndex := slices.Index(row, "^")
		if guardIndex > -1 {
			guard.location.x = guardIndex
			guard.location.y = len(room) - 1
		}
	}

	var loops int
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[y]); x++ {
			if room[y][x] == "." {
				room[y][x] = "#"
				if detectLoop(room, guard) {
					loops++
				}
				room[y][x] = "."
			}
		}
	}
	fmt.Println(loops)
}

func detectLoop(room [][]string, guard guard) bool {
	roomWidth := len(room[0])
	roomHeight := len(room)
	visited := make(map[string]bool)

	for {
		cords := fmt.Sprintf("%v,%v,%v,%v", guard.location.x, guard.location.y, guard.direction.xMovement, guard.direction.yMovement)
		if alreadyTraveled, ok := visited[cords]; ok && alreadyTraveled {
			return true
		}
		visited[cords] = true

		nextX := guard.location.x + guard.direction.xMovement
		nextY := guard.location.y + guard.direction.yMovement
		if nextX < 0 || nextX >= roomWidth || nextY < 0 || nextY >= roomHeight {
			return false
		}

		if room[nextY][nextX] == "#" {
			guard.direction = guard.direction.turnRight()
			continue
		}

		guard.location.x = nextX
		guard.location.y = nextY
	}
}
