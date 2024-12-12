package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	locationIds := []int{}
	var frequencies = make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		newLocationId, _ := strconv.Atoi(parts[0])
		locationIds = append(locationIds, newLocationId)

		newFrequency, _ := strconv.Atoi(parts[1])
		frequencies[newFrequency] += 1
	}

	distance := 0
	for _, locationId := range locationIds {
		if multiplyer, ok := frequencies[locationId]; ok {
			distance += locationId * multiplyer
		}
	}

	fmt.Println(distance)
}
