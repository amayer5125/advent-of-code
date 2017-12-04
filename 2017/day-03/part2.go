package main

import (
	"fmt"
	"os"
	"strconv"
)

func GetSumSurround(m [11][11]int, p [2]int) (sum int) {
	for x := p[0] - 1; x <= p[0]+1; x++ {
		for y := p[1] - 1; y <= p[1]+1; y++ {
			if x == p[0] && y == p[1] {
				continue
			}

			sum += m[y][x]
		}
	}

	return
}

func main() {
	var d, i int
	var memory [11][11]int

	// get the input
	d, _ = strconv.Atoi(os.Args[1])

	// the current position [x, y]
	p := [2]int{len(memory) / 2, len(memory) / 2}

	memory[p[1]][p[0]] = 1

	// keep track of which direction we are moving in the swirl
	direction := 1
	for {
		i++

		// move left and right
		for x := 0; x < i; x++ {
			p[0] += direction
			memory[p[1]][p[0]] = GetSumSurround(memory, p)

			if memory[p[1]][p[0]] > d {
				fmt.Println(memory[p[1]][p[0]])
				return
			}
		}

		// move up and down
		for y := 0; y < i; y++ {
			p[1] += -direction

			memory[p[1]][p[0]] = GetSumSurround(memory, p)

			if memory[p[1]][p[0]] > d {
				fmt.Println(memory[p[1]][p[0]])
				return
			}
		}

		direction *= -1
	}
}
