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
	shapes, error := openShapeFile("input.txt")
	if error != nil {
		fmt.Printf("unable to open input: %v", error)
		os.Exit(1)
	}

	part1Count := part1(shapes)
	fmt.Println("Part1: ", part1Count)
	part2Count := part2(shapes)
	fmt.Println("Part2: ", part2Count)
}

func openShapeFile(name string) ([]shape, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
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

	return shapes, nil
}

func part1(shapes []shape) int {
	return countTriangles(shapes)
}

func part2(shapes []shape) int {
	pivoted := pivot(shapes)
	return countTriangles(pivoted)
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
