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

	roomWidth := len(room[0])
	roomHeight := len(room)
	visited := make(map[string]bool)

	for {
		visited[fmt.Sprintf("%v,%v", guard.location.x, guard.location.y)] = true

		nextX := guard.location.x + guard.direction.xMovement
		nextY := guard.location.y + guard.direction.yMovement
		if nextX < 0 || nextX >= roomWidth || nextY < 0 || nextY >= roomHeight {
			break
		}

		if room[nextY][nextX] == "#" {
			guard.direction = guard.direction.turnRight()
			continue
		}

		guard.location.x = nextX
		guard.location.y = nextY
	}

	fmt.Println(len(visited))
}
