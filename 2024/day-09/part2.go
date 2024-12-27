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

	endPoint := len(newFilesystem) - 1
	startPoint := endPoint
	for {
		if startPoint <= firstEmptySlot(newFilesystem) {
			break
		}

		if newFilesystem[endPoint] == "." {
			endPoint = startPoint - 1
			startPoint = endPoint
			continue
		}

		if newFilesystem[startPoint] == newFilesystem[endPoint] {
			startPoint--
			continue
		}

		startPoint += 1
		newLocation := findFileMoveLocation(newFilesystem[:startPoint], endPoint-startPoint+1)
		if newLocation > -1 {
			for j := newLocation; j < newLocation+endPoint-startPoint+1; j++ {
				newFilesystem[j], newFilesystem[j+startPoint-newLocation] = newFilesystem[j+startPoint-newLocation], newFilesystem[j]
			}
		}

		endPoint = startPoint - 1
		startPoint = endPoint
	}

	return newFilesystem
}

func firstEmptySlot(filesystem []string) int {
	for i := 0; i < len(filesystem); i++ {
		if filesystem[i] == "." {
			return i
		}
	}

	return -1
}

func findFileMoveLocation(filesystem []string, fileLength int) (index int) {
	index = -1

	for i := 0; i < len(filesystem); i++ {
		if filesystem[i] != "." {
			index = -1
			continue
		}

		if index == -1 {
			index = i
		}

		if i-index+1 == fileLength {
			return
		}
	}

	return -1
}

func calculateChecksum(filesystem []string) (checksum int) {
	for position, contents := range filesystem {
		if contents == "." {
			continue
		}

		checksumBit, _ := strconv.Atoi(contents)
		checksum += position * checksumBit
	}

	return
}
