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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		instructionPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := instructionPattern.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			firstParameter, _ := strconv.Atoi(match[1])
			secondParameter, _ := strconv.Atoi(match[2])
			total += firstParameter * secondParameter
		}
	}

	fmt.Println(total)
}
