package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	Direction int
	Clicks    int
}

func main() {
	var instructions []Instruction
	var direction int
	var clicks int
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		if string(instruction[0]) == "L" {
			direction = -1
		} else if string(instruction[0]) == "R" {
			direction = 1
		} else {
			fmt.Fprintf(os.Stderr, "Bad instruction (direction): %s\n", instruction)
			os.Exit(1)
		}

		clicks, err = strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Bad instruction (number): %s\n", instruction)
			os.Exit(1)
		}

		instructions = append(instructions, Instruction{direction, clicks})
	}

	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

func part1(instructions []Instruction) (zeros int) {
	var position int = 50

	for _, instruction := range instructions {
		// ensure modulo result is positive
		position = ((position+(instruction.Clicks*instruction.Direction))%100 + 100) % 100
		if position == 0 {
			zeros += 1
		}
	}

	return
}

func part2(instructions []Instruction) (zeros int) {
	var position int = 50

	for _, instruction := range instructions {
		for range instruction.Clicks {
			// ensure modulo result is positive
			position = ((position+instruction.Direction)%100 + 100) % 100
			if position == 0 {
				zeros += 1
			}
		}
	}

	return
}
