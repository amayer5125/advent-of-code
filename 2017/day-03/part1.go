package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var d, rl, r, s int

	// get the input
	d, _ = strconv.Atoi(os.Args[1])

	// get the row length (1/4 of swirl square)
	// since distance to center is a square all 4 sides have the same distance values
	rl = int(math.Floor(math.Sqrt(float64(d))))
	if rl < 2 {
		rl = 2
	}

	// how many rows we are out from the center of the swirl
	r = int(math.Ceil(float64(rl) / 2))

	// get the first number in the swirl row
	// start at 2 because first row starts with 2
	s = 2
	for i := 1; i < r; i++ {
		s += i * 8
	}

	// slice of values coresponding to the distance to the center of the square
	dtc := []int{rl - 1}
	// fill the values with the distance to the center
	for j := 0; j < rl-1; j++ {
		adder := -1
		if j >= (rl-1)/2 {
			adder = 1
		}

		dtc = append(dtc, dtc[len(dtc)-1]+adder)
	}

	// get output our starting numbers coresponding slot in the distance to center slice
	fmt.Println(dtc[(d-s)%rl])
}
