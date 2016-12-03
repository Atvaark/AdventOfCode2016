package main

import (
	"bufio"
	"fmt"
	"os"
)

type shape struct {
	sides [3]int
}

func main() {
	file, error := os.Open("input.txt")
	if error != nil {
		panic(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var shapes []shape
	for scanner.Scan() {
		line := scanner.Text()
		var s shape
		fmt.Sscanf(line, "%d %d %d", &s.sides[0], &s.sides[1], &s.sides[2])
		shapes = append(shapes, s)
	}

	part1(shapes)
	part2(shapes)
}

func part1(shapes []shape) {
	count := countTriangles(shapes)
	fmt.Println("Part1: ", count)
}

func part2(shapes []shape) {
	pivoted := pivot(shapes)
	count := countTriangles(pivoted)
	fmt.Println("Part2: ", count)
}

func pivot(shapes []shape) []shape {
	var pivoted []shape
	var tmpShapes [3]shape
	for sIndex, s := range shapes {
		sideIndex := sIndex % 3

		for i := 0; i < 3; i++ {
			if sideIndex == 0 {
				tmpShapes[i] = shape{}
			}

			tmpShapes[i].sides[sideIndex] = s.sides[i]

			if sideIndex == 2 {
				pivoted = append(pivoted, tmpShapes[i])
			}
		}
	}

	return pivoted
}

func countTriangles(shapes []shape) int {
	var count int
	for _, s := range shapes {

		if isTriangle(s) {
			count = count + 1
		}
	}

	return count
}

func isTriangle(s shape) bool {
	if (s.sides[0] < (s.sides[1] + s.sides[2])) && (s.sides[1] < (s.sides[0] + s.sides[2])) && (s.sides[2] < (s.sides[0] + s.sides[1])) {
		return true
	}

	return false
}
