package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, error := openInput("input.txt")
	if error != nil {
		fmt.Printf("unable to read input: %v", error)
		os.Exit(1)
	}

	part1Result, part2Result := run(lines)

	fmt.Println("max: ", part1Result)
	fmt.Println("min: ", part2Result)
}

func openInput(name string) ([]string, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func run(messages []string) (part1Result string, part2Result string) {
	m := make(map[int]map[rune]int)
	//messageCount := 0
	for _, message := range messages {
		for msgIdx, msgRune := range message {
			m1 := m[msgIdx]
			if m1 == nil {
				m1 = make(map[rune]int)
				m[msgIdx] = m1
			}

			m[msgIdx][msgRune] = m[msgIdx][msgRune] + 1
			//messageCount++
		}
	}

	msgIdxMax := -1
	for i := range m {
		if i > msgIdxMax {
			msgIdxMax = i
		}
	}

	maxResult := make([]rune, msgIdxMax+1)
	minResult := make([]rune, msgIdxMax+1)

	for i, m2 := range m {
		rMax := '?'
		rMaxCount := -1

		rMin := '?'
		rMinCount := len(messages) + 1

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

	return maxResultString, minResultString
}
