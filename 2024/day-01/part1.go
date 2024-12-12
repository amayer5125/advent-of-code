package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		newLeft, _ := strconv.Atoi(parts[0])
		left = append(left, newLeft)

		newRight, _ := strconv.Atoi(parts[1])
		right = append(right, newRight)
	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	distance := 0
	for i := range left {
		if left[i] >= right[i] {
			distance += left[i] - right[i]
		} else {
			distance += right[i] - left[i]
		}
	}

	fmt.Println(distance)
}
