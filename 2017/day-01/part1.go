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

		// loop through the input
		for i := 0; i < len(d); i++ {
			n := (i + 1) % len(d)

			if d[i:i+1] == d[n:n+1] {
				c, _ := strconv.Atoi(d[i : i+1])
				t += c
			}
		}
	}

	fmt.Println(t)
}
