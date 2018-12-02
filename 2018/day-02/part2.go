package main

import (
	"bufio"
	"fmt"
	"os"
)

func inCommon(x, y string) (int, string) {
	var diff int
	var common []byte
	// loop through each letter one at a time
	for i := 0; i < len(x); i++ {
		// check if the letters are different
		if x[i] != y[i] {
			diff++
			continue
		}

		common = append(common, x[i])
	}

	return diff, string(common)
}

func main() {
	var inputs []string

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	// loop through each input (except for the last one)
	for i := 0; i < len(inputs)-1; i++ {
		// loop through each input after the current one
		for j := i + 1; j < len(inputs); j++ {
			// get the differences in the strings
			diff, common := inCommon(inputs[i], inputs[j])
			// check if the strings only have one letter different-
			if diff == 1 {
				fmt.Println(common)
				return
			}
		}
	}
}
