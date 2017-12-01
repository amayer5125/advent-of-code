package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
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

func (a ByCount) Len() int      { return len(a) }
func (a ByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
	// if the counts are the same use charater code for sorting (alphabetical)
	if a[i].Count == a[j].Count {
		// return []rune(a[i].Name)[0] < []rune(a[j].Name)[0]
		return a[i].Name[0] < a[j].Name[0]
	}
	// sort by count
	return a[i].Count > a[j].Count
}

func main() {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	directions := strings.Split(ReadFile("input.txt"), "\n")

	// secret bunny encryption format
	bunny_encryption, _ := regexp.Compile(`^([a-z\-]+)-(\d+)\[(.+)\]$`)

	for i := 0; i < len(directions)-1; i++ {
		// the parsed direction
		decode := bunny_encryption.FindAllStringSubmatch(directions[i], -1)

		// remove "-" from hash
		encrypted_message := decode[0][1]

		// split encrypted_message into letter slice
		letters := strings.Split(strings.Replace(encrypted_message, "-", "", -1), "")
		var letter_counts []Letter
		// map of found letters
		found := make(map[string]bool)
		// fill letter counts with letters and their count
		for _, letter := range letters {
			if found[letter] {
				continue
			}
			found[letter] = true
			letter_counts = append(letter_counts, Letter{Name: letter, Count: strings.Count(encrypted_message, letter)})
		}

		// sort the slice
		sort.Sort(ByCount(letter_counts))

		match := true
		checksum := strings.Split(decode[0][3], "")
		// check if checksum and letter_counts are in the same order
		for i := 0; i < len(checksum); i++ {
			if checksum[i] != letter_counts[i].Name {
				match = false
				break
			}
		}

		// check if we have a match
		if !match {
			continue
		}

		sector_id, _ := strconv.Atoi(decode[0][2])

		// generate bunny cipher (I am a master cryptographer after all)
		bunny_cipher := make(map[string]string)
		for i := 0; i < len(alphabet); i++ {
			bunny_cipher[alphabet[i]] = alphabet[(i+sector_id)%len(alphabet)]
		}
		// add space to the bunny cipher
		bunny_cipher["-"] = " "

		// buffer to store message
		var message bytes.Buffer

		message_chars := strings.Split(encrypted_message, "")
		// loop through encrypted message
		for _, letter := range message_chars {
			message.WriteString(bunny_cipher[letter])
		}

		// output secter id of northpole object storage
		if message.String() == "northpole object storage" {
			fmt.Println(message.String(), sector_id)
		}
	}
}
