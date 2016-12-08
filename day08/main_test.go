package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	instructions, error := openInput("input.txt")
	if error != nil {
		t.Errorf("unable to open input: %v", error)
		return
	}

	pixelCount, text := run(instructions)

	const expectedPixelCount = 119
	if pixelCount != expectedPixelCount {
		t.Errorf("unexpected pixel count %d expected %d", pixelCount, expectedPixelCount)
	}

	const expectedText = "ZFHFSFOGPO"
	if text != expectedText {
		t.Errorf("unexpected text %s expected %s", text, expectedText)
	}
}
