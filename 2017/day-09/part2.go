package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, i int
	var d []byte
	var garbage bool

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d = scanner.Bytes()
	}

	// decide what to do with the character
	for i < len(d) {
		switch d[i] {
		case '!':
			i += 2
			continue
		case '<':
			garbage = true
		case '>':
			t--
			garbage = false
		}

		if garbage {
			t++
		}

		i++
	}

	fmt.Println(t)
}
