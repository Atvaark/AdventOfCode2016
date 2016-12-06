package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, error := os.Open("input.txt")
	_ = error

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make(map[int]map[rune]int)
	messageCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i, r := range line {
			m1 := m[i]
			if m1 == nil {
				m1 = make(map[rune]int)
				m[i] = m1
			}

			m[i][r] = m[i][r] + 1
			messageCount++
		}
	}

	iMax := -1
	for i := range m {
		if i > iMax {
			iMax = i
		}
	}

	maxResult := make([]rune, iMax+1)
	minResult := make([]rune, iMax+1)

	for i, m2 := range m {

		rMax := '?'
		rMaxCount := -1

		rMin := '?'
		rMinCount := messageCount + 1

		for r, c := range m2 {
			if c > rMaxCount {
				rMaxCount = c
				rMax = r
			}

			if c < rMinCount {
				rMinCount = c
				rMin = r
			}
		}

		maxResult[i] = rMax
		minResult[i] = rMin
	}

	maxResultString := string(maxResult[:])
	minResultString := string(minResult[:])

	fmt.Println("max: ", maxResultString)
	fmt.Println("min: ", minResultString)
}
