package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return []byte{}
	}

	return bytes.TrimSpace(data)
}

// compile expander patter
var expander_patt = regexp.MustCompile(`\((\d+)x(\d+)\)`)

// accept compressed byte array and return expanded byte array
func decompress(cb []byte, recursion chan []byte) {
	// get position of next expander
	expander := expander_patt.FindIndex(cb)
	// return the sent bytes  (nothing to expand)
	if expander == nil {
		recursion <- cb
		return
	}
	rerecursion := make(chan []byte)
	// return bytes up to expander and expand the rest
	if expander[0] > 0 {
		go decompress(cb[expander[0]:], rerecursion)
		recursion <- append(cb[0:expander[0]], <-rerecursion...)
		return
	}

	// get expander parts
	expander_parts := expander_patt.FindAllSubmatch(cb[expander[0]:expander[1]], -1)
	// remove expander from compressed bytes
	cb = cb[expander[1]:]
	// number of chars and how to repeat them
	byte_len, _ := strconv.Atoi(string(expander_parts[0][1]))
	repeat, _ := strconv.Atoi(string(expander_parts[0][2]))

	// expand to_repeat and concat the rest of the decrypted message
	go decompress(cb[byte_len:], rerecursion)
	recursion <- append(bytes.Repeat(cb[:byte_len], repeat), <-rerecursion...)
	return
}

func main() {
	recursion := make(chan []byte)
	go decompress(ReadFile("input.txt"), recursion)
	fmt.Println(len(<-recursion))
}
