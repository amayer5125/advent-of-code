package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var d string
	var t int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d = scanner.Text()

		// get the half length
		h := len(d) / 2

		// loop through half the string
		// if they didn't match the first time they will not match the second
		for i := 0; i < h; i++ {
			n := (i + h) % len(d)

			if d[i:i+1] == d[n:n+1] {
				c, _ := strconv.Atoi(d[i : i+1])
				t += c * 2
			}
		}
	}

	fmt.Println(t)
}
