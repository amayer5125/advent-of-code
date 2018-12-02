package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total int

	// read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		freqChange, _ := strconv.Atoi(scanner.Text())

		// add frequency change to our total
		total += freqChange
	}

	fmt.Println(total)
}
