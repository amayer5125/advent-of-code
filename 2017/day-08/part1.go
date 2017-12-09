package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var l int
	var registers = make(map[string]int)
	var instruction = regexp.MustCompile(`([a-z]+) (inc|dec) (-?\d+) if ([a-z]+) ([<>=!]+) (-?\d+)`)

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// separate the instructions into parts
		parts := instruction.FindStringSubmatch(scanner.Text())

		// get the value we will be compairing against
		x, _ := strconv.Atoi(parts[6])

		// does the register meet the requirements
		switch parts[5] {
		case "==":
			if registers[parts[4]] != x {
				continue
			}
		case "!=":
			if registers[parts[4]] == x {
				continue
			}
		case ">":
			if registers[parts[4]] <= x {
				continue
			}
		case ">=":
			if registers[parts[4]] < x {
				continue
			}
		case "<":
			if registers[parts[4]] >= x {
				continue
			}
		case "<=":
			if registers[parts[4]] > x {
				continue
			}
		}

		i, _ := strconv.Atoi(parts[3])
		if parts[2] == "dec" {
			i *= -1
		}

		registers[parts[1]] += i
	}

	for _, v := range registers {
		if v > l {
			l = v
		}
	}
	fmt.Println(l)
}
