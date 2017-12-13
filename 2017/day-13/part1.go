package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t int
	var layers [][2]int

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ": ")

		var new_layer [2]int
		new_layer[0], _ = strconv.Atoi(tmp[0])
		new_layer[1], _ = strconv.Atoi(tmp[1])

		layers = append(layers, new_layer)
	}

	// loop through each security layer
	for _, v := range layers {
		// check if the scanner will be at the top of the layer when we get there
		if v[0]%((v[1]*2)-2) == 0 {
			t += v[0] * v[1]
		}
	}

	fmt.Println(t)
}
