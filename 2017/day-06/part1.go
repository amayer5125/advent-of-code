package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLargestBlock(d []int) (i int) {
	for k, v := range d {
		if v > d[i] {
			i = k
		}
	}

	return
}

func distributeBlock(d []int, b int) (n []int) {
	n = d

	// get the value of the largest block
	dist := n[b]
	// set the value of the largest block to zero
	n[b] = 0

	// distribute the block
	for i := 1; i <= dist; i++ {
		n[(b+i)%len(d)]++
	}

	return
}

func sliceToString(d []int) (s string) {
	for i := 0; i < len(d); i++ {
		s += strconv.Itoa(d[i])
	}

	return
}

func main() {
	var t, l int
	var d []int
	var seen = make(map[string]bool)

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// split the input by tabs
		tmp := strings.Split(scanner.Text(), "\t")
		// add each item to d
		for _, v := range tmp {
			n, _ := strconv.Atoi(v)
			d = append(d, n)
		}
	}

	for {
		// count the number of moves we need to make
		t++

		// get the largest block
		l = getLargestBlock(d)

		// distribute the block
		d = distributeBlock(d, l)

		// get the slice signature
		s := sliceToString(d)

		// check if we have seen this signature before
		if seen[s] {
			break
		}
		// add signature to seen slice
		seen[s] = true
	}

	fmt.Println(t)
}
