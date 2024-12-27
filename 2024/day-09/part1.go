package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filesystem := make([]string, 0)
	var currentFileId int = 0

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	headers := strings.Split(text, "")
	for i := 0; i < len(headers); i += 2 {
		blocks, _ := strconv.Atoi(headers[i])
		for block := 0; block < blocks; block++ {
			filesystem = append(filesystem, strconv.Itoa(currentFileId))
		}

		if i+1 < len(headers) {
			freeBlocks, _ := strconv.Atoi(headers[i+1])
			for block := 0; block < freeBlocks; block++ {
				filesystem = append(filesystem, ".")
			}
		}

		currentFileId++
	}

	fmt.Println(calculateChecksum(defragFilesystem(filesystem)))
}

func defragFilesystem(filesystem []string) []string {
	newFilesystem := make([]string, len(filesystem))
	copy(newFilesystem, filesystem)
	endPointer := len(filesystem) - 1

	for startPointer := 0; startPointer < endPointer; startPointer++ {
		if newFilesystem[startPointer] != "." {
			continue
		}

		for {
			if endPointer == 0 || newFilesystem[endPointer] != "." {
				break
			}

			endPointer--
		}

		if endPointer < startPointer {
			break
		}

		newFilesystem[endPointer], newFilesystem[startPointer] = newFilesystem[startPointer], newFilesystem[endPointer]
	}

	return newFilesystem
}

func calculateChecksum(filesystem []string) (checksum int) {
	for position, contents := range filesystem {
		if contents == "." {
			return
		}

		checksumBit, _ := strconv.Atoi(contents)
		checksum += position * checksumBit
	}

	return
}
