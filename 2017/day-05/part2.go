package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t, o, n int
	var i []int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())

		// add input to instructions
		i = append(i, d)
	}

	// loop until the offset is out of bounds
	for o < len(i) && o >= 0 {
		// get the value at the current offset
		n = i[o]

		// increment or decrement the current instruction
		if n >= 3 {
			i[o]--
		} else {
			i[o]++
		}

		// update the offset to the next jump
		o += n

		// increment the number of jumps
		t++
	}

	fmt.Println(t)
}
