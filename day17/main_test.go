package main

import "testing"

func TestRunPart1(t *testing.T) {
	const passcode = "gdjjyniy"
	actualPart1, actualPart2 := run(passcode)
	const expecterPart1 = "DUDDRLRRRD"
	if actualPart1 != expecterPart1 {
		t.Errorf("invalid result. got '%v' expected '%v'", actualPart1, expecterPart1)
	}

	const expecterPart2 = 578
	if actualPart2 != expecterPart2 {
		t.Errorf("invalid result. got '%v' expected '%v'", actualPart1, expecterPart1)
	}
}

func TestRun1(t *testing.T) {
	const passcode = "ihgpwlah"
	actual, _ := run(passcode)
	const expected = "DDRRRD"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRun2(t *testing.T) {
	const passcode = "kglvqrro"
	actual, _ := run(passcode)
	const expected = "DDUDRLRRUDRD"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRun3(t *testing.T) {
	const passcode = "ulqzkmiv"
	actual, _ := run(passcode)
	const expected = "DRURDRUDDLLDLUURRDULRLDUUDDDRR"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestCalcDoorState(t *testing.T) {
	const passcode = "hijkl"
	const path = ""
	pos := position{}
	actual := calcDoorState(passcode, path, pos)

	expected := doorState{upOpen: false, downOpen: true, leftOpen: false, rightOpen: false}
	if actual != expected {
		t.Errorf("invalid result. got '%+v' expected '%+v'", actual, expected)
	}
}
