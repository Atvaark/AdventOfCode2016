package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type pos struct {
	x int
	y int
}

func main() {
	input1, error := openInput("input.txt")
	if error != nil {
		fmt.Printf("unable to open input: %v", error)
		os.Exit(1)
	}

	part1Keypad, part1Pos := part1Keyboard()
	keycode1 := run(part1Keypad, part1Pos, input1)
	println("Part1: ", keycode1)

	part2Keypad, part2Pos := part2Keyboard()
	keycode2 := run(part2Keypad, part2Pos, input1)
	println("Part2: ", keycode2)
}

func openInput(name string) (string, error) {
	data, error := ioutil.ReadFile(name)
	if error != nil {
		return "", error
	}

	return string(data), nil
}

func part1Keyboard() ([][]rune, pos) {
	keypad := [][]rune{}
	row1 := []rune{'1', '2', '3'}
	row2 := []rune{'4', '5', '6'}
	row3 := []rune{'7', '8', '9'}
	keypad = append(keypad, row1)
	keypad = append(keypad, row2)
	keypad = append(keypad, row3)
	return keypad, pos{1, 1}
}

func part2Keyboard() ([][]rune, pos) {
	keypad := [][]rune{}
	row1 := []rune{' ', ' ', '1', ' ', ' '}
	row2 := []rune{' ', '2', '3', '4', ' '}
	row3 := []rune{'5', '6', '7', '8', '9'}
	row4 := []rune{' ', 'A', 'B', 'C', ' '}
	row5 := []rune{' ', ' ', 'D', ' ', ' '}
	keypad = append(keypad, row1)
	keypad = append(keypad, row2)
	keypad = append(keypad, row3)
	keypad = append(keypad, row4)
	keypad = append(keypad, row5)
	return keypad, pos{0, 2}
}

func run(keypad [][]rune, initial pos, in string) string {
	maxRow := len(keypad) - 1
	maxColumn := len(keypad[0]) - 1

	var keycode []rune
	p := initial
	for _, c := range in {
		pNext := p

		switch c {
		case 'U':
			if p.y > 0 {
				pNext.y = p.y - 1
			}
		case 'R':
			if p.x < maxColumn {
				pNext.x = p.x + 1
			}
		case 'D':
			if p.y < maxRow {
				pNext.y = p.y + 1
			}
		case 'L':
			if p.x > 0 {
				pNext.x = p.x - 1
			}
		case '\n':
			keycode = append(keycode, keypad[pNext.y][pNext.x])
		default:
			continue
		}

		if keypad[pNext.y][pNext.x] == ' ' {
			continue
		}

		p = pNext
	}

	return string(keycode)
}
