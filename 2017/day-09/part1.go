package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, i, l int
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
			i++
		case '<':
			garbage = true
		case '>':
			garbage = false
		case '{':
			if !garbage {
				l++
			}
		case '}':
			if !garbage {
				t += l
				l--
			}
		}

		i++
	}

	fmt.Println(t)
}
