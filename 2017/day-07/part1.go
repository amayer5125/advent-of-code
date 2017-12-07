package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var program_description = regexp.MustCompile(`([a-z]+) \((\d+)\)(?: -> (.*))?`)

func main() {
	var parents []string
	var isChild = make(map[string]bool)

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := program_description.FindStringSubmatch(scanner.Text())

		// ignore programs that do not have children
		if len(parts[3]) == 0 {
			continue
		}

		parents = append(parents, parts[1])

		children := strings.Split(parts[3], ", ")
		for _, i := range children {
			isChild[i] = true
		}
	}

	for _, p := range parents {
		if isChild[p] {
			continue
		}

		fmt.Println(p)
	}
}
