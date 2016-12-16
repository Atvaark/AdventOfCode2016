package main

import "testing"

func TestStep(t *testing.T) {
	actual := step("111100001010")
	const expected = "1111000010100101011110000"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestGetChecksumPart1(t *testing.T) {
	const length = 272
	const input = "10010000000110000"

	actual := getChecksum(input, length)

	const expected = "10010110010011110"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestGetChecksumPart2(t *testing.T) {
	const length = 35651584
	const input = "10010000000110000"

	actual := getChecksum(input, length)

	const expected = "01101011101100011"
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}
