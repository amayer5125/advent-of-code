package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), ",")

		var dirs = make(map[string]float64)
		var m, t, l float64

		for _, v := range d {
			dirs[v]++

			// remove opposite moves
			m = math.Min(dirs["ne"], dirs["sw"])
			dirs["ne"] -= m
			dirs["sw"] -= m

			m = math.Min(dirs["se"], dirs["nw"])
			dirs["se"] -= m
			dirs["nw"] -= m

			m = math.Min(dirs["n"], dirs["s"])
			dirs["n"] -= m
			dirs["s"] -= m

			// convert opposite diagonal moves to polar moves
			m = math.Min(dirs["se"], dirs["sw"])
			dirs["se"] -= m
			dirs["sw"] -= m
			dirs["s"] += m

			m = math.Min(dirs["ne"], dirs["nw"])
			dirs["ne"] -= m
			dirs["nw"] -= m
			dirs["n"] += m

			// convert opposite corner moves
			m = math.Min(dirs["ne"], dirs["s"])
			dirs["ne"] -= m
			dirs["s"] -= m
			dirs["se"] += m

			m = math.Min(dirs["nw"], dirs["s"])
			dirs["nw"] -= m
			dirs["s"] -= m
			dirs["sw"] += m

			m = math.Min(dirs["se"], dirs["n"])
			dirs["se"] -= m
			dirs["n"] -= m
			dirs["se"] += m

			m = math.Min(dirs["sw"], dirs["n"])
			dirs["sw"] -= m
			dirs["n"] -= m
			dirs["nw"] += m

			t = 0
			for _, v := range dirs {
				t += v
			}
			if t > l {
				l = t
			}
		}

		fmt.Println(l)
	}
}
