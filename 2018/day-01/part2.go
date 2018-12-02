package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var index, total int
	var input []int
	seen := make(map[int]bool)

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		freqChange, _ := strconv.Atoi(scanner.Text())

		// append frequency change to our input
		input = append(input, freqChange)
	}

	for {
		total += input[index]

		// check if we have seen this total before
		if seen[total] {
			break
		}

		// mark total as seen
		seen[total] = true

		// increment index wrapping on input length
		index = (index + 1) % len(input)
	}

	// output total
	fmt.Println(total)
}
