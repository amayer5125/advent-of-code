package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, stone := range strings.Split(scanner.Text(), " ") {
			stones[stone]++
		}
	}

	for i := 0; i < 75; i++ {
		stones = stonesAfterBlink(stones)
	}

	total := 0
	for _, count := range stones {
		total += count
	}
	fmt.Println(total)
}

func stonesAfterBlink(stones map[string]int) map[string]int {
	newStones := make(map[string]int)

	for number, count := range stones {
		if number == "0" {
			newStones["1"] += count
			continue
		}

		if len(number)%2 == 0 {
			middle := len(number) / 2
			newStones[number[:middle]] += count
			rightNumber := strings.TrimLeft(number[middle:], "0")
			if rightNumber == "" {
				rightNumber = "0"
			}

			newStones[rightNumber] += count
			continue
		}

		newNumber, _ := strconv.Atoi(number)
		newStones[strconv.Itoa(newNumber*2024)] += count
	}

	return newStones
}
