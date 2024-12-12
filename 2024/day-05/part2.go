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
	var checkItAgain bool
	var orderUpdated bool

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
		orderUpdated = false
		for {
			checkItAgain = false

			for pageIndex, page := range pages {
				for _, rule := range rules[page] {
					ruleIndex := slices.Index(pages, rule)
					if ruleIndex > pageIndex {
						pages = slices.Delete(pages, ruleIndex, ruleIndex+1)
						pages = slices.Insert(pages, pageIndex, rule)

						orderUpdated = true
						checkItAgain = true
					}
				}
			}

			if !checkItAgain {
				break
			}
		}

		if orderUpdated {
			middlePage, _ := strconv.Atoi(pages[len(pages)/2])
			total += middlePage
		}
	}

	fmt.Println(total)
}
