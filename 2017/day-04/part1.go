package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(s string) bool {
	// separate each word
	w := strings.Split(s, " ")
	// keep track of words we have already seen
	u := make(map[string]bool)

	// loop through each word
	for _, v := range w {
		// checkf if we havd seen this word before
		if u[v] {
			return false
		}
		// mark the word as seen
		u[v] = true
	}

	// no two words were the same
	return true
}

func main() {
	var t int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)

	// loop through each row of input
	for scanner.Scan() {
		// check if the passphrase is valid
		if isValid(scanner.Text()) {
			t++
		}
	}

	fmt.Println(t)
}
