package main

import (
	"fmt"
	"math"

	"os"
	"strings"
	"text/scanner"
)

// // --- Day 1: No Time for a Taxicab ---
// const inputSequence string = "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3, L1, L4, R2, L5, R3, R4, L1, R2, L1, R3, L4, R5, L4, L5, R5, L3, R2, L3, L3, R1, R3, L4, R2, R5, L4, R1, L1, L1, R5, L2, R1, L2, R188, L5, L3, R5, R1, L2, L4, R3, R5, L3, R3, R45, L4, R4, R72, R2, R3, L1, R1, L1, L1, R192, L1, L1, L1, L4, R1, L2, L5, L3, R5, L3, R3, L4, L3, R1, R4, L2, R2, R3, L5, R3, L1, R1, R4, L2, L3, R1, R3, L4, L3, L4, L2, L2, R1, R3, L5, L1, R4, R2, L4, L1, R3, R3, R1, L5, L2, R4, R4, R2, R1, R5, R5, L4, L1, R5, R3, R4, R5, R3, L1, L2, L4, R1, R4, R5, L2, L3, R4, L4, R2, L2, L4, L2, R5, R1, R4, R3, R5, L4, L4, L5, L5, R3, R4, L1, L3, R2, L2, R1, L3, L5, R5, R5, R3, L4, L2, R4, R5, R1, R4, L3"

type input struct {
	direction rune
	distance  int
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

	var result1 int
	var result2 int
	result1 = computeResult(currentState)
	if firstIntersectionFound {
		result2 = computeResult(firstIntersectionState)
	}

	fmt.Println("Day 1 Part 1: ", result1)
	fmt.Println("Day 1 Part 2: ", result2)
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

	//fmt.Println(s, i, sNext)

	return newStates
}

func computeResult(s state) int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}
