package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction interface {
	apply(input string) string
	reverseApply(input string) string
}

func main() {
	instructions, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	inputPart1 := "abcdefgh"
	resultPart1 := run(instructions, inputPart1)
	fmt.Println("result part 1: ", resultPart1)

	inputPart2 := "fbgdceah"
	resultPart2 := runReverse(instructions, inputPart2)
	fmt.Println("result part 2: ", resultPart2)
}

func run(instructions []instruction, input string) string {
	password := input
	for _, inst := range instructions {
		password = inst.apply(password)
	}
	return password
}

func runReverse(instructions []instruction, input string) string {
	password := input
	for i := len(instructions) - 1; i >= 0; i-- {
		inst := instructions[i]
		password = inst.reverseApply(password)
	}
	return password
}

func openInput(name string) ([]instruction, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var instructions []instruction
	for scanner.Scan() {
		line := scanner.Text()
		instr := parseInstruction(line)
		instructions = append(instructions, instr)
	}

	return instructions, nil
}

type swapPos struct {
	x int
	y int
}

func (i swapPos) apply(input string) string {
	work := []rune(input)
	tmp := work[i.x]
	work[i.x] = work[i.y]
	work[i.y] = tmp
	return string(work)
}

func (i swapPos) reverseApply(input string) string {
	tmp := i.x
	i.x = i.y
	i.y = tmp
	return i.apply(input)
}

type swapLetter struct {
	x rune
	y rune
}

func (i swapLetter) apply(input string) string {
	work := []rune(input)
	var x, y int

	for j, r := range work {
		if r == i.x {
			x = j
			continue
		}

		if r == i.y {
			y = j
			continue
		}
	}

	tmp := work[x]
	work[x] = work[y]
	work[y] = tmp
	return string(work)
}

func (i swapLetter) reverseApply(input string) string {
	tmp := i.x
	i.x = i.y
	i.y = tmp
	return i.apply(input)
}

type rotateBased struct {
	x rune
}

func (i rotateBased) apply(input string) string {
	work := []rune(input)
	var xPos int
	for j, r := range work {
		if r == i.x {
			xPos = j
			break
		}
	}

	rotationCount := 1
	rotationCount += xPos
	if xPos >= 4 {
		rotationCount++
	}

	tmp := rotateDirection{direction: "right", x: rotationCount}
	return tmp.apply(input)
}

func (i rotateBased) reverseApply(input string) string {
	for j := len(input) - 1; j >= 0; j-- {
		reverseInstr := rotateDirection{direction: "left", x: j}
		reversed := reverseInstr.apply(input)
		reversedReversed := i.apply(reversed)
		if input != reversed && input == reversedReversed {
			return reversed
		}
	}

	return input
}

type rotateDirection struct {
	direction string
	x         int
}

func (i rotateDirection) apply(input string) string {
	work := []rune(input)
	rotationCount := i.x % len(work)

	var left, right []rune
	switch i.direction {
	case "right":
		left = work[len(work)-rotationCount:]
		right = work[:len(work)-rotationCount]
	case "left":
		left = work[rotationCount:]
		right = work[:rotationCount]
	}

	result := append(left, right...)
	return string(result)
}

func (i rotateDirection) reverseApply(input string) string {
	switch i.direction {
	case "right":
		i.direction = "left"
	case "left":
		i.direction = "right"
	}

	return i.apply(input)
}

type reverse struct {
	x int
	y int
}

func (i reverse) apply(input string) string {
	work := []rune(input)
	left := work[:i.x]
	middle := work[i.x : i.y+1]
	right := work[i.y+1:]
	middleReversed := make([]rune, len(middle))
	for i := 0; i < len(middle); i++ {
		middleReversed[i] = middle[len(middle)-i-1]
	}

	result := append(left, middleReversed...)
	result = append(result, right...)
	return string(result)
}

func (i reverse) reverseApply(input string) string {
	return i.apply(input)
}

type move struct {
	x int
	y int
}

func (i move) apply(input string) string {
	work := []rune(input)

	middle := work[i.x]
	work = append(work[:i.x], work[i.x+1:]...)
	left := work[:i.y]
	right := work[i.y:]

	result := make([]rune, 0, len(work))
	result = append(result, left...)
	result = append(result, middle)
	result = append(result, right...)
	return string(result)
}

func (i move) reverseApply(input string) string {
	tmp := i.x
	i.x = i.y
	i.y = tmp
	return i.apply(input)
}

func parseInstruction(line string) instruction {
	split := strings.Split(line, " ")
	switch split[0] {
	case "swap":
		switch split[1] {
		case "position":
			var x, y int
			fmt.Sscanf(split[2], "%d", &x)
			fmt.Sscanf(split[5], "%d", &y)
			return swapPos{x, y}
		case "letter":
			var x, y rune
			x = []rune(split[2])[0]
			y = []rune(split[5])[0]
			return swapLetter{x, y}
		}
	case "rotate":
		switch split[1] {
		case "based":
			var x rune
			x = []rune(split[6])[0]
			return rotateBased{x}
		default:
			x := split[1]
			var p int
			fmt.Sscanf(split[2], "%d", &p)
			return rotateDirection{x, p}
		}

	case "reverse":
		var x, y int
		fmt.Sscanf(split[2], "%d", &x)
		fmt.Sscanf(split[4], "%d", &y)
		return reverse{x, y}
	case "move":
		var x, y int
		fmt.Sscanf(split[2], "%d", &x)
		fmt.Sscanf(split[5], "%d", &y)
		return move{x, y}
	}

	return nil
}
