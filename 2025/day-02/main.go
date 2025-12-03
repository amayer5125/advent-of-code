package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	var idRanges []Range

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, idRange := range strings.Split(scanner.Text(), ",") {
			ids := strings.Split(idRange, "-")

			start, err := strconv.Atoi(ids[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			end, err := strconv.Atoi(ids[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			idRanges = append(idRanges, Range{start, end})
		}
	}

	fmt.Println(part1(idRanges))
	fmt.Println(part2(idRanges))
}

func part1(idRanges []Range) (total int) {
	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			idString := strconv.Itoa(id)
			if len(idString)%2 != 0 {
				continue
			}

			if idString[:len(idString)/2] == idString[len(idString)/2:] {
				total += id
			}
		}
	}

	return
}

func part2(idRanges []Range) (total int) {
	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			idString := strconv.Itoa(id)

			for length := 1; length <= len(idString)/2; length++ {
				if len(idString)%length != 0 {
					continue
				}

				if idString == strings.Repeat(idString[:length], len(idString)/length) {
					total += id
					break
				}
			}
		}
	}

	return
}
