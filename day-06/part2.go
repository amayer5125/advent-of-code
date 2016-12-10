package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}

	return string(data)
}

type Letter struct {
	Name  string
	Count int
}

type ByCount []Letter

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func main() {
	directions := strings.Split(ReadFile("input.txt"), "\n")

	var col_letters bytes.Buffer
	var answer bytes.Buffer

	for x := 0; x < len(directions[0]); x++ {
		// reset col letters buffer for new col
		col_letters.Reset()

		// write entire col to col_letters buffer
		for y := 0; y < len(directions)-1; y++ {
			col_letters.WriteByte(directions[y][x])
		}

		var letter_counts []Letter
		// map of found letters
		found := make(map[string]bool)
		for _, letter := range strings.Split(col_letters.String(), "") {
			if found[letter] {
				continue
			}
			found[letter] = true
			letter_counts = append(letter_counts, Letter{Name: letter, Count: strings.Count(col_letters.String(), letter)})
		}

		// sort the slice
		sort.Sort(ByCount(letter_counts))
		// add most frequent letter to answer
		answer.WriteString(letter_counts[len(letter_counts)-1].Name)
	}
	fmt.Println(answer.String())
}
