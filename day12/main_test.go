package main

import "testing"

func TestRun(t *testing.T) {
	input := openInput()

	cpu1 := newCPU1()
	actual1 := run(cpu1, input)
	const expected1 = 318083
	if actual1 != expected1 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual1, expected1)
	}

	cpu2 := newCPU2()
	actual2 := run(cpu2, input)
	const expected2 = 9227737
	if actual2 != expected2 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual1, expected2)
	}
}
