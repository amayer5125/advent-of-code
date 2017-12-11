package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateList(start, end int) (list []int) {
	for i := start; i <= end; i++ {
		list = append(list, i)
	}

	return
}

func stringToIntSlice(s string) (i []int) {
	// split the string
	t := strings.Split(s, ",")

	// convert each string to integer
	for _, v := range t {
		n, _ := strconv.Atoi(v)
		i = append(i, n)
	}

	return
}

func reverseSection(list []int, start, length int) []int {
	newlist := make([]int, len(list))
	copy(newlist, list)

	// adjust for zero base
	length--

	// reverse the section
	for i := 0; i <= length; i++ {
		s1 := (start + length - i) % len(newlist)
		s2 := (start + i) % len(newlist)

		newlist[s1] = list[s2]
	}

	return newlist
}

func main() {
	var c int
	var l []int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l = stringToIntSlice(scanner.Text())
	}

	list := generateList(0, 255)

	// loop through each skip length
	for i := 0; i < len(l); i++ {
		list = reverseSection(list, c, l[i])

		c += l[i] + i
	}

	fmt.Println(list[0] * list[1])
}
