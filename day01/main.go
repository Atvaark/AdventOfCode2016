// --- Day 1: No Time for a Taxicab ---
package main

import (
	"fmt"
	"math"

	"os"
	"strings"
	"text/scanner"
)

type input struct {
	direction rune
	distance  int
}

type output struct {
	distanceEnd               int
	distanceFirstIntersection int
}

type state struct {
	x         int
	y         int
	direction byte
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no sequence specified")
		return
	}

	inputSequence := os.Args[1]
	result := run(inputSequence)

	fmt.Println("Day 1 Part 1: ", result.distanceEnd)
	fmt.Println("Day 1 Part 2: ", result.distanceFirstIntersection)
}

func run(inputSequence string) output {
	var instructions []input
	instructions = parseInstructions(inputSequence)

	var currentState state
	var nextStates []state

	var firstIntersectionState state
	var firstIntersectionFound bool

	var visitedStates []state
	visitedStates = append(visitedStates, currentState)

	for _, instruction := range instructions {
		nextStates = applyInput(currentState, instruction)
		for _, nextState := range nextStates {
			if !firstIntersectionFound {
				if containsIntersection(visitedStates, nextState) {
					firstIntersectionFound = true
					firstIntersectionState = nextState
				}

				visitedStates = append(visitedStates, nextState)
			}

			currentState = nextState
		}
	}

	var result output
	result.distanceEnd = computeResult(currentState)
	if firstIntersectionFound {
		result.distanceFirstIntersection = computeResult(firstIntersectionState)
	}
	return result
}

func parseInstructions(inputSequence string) []input {
	var instructions []input
	inputReader := strings.NewReader(inputSequence)
	var inputScanner scanner.Scanner
	inputScanner.Init(inputReader)

	var token rune
	for token != scanner.EOF {
		token = inputScanner.Scan()
		if token == ',' || token == ' ' {
			continue
		}

		var tokenText string
		tokenText = inputScanner.TokenText()

		instruction := input{direction: 'R', distance: 2}
		tokenLength := len(tokenText)

		if tokenLength <= 1 {
			continue
		}

		instruction.direction = rune(tokenText[0])
		fmt.Sscanf(tokenText[1:], "%d", &instruction.distance)

		instructions = append(instructions, instruction)
	}

	return instructions
}

func containsIntersection(states []state, findState state) bool {
	for _, currentState := range states {
		if currentState.x == findState.x && currentState.y == findState.y {
			return true
		}

	}
	return false
}

func applyInput(s state, i input) []state {
	var newStates []state

	// 0 = N, 1 = E, 2 = S, 3 = W
	var sNext state
	switch i.direction {
	case 'R':
		sNext.direction = (4 + s.direction + 1) % 4
	case 'L':
		sNext.direction = (4 + s.direction - 1) % 4
	}

	var xCoef int
	var yCoef int
	switch sNext.direction {
	case 0:
		xCoef = 0
		yCoef = 1
	case 1:
		xCoef = 1
		yCoef = 0
	case 2:
		xCoef = 0
		yCoef = -1
	case 3:
		xCoef = -1
		yCoef = 0
	}

	sNext.x = s.x
	sNext.y = s.y

	for index := 0; index < i.distance; index++ {
		sNext.x = sNext.x + xCoef
		sNext.y = sNext.y + yCoef
		newStates = append(newStates, sNext)
	}

	return newStates
}

func computeResult(s state) int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}
