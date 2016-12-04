package main

import (
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	inputData, error := ioutil.ReadFile("input.txt")
	if error != nil {
		t.Errorf("unable to open input: %v", error)
		return
	}

	inputSequence := string(inputData)

	result := run(inputSequence)
	if result.distanceEnd != 307 {
		t.Error("unexpected distanceEnd")
	}

	if result.distanceFirstIntersection != 165 {
		t.Error("unexpected distanceFirstIntersection")
	}
}
