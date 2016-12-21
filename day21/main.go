package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type instruction interface {
	apply(input string) string
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
}

func run(instructions []instruction, input string) string {
	c := input
	for _, inst := range instructions {
		c = inst.apply(c)
	}

	return c
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
		i := parseInstruction(line)
		instructions = append(instructions, i)
	}

	return instructions, nil
}

type swapPos struct {
	p1 int
	p2 int
}

func (i swapPos) apply(input string) string {
	work := []rune(input)
	tmp := work[i.p1]
	work[i.p1] = work[i.p2]
	work[i.p2] = tmp
	return string(work)
}

type swapLetter struct {
	l1 rune
	l2 rune
}

func (i swapLetter) apply(input string) string {
	work := []rune(input)
	var p1, p2 int

	for j, r := range work {
		if r == i.l1 {
			p1 = j
			continue
		}

		if r == i.l2 {
			p2 = j
			continue
		}
	}

	tmp := work[p1]
	work[p1] = work[p2]
	work[p2] = tmp
	return string(work)
}

type rotateBased struct {
	l rune
}

func (i rotateBased) apply(input string) string {
	work := []rune(input)
	var p int
	for j, r := range work {
		if r == i.l {
			p = j
			break
		}
	}

	rotationCount := 1
	rotationCount += p
	if p >= 4 {
		rotationCount++
	}

	tmp := rotateDirection{dir: "right", s: rotationCount}
	return tmp.apply(input)
}

type rotateDirection struct {
	dir string
	s   int
}

func (i rotateDirection) apply(input string) string {
	work := []rune(input)
	by := i.s % len(work)

	var left, right []rune
	switch i.dir {
	case "right":
		// abcd 1
		// d abc
		left = work[len(work)-by:]
		right = work[:len(work)-by]
	case "left":
		// abcd 1
		// bcda
		left = work[by:]
		right = work[:by]
	}

	out := append(left, right...)
	return string(out)
}

type reverse struct {
	p1 int
	p2 int
}

func (i reverse) apply(input string) string {
	work := []rune(input)
	before := work[:i.p1]
	between := work[i.p1 : i.p2+1]
	after := work[i.p2+1:]
	betweenReversed := make([]rune, len(between))
	for i := 0; i < len(between); i++ {
		betweenReversed[i] = between[len(between)-i-1]
	}

	sum := append(before, betweenReversed...)
	sum = append(sum, after...)
	return string(sum)
}

type move struct {
	p1 int
	p2 int
}

func (i move) apply(input string) string {
	work := []rune(input)

	tmp := work[i.p1]
	workTmp := make([]rune, 0, len(work))

	work = append(work[:i.p1], work[i.p1+1:]...)
	work1 := work[:i.p2]
	work2 := work[i.p2:]
	workTmp = append(workTmp, work1...)
	workTmp = append(workTmp, tmp)
	workTmp = append(workTmp, work2...)
	return string(workTmp)

	return input
}

func parseInstruction(line string) instruction {
	split := strings.Split(line, " ")
	switch split[0] {
	case "swap":
		switch split[1] {
		case "position":
			// swap position 4 with position 0
			var p1, p2 int
			fmt.Sscanf(split[2], "%d", &p1)
			fmt.Sscanf(split[5], "%d", &p2)

			// fmt.Printf("swap p %d with %d\n", p1, p2)
			return swapPos{p1, p2}
		case "letter":
			// swap letter d with letter b
			var l1, l2 rune
			l1, _ = utf8.DecodeRuneInString(split[2])
			l2, _ = utf8.DecodeRuneInString(split[5])
			// fmt.Printf("swap letter %c with %c\n", l1, l2)
			return swapLetter{l1, l2}
		}
	case "rotate":
		switch split[1] {
		case "based":
			// rotate based on position of letter b
			// rotate based on position of letter d
			var l rune
			l, _ = utf8.DecodeRuneInString(split[6])
			// fmt.Printf("rotate based on %c\n", l)
			return rotateBased{l}
		default:
			// rotate left 1 step
			d := split[1]
			var p int
			fmt.Sscanf(split[2], "%d", &p)

			// fmt.Printf("rotate %s %d\n", d, p)
			return rotateDirection{d, p}
		}

	case "reverse":
		// reverse positions 0 through 4
		var p1, p2 int
		fmt.Sscanf(split[2], "%d", &p1)
		fmt.Sscanf(split[4], "%d", &p2)

		// fmt.Printf("reverse %d through %d\n", p1, p2)
		return reverse{p1, p2}
	case "move":
		// move position 1 to position 4
		// move position 3 to position 0
		var p1, p2 int
		fmt.Sscanf(split[2], "%d", &p1)
		fmt.Sscanf(split[5], "%d", &p2)
		// fmt.Printf("move %d to %d\n", p1, p2)
		return move{p1, p2}
	}

	return nil
}
