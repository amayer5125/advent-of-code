package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func IsValid(s1, s2, s3 int) bool {
	return s1+s2 > s3 && s2+s3 > s1 && s3+s1 > s2
}

func main() {
	valid_triangles := 0

	directions := strings.Split(ReadFile("input.txt"), "\n")

	for i := 0; i < len(directions)-1; i += 3 {
		row1 := strings.Fields(directions[i])
		row2 := strings.Fields(directions[i+1])
		row3 := strings.Fields(directions[i+2])

		for j := 0; j < 3; j++ {
			side1, _ := strconv.Atoi(row1[j])
			side2, _ := strconv.Atoi(row2[j])
			side3, _ := strconv.Atoi(row3[j])

			if IsValid(side1, side2, side3) {
				valid_triangles++
			}
		}
	}

	fmt.Println(valid_triangles)
}
