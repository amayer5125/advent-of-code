package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "cxdnnyjw"
	hasher := md5.New()
	var password [8]byte
	i := 0

	for {
		hasher.Reset()
		hasher.Write([]byte(input + strconv.Itoa(i)))
		// get our hash string
		hash := hex.EncodeToString(hasher.Sum(nil))

		// check if has starts with 5 zeros and has 0-7 in the index slot
		if strings.HasPrefix(hash, "00000") && strings.ContainsAny(string(hash[5]), "01234567") {
			slot, _ := strconv.Atoi(string(hash[5]))
			// check if the slot trying to be writen to is empty
			if bytes.Equal([]byte{password[slot]}, []byte{0}) {
				password[slot] = hash[6]
				// passphrase output for extra credit
				passphrase := bytes.Replace(password[:], []byte{0}, []byte("-"), -1)
				fmt.Printf("|%s|\n", passphrase[:])
			}
		}

		// we are dont if there are no zeros bytes (null) left in our password
		if !bytes.Contains(password[:], []byte{0}) {
			break
		}

		i++
	}
}
