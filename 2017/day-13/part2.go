package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var pico_second int
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

	for {
		// start at the first pico second (zere pico seconds is a failure ever time)
		pico_second++

		// aren't we being a little over optimistic?
		escaped := true

		// loop through each layer and check if we actually escaped
		for _, v := range layers {
			if (pico_second+v[0])%((v[1]*2)-2) == 0 {
				escaped = false
				break
			}
		}

		if escaped {
			fmt.Println(pico_second)
			break
		}
	}
}
