package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// get the input
	d, _ := strconv.ParseFloat(os.Args[1], 64)

	// distance from the center to the swirl our number is in
	radius := math.Floor(math.Ceil(math.Sqrt(d)) / 2)

	// first number in swirl row
	row_start := math.Pow(2*radius-1, 2) + 1

	// distance from a cardinal direction (n,s,e,w)
	dist_from_cardinal := math.Abs(math.Mod(d-row_start, 2*radius) - (radius - 1))

	fmt.Println(radius + dist_from_cardinal)
}
