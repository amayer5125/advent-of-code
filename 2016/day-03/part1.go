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

	for i := 0; i < len(directions)-1; i++ {
		sides := strings.Fields(directions[i])

		side1, _ := strconv.Atoi(sides[0])
		side2, _ := strconv.Atoi(sides[1])
		side3, _ := strconv.Atoi(sides[2])

		if IsValid(side1, side2, side3) {
			valid_triangles++
		}
	}

	fmt.Println(valid_triangles)
}
