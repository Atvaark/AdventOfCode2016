package main

import (
	"testing"
)

func TestOpenShapeFile(t *testing.T) {
	shapes, error := openShapeFile("input.txt")
	if error != nil {
		t.Errorf("open shape file failed: %v", error)
	}
	_ = shapes
}

func TestPart1(t *testing.T) {
	shapes, error := openShapeFile("input.txt")
	if error != nil {
		t.Error("open shape file failed")
	}

	count := part1(shapes)
	const expected = 983
	if count != expected {
		t.Errorf("invalid result '%d' expected '%d'", count, expected)
	}
}

func TestPart2(t *testing.T) {
	shapes, error := openShapeFile("input.txt")
	if error != nil {
		t.Error("open shape file failed")
	}

	count := part2(shapes)
	const expected = 1836
	if count != expected {
		t.Errorf("invalid result '%d' expected '%d'", count, expected)
	}
}
