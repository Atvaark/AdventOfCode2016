package main

import (
	"fmt"
)

const trapSymbol = '^'
const safeSymbol = '.'

func main() {
	const inputRow = ".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^"
	const part1RowCount = 40
	const part2RowCount = 400000

	safeCountPart1 := run(inputRow, part1RowCount)
	fmt.Println("part 1: ", safeCountPart1)

	safeCountPart2 := run(inputRow, part2RowCount)
	fmt.Println("part 2: ", safeCountPart2)
}

func run(inputRow string, rowCount int) int {
	var safeCount int

	for _, symbol := range inputRow {
		if symbol == safeSymbol {
			safeCount++
		}
	}

	previousRow := inputRow
	rowLen := len(inputRow)
	nextRow := make([]rune, rowLen, rowLen)
	for i := 1; i < rowCount; i++ {
		for j := 0; j < rowLen; j++ {
			var left rune
			var center rune
			var right rune
			if j == 0 {
				left = safeSymbol
			} else {
				left = rune(previousRow[j-1])
			}

			center = rune(previousRow[j])

			if j == rowLen-1 {
				right = safeSymbol
			} else {
				right = rune(previousRow[j+1])
			}

			var isTrap bool
			switch {
			case left == trapSymbol && center == trapSymbol && right == safeSymbol:
				isTrap = true
			case left == safeSymbol && center == trapSymbol && right == trapSymbol:
				isTrap = true
			case left == trapSymbol && center == safeSymbol && right == safeSymbol:
				isTrap = true
			case left == safeSymbol && center == safeSymbol && right == trapSymbol:
				isTrap = true
			}

			var nextRune rune
			if isTrap {
				nextRune = trapSymbol
			} else {
				nextRune = safeSymbol
			}

			if !isTrap {
				safeCount++
			}

			nextRow[j] = nextRune
		}

		previousRow = string(nextRow)
	}

	return safeCount
}
