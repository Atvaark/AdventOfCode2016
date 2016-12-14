package main

import "testing"

func TestGetTripletRune(t *testing.T) {
	input := "ae2e85dd75d63e916a525df95e999ea0"
	actual, ok := getTripletRune(input)
	const expectedRune = '9'
	const expectedOk = true
	if actual != expectedRune {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expectedRune)
	}

	if ok != expectedOk {
		t.Errorf("invalid result. got '%v' expected '%v'", ok, expectedOk)
	}
}

func TestGetStretchedHash(t *testing.T) {
	knownHashes := make(map[string]string, 0)
	dataString := "abc0"
	actualHash := getStretchedHash(knownHashes, dataString)

	const expectedHash = "a107ff634856bb300138cac6568c0f24"
	if actualHash != expectedHash {
		t.Errorf("invalid result. got '%v' expected '%v'", actualHash, expectedHash)
	}
}

func TestRun(t *testing.T) {
	const salt = "jlmsuwbz"
	actualResult, ok := run(salt, getHash)
	if !ok {
		t.Errorf("no result found")
	}

	const expectedResult = 35186
	if actualResult != expectedResult {
		t.Errorf("invalid result. got '%v' expected '%v'", actualResult, expectedResult)
	}

}
