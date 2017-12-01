package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

func containsABBA(s string) bool {
	// loop compairing every substring of 4 letters
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func main() {
	directions := strings.Split(ReadFile("input.txt"), "\n")

	var ip_set bytes.Buffer
	tls_ips := 0

	// loop through each ip set
	for i := 0; i < len(directions)-1; i++ {
		chars := strings.Split(directions[i], "")

		has_supernet_aba := false
		has_hypernet_abba := false
		ip_set.Reset()

		// loop through each charater
		for index, char := range chars {
			// check if we are at the end of a hypernet
			if char == "]" {
				if containsABBA(ip_set.String()) {
					has_hypernet_abba = true
					break
				}
				ip_set.Reset()
				continue
			}
			// check if we are at the begining of a hypernet or on the last char of the ip
			if char == "[" || index == len(chars)-1 {
				if index == len(chars)-1 {
					// add char to current ip_set
					ip_set.WriteString(char)
				}
				has_supernet_aba = has_supernet_aba || containsABBA(ip_set.String())
				ip_set.Reset()
				continue
			}

			// add char to current ip_set
			ip_set.WriteString(char)
		}

		// add 1 to ips if it matches criteria
		if has_supernet_aba && !has_hypernet_abba {
			tls_ips += 1
		}
	}
	fmt.Println(tls_ips)
}
