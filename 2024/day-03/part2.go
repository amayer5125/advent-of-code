package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var total = 0
	var active = true

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		instructionPattern := regexp.MustCompile(`(do|don't|mul)\((?:(\d{1,3}),(\d{1,3}))?\)`)
		matches := instructionPattern.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if !active {
				active = match[1] == "do"
				continue
			}

			if match[1] == "don't" {
				active = false
				continue
			}

			if match[1] == "mul" && len(match) == 4 {
				firstParameter, _ := strconv.Atoi(match[2])
				secondParameter, _ := strconv.Atoi(match[3])
				total += firstParameter * secondParameter
			}
		}
	}

	fmt.Println(total)
}
