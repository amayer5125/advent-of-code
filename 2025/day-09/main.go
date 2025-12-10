package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func (l Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l Line) ContainsPoint(point Point) bool {
	if l.IsVertical() {
		return point.X == l.Start.X && point.Y >= intMin(l.Start.Y, l.End.Y) && point.Y <= intMax(l.Start.Y, l.End.Y)
	}

	return point.Y == l.Start.Y && point.X >= intMin(l.Start.X, l.End.X) && point.X <= intMax(l.Start.X, l.End.X)
}

func main() {
	var redTiles []Point

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		redTiles = append(redTiles, Point{x, y})
	}

	fmt.Println(part1(redTiles))
	fmt.Println(part2(redTiles))
}

func part1(redTiles []Point) (area int) {
	var height int
	var width int
	for firstIndex, firstTile := range redTiles {
		for _, secondTile := range redTiles[firstIndex+1:] {
			height = intAbs(secondTile.Y-firstTile.Y) + 1
			width = intAbs(secondTile.X-firstTile.X) + 1
			if height*width > area {
				area = height * width
			}
		}
	}
	return
}

func part2(redTiles []Point) (area int) {
	var height int
	var width int
	for firstIndex, firstTile := range redTiles {
		for _, secondTile := range redTiles[firstIndex+1:] {
			if !pointInPolygon(Point{firstTile.X, secondTile.Y}, redTiles) {
				continue
			}
			if !pointInPolygon(Point{secondTile.X, firstTile.Y}, redTiles) {
				continue
			}

			height = intAbs(secondTile.Y-firstTile.Y) + 1
			width = intAbs(secondTile.X-firstTile.X) + 1
			if height*width > area {
				area = height * width
			}
		}
	}
	return
}

func pointInPolygon(point Point, polygon []Point) bool {
	var intersections int
	for i := range polygon {
		line := Line{polygon[i], polygon[(i+1)%len(polygon)]}

		if line.ContainsPoint(point) {
			return true
		}

		if !line.IsVertical() {
			continue
		}

		if line.Start.X < point.X {
			continue
		}

		if point.Y < intMin(line.Start.Y, line.End.Y) || point.Y > intMax(line.Start.Y, line.End.Y) {
			continue
		}

		intersections++
	}

	return intersections%2 != 0
}

func intAbs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func intMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func intMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
