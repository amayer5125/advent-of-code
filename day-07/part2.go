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

func containsABA(s string) (r []string) {
	// loop compairing every substring of 3 letters
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			// add found aba to return string array
			r = append(r, s[i:i+3])
		}
	}
	return
}

func main() {
	directions := strings.Split(ReadFile("input.txt"), "\n")

	var ip_set bytes.Buffer
	ssl_ips := 0

	// loop through each ip set
	for i := 0; i < len(directions)-1; i++ {
		chars := strings.Split(directions[i], "")

		var supernet_abas []string
		var hypernet_babs []string
		ip_set.Reset()

		// loop through each charater
		for index, char := range chars {
			// check if we are at the end of a hypernet
			if char == "]" {
				// add any found babs to hypernet babs
				for _, bab := range containsABA(ip_set.String()) {
					hypernet_babs = append(hypernet_babs, bab)
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
				// add any found abas to supernet abas
				for _, aba := range containsABA(ip_set.String()) {
					supernet_abas = append(supernet_abas, aba)
				}
				ip_set.Reset()
				continue
			}

			// add char to current ip_set
			ip_set.WriteString(char)
		}

	SSL_CHECK:
		// loop through each supernet
		for _, sn := range supernet_abas {
			// loop through each hypernet
			for _, hn := range hypernet_babs {
				// check if hypernet and supernet match
				if string([]byte{hn[1], hn[0], hn[1]}) == sn {
					ssl_ips += 1
					break SSL_CHECK
				}
			}
		}
	}
	fmt.Println(ssl_ips)
}
