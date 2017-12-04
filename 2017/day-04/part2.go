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
		// loop through each word we have seen
		for k, _ := range u {
			// skip the word if the lengths are not the same
			if len(k) != len(v) {
				continue
			}

			// loop through each letter in the current word
			for i := 0; i < len(v); i++ {
				// if known word does not contain letter skip the known workd
				if strings.Index(k, v[i:i+1]) == -1 {
					break
				}

				// if we are on the last letter of the word
				// and they all matched then this is not a valid passphrase
				if i == len(v)-1 {
					return false
				}
			}
		}

		// mark the word as seen
		u[v] = true
	}

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
