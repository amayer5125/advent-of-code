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
	var password bytes.Buffer
	i := 0

	for {
		hasher.Reset()
		hasher.Write([]byte(input + strconv.Itoa(i)))

		// get our hash string
		hash := hex.EncodeToString(hasher.Sum(nil))

		// check if has starts with 5 zeros
		if strings.HasPrefix(hash, "00000") {
			// put charater in our pasword buffer
			password.WriteByte(hash[5])
		}

		// check if we have an 8 digit password
		if password.Len() == 8 {
			break
		}

		i++
	}
	fmt.Println(password.String())
}
