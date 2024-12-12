package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var total int
	rules := make(map[string][]string)
	phaseTwo := false
	var goodOrder bool

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			phaseTwo = true
			continue
		}

		if !phaseTwo {
			rule := strings.Split(text, "|")

			if _, ok := rules[rule[1]]; !ok {
				rules[rule[1]] = make([]string, 0)
			}
			rules[rule[1]] = append(rules[rule[1]], rule[0])
			continue
		}

		pages := strings.Split(text, ",")
		for pageIndex, page := range pages {
			goodOrder = true

			for _, rule := range rules[page] {
				if slices.Index(pages, rule) > pageIndex {
					goodOrder = false
					break
				}
			}

			if !goodOrder {
				break
			}
		}

		if goodOrder {
			middlePage, _ := strconv.Atoi(pages[len(pages)/2])
			total += middlePage
		}
	}

	fmt.Println(total)
}
