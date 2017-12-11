package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateList(start, end int) (list []byte) {
	for i := start; i <= end; i++ {
		list = append(list, byte(i))
	}

	return
}

func reverseSection(list []byte, start, length int) []byte {
	newlist := make([]byte, len(list))
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

func convertToDenseHash(list []byte) []byte {
	newlist := make([]byte, 16)

	for i := 0; i < 16; i++ {
		start := i * 16
		end := start + 16

		t := 0
		for _, v := range list[start:end] {
			t = t ^ int(v)
		}

		newlist[i] = byte(t)
	}

	return newlist
}

func convertToHex(list []byte) (s string) {
	for _, v := range list {
		n := strconv.FormatInt(int64(v), 16)

		// pad with zero
		if len(n) < 2 {
			s += "0"
		}

		s += n
	}

	return
}

func main() {
	var current_pos, skip int
	var lengths []byte
	var suffix = []byte{17, 31, 73, 47, 23}

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lengths = scanner.Bytes()
	}
	// add the suffix to the lengths
	lengths = append(lengths, suffix...)

	// generate the list
	list := generateList(0, 255)

	// loop through each round
	for r := 0; r < 64; r++ {
		// loop through each skip length
		for i := 0; i < len(lengths); i++ {
			l := int(lengths[i])

			list = reverseSection(list, current_pos, l)

			current_pos += (l + skip) % len(list)
			skip++
		}
	}

	// get dense hash
	dense_hash := convertToDenseHash(list)

	// get hex representation
	hash := convertToHex(dense_hash)

	fmt.Println(hash)
}
