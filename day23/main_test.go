package main

import "testing"

const inputDay12 = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 16 c
cpy 17 d
inc a
dec d
jnz d -2
dec c
jnz c -5
`

const testInputDay23 = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
`

const inputDay23Test = `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 81 c
jnz 73 d
inc a
inc d
jnz d -2
inc c
jnz c -5
`

func TestRunDay12Part1(t *testing.T) {
	input := openInput(inputDay12)

	cpu1 := newCPU1()
	actual1 := run(cpu1, input)
	const expected1 = 318083
	if actual1 != expected1 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual1, expected1)
	}
}

func TestRunDay12(t *testing.T) {
	input := openInput(inputDay12)

	cpu2 := newCPU2()
	actual2 := run(cpu2, input)
	const expected2 = 9227737
	if actual2 != expected2 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual2, expected2)
	}
}

func TestRunDay23Test(t *testing.T) {
	input := openInput(testInputDay23)

	cpu1 := newCPU1()
	actual1 := run(cpu1, input)
	const expected1 = 3
	if actual1 != expected1 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual1, expected1)
	}
}

func TestRunDay23Part1(t *testing.T) {
	input := openInput(inputDay23Test)

	cpu := newCPU3()
	actual := run(cpu, input)
	const expected = 10953
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

// TODO: Needs peephole optimization. Runs for about 10m
// func TestRunDay23Part2(t *testing.T) {
// 	input := openInput(inputDay23Test)

// 	cpu := newCPU4()
// 	actual := run(cpu, input)
// 	const expected = 479007513
// 	if actual != expected {
// 		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
// 	}
// }
