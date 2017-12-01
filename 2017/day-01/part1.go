package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var d string
	var nums []int
	var total int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d = scanner.Text()
	}

	// put the input into an integer slice
	for i := 0; i < len(d); i++ {
		c, _ := strconv.Atoi(d[i : i+1])
		nums = append(nums, c)
	}

	// loop through each integer
	for k, v := range nums {
		// get the next index
		n := (k + 1) % len(d)

		// check if the value is equal to the next index
		if v == nums[n] {
			total += v
		}
	}

	fmt.Println(total)
}
