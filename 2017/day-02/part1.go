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
		for _, v := range r {
			n, _ := strconv.Atoi(v)

			if n > h {
				h = n
			}
			if l == 0 || n < l {
				l = n
			}
		}

		// add difference to total
		t += h - l
	}

	fmt.Println(t)
}
