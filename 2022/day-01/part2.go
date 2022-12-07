package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var calories int
	totalCalories := []int{0, 0, 0}

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if calories > totalCalories[0] {
				totalCalories[0] = calories
				sort.Ints(totalCalories)
			}

			calories = 0

			continue
		}

		value, _ := strconv.Atoi(text)
		calories += value
	}

	fmt.Println(sum(totalCalories))
}

func sum(input []int) int {
	sum := 0

	for _, i := range input {
		sum += i
	}

	return sum
}
