package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var twosAndThrees [2]int

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		letters := scanner.Text()

		hasTwo := false
		hasThree := false
		for i := 0; i < len(letters); i++ {
			count := strings.Count(letters, string(letters[i]))
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twosAndThrees[0]++
		}
		if hasThree {
			twosAndThrees[1]++
		}
	}

	fmt.Println(twosAndThrees[0] * twosAndThrees[1])
}
