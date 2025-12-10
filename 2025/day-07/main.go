package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	var tachyonManifold [][]string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tachyonManifold = append(tachyonManifold, strings.Split(scanner.Text(), ""))
	}

	fmt.Println(part1(tachyonManifold))
	fmt.Println(part2(tachyonManifold))
}

func part1(tachyonManifold [][]string) (splits int) {
	var start int = slices.Index(tachyonManifold[0], "S")
	if start == -1 {
		fmt.Fprintln(os.Stderr, "Did not find 'S' in first row")
		os.Exit(1)
	}

	var beams []int = []int{start}
	var beamIndex int

	for row := 1; row < len(tachyonManifold); row++ {
		for column := range tachyonManifold[row] {
			if tachyonManifold[row][column] != "^" {
				continue
			}

			beamIndex = slices.Index(beams, column)
			if beamIndex == -1 {
				continue
			}

			splits++

			beams = append(beams[:beamIndex], beams[beamIndex+1:]...)

			if slices.Index(beams, column-1) == -1 {
				beams = append(beams, column-1)
			}
			if slices.Index(beams, column+1) == -1 {
				beams = append(beams, column+1)
			}
		}
	}

	return
}

func part2(tachyonManifold [][]string) int {
	var start int = slices.Index(tachyonManifold[0], "S")
	if start == -1 {
		fmt.Fprintln(os.Stderr, "Did not find 'S' in first row")
		os.Exit(1)
	}

	return countPaths(tachyonManifold, Point{start, 1}, make(map[Point]int))
}

func countPaths(tachyonManifold [][]string, beam Point, cache map[Point]int) int {
	if beam.Y == len(tachyonManifold)-1 {
		return 1
	}

	if numberOfPaths, ok := cache[beam]; ok {
		return numberOfPaths
	}

	if tachyonManifold[beam.Y][beam.X] == "^" {
		cache[beam] = countPaths(tachyonManifold, Point{beam.X - 1, beam.Y + 1}, cache) + countPaths(tachyonManifold, Point{beam.X + 1, beam.Y + 1}, cache)
		return cache[beam]
	}

	cache[beam] = countPaths(tachyonManifold, Point{beam.X, beam.Y + 1}, cache)
	return cache[beam]
}
