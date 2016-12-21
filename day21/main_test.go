package main

import "testing"

func TestRun(t *testing.T) {

	instructions, err := openInput("input.txt")
	if err != nil {
		t.Errorf("unable to open input: %v", err)
		return
	}

	inputPart1 := "abcdefgh"
	actual := run(instructions, inputPart1)

	const expected = "aefgbcdh"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrSwapPos(t *testing.T) {
	i := swapPos{4, 0}
	in := "abcde"

	actual := i.apply(in)
	const expected = "ebcda"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrSwapLetter(t *testing.T) {
	i := swapLetter{l1: 'd', l2: 'b'}
	in := "ebcda"

	actual := i.apply(in)
	const expected = "edcba"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrReverse(t *testing.T) {
	i := reverse{0, 4}
	in := "edcba"

	actual := i.apply(in)
	const expected = "abcde"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrRotateDirection(t *testing.T) {
	i := rotateDirection{"left", 1}
	in := "abcde"

	actual := i.apply(in)
	const expected = "bcdea"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrMove(t *testing.T) {
	i := move{1, 4}
	in := "bcdea"

	actual := i.apply(in)
	const expected = "bdeac"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestInstrRotateBased(t *testing.T) {
	i := rotateBased{'b'}
	in := "abdec"

	actual := i.apply(in)
	const expected = "ecabd"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}
