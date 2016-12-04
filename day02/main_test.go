package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input1, error := openInput("input.txt")
	if error != nil {
		t.Errorf("unable to open input: %v", error)
		return
	}

	keypad, pos := part1Keyboard()
	keycode := run(keypad, pos, input1)
	const expected = "36629"
	if keycode != expected {
		t.Errorf("invalid keycode '%s' expected '%s'", keycode, expected)
	}
}

func TestPart2(t *testing.T) {
	input1, error := openInput("input.txt")
	if error != nil {
		t.Errorf("unable to open input: %v", error)
		return
	}

	keypad, pos := part2Keyboard()
	keycode := run(keypad, pos, input1)
	const expected = "99C3D"
	if keycode != expected {
		t.Errorf("invalid keycode '%s' expected '%s'", keycode, expected)
	}
}
