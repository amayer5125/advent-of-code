package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var paperRollGrid [][]string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		paperRollGrid = append(paperRollGrid, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(part1(paperRollGrid))
	fmt.Println(part2(paperRollGrid))
}

func part1(paperRollGrid [][]string) (total int) {
	for y := range paperRollGrid {
		for x := range paperRollGrid[y] {
			if paperRollGrid[y][x] != "@" {
				continue
			}

			if countAdjacentRolls(paperRollGrid, x, y) < 4 {
				total++
			}
		}
	}
	return
}

func part2(paperRollGrid [][]string) (total int) {
	var removed bool = true
	for removed {
		removed = false

		for y := range paperRollGrid {
			for x := range paperRollGrid[y] {
				if paperRollGrid[y][x] != "@" {
					continue
				}

				if countAdjacentRolls(paperRollGrid, x, y) < 4 {
					total++
					paperRollGrid[y][x] = "x"
					removed = true
				}
			}
		}
	}
	return
}

func countAdjacentRolls(paperRollGrid [][]string, x, y int) (count int) {
	cordsToCheck := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}
	for _, cords := range cordsToCheck {
		cordsX := x + cords[0]
		if cordsX < 0 || cordsX >= len(paperRollGrid) {
			continue
		}

		cordsY := y + cords[1]
		if cordsY < 0 || cordsY >= len(paperRollGrid) {
			continue
		}

		if paperRollGrid[cordsY][cordsX] == "@" {
			count++
		}
	}
	return
}
