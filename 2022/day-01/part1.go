package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var current int
	var max int

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			if current > max {
				max = current
			}

			current = 0
			continue
		}

		calories, _ := strconv.Atoi(text)
		current += calories
	}

	fmt.Println(max)
}
