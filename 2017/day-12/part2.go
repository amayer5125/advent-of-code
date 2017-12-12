package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getGroupList(d map[string][]string, index string, tried map[string]bool) (list []string) {
	if tried[index] {
		return
	}

	list = append(list, index)
	tried[index] = true

	for _, v := range d[index] {
		list = append(list, getGroupList(d, v, tried)...)
	}

	return
}

func inSlice(h []string, n string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}

	return false
}

func main() {
	var t int
	var d = make(map[string][]string)
	var found_elements []string

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " <-> ")
		d[tmp[0]] = strings.Split(tmp[1], ", ")
	}

	// loop through each program
	for k, _ := range d {
		// skip programs we have already found/grouped
		if inSlice(found_elements, k) {
			continue
		}

		// add the programs to the found list
		found_elements = append(found_elements, getGroupList(d, k, map[string]bool{})...)

		t++
	}

	fmt.Println(t)
}
