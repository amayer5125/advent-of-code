package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)

	// loop through each row of input
	for scanner.Scan() {
		var h, l int

		// split the row into slice
		r := strings.Split(scanner.Text(), "\t")

		// loop through row fields
		for k, v := range r {
			n, _ := strconv.Atoi(v)

			// loop through slice again to compare values
			for k2, v2 := range s {
				// do not compare value against itself
				if k2 == k {
					continue
				}

				n2, _ := strconv.Atoi(v2)
				// check if the numbers are divisible
				if n%n2 == 0 {
					h = n
					l = n2
					break
				}
			}

			// if high is set we have all the info we need from this set
			if h != 0 {
				break
			}
		}

		t += h / l
	}

	fmt.Println(t)
}
