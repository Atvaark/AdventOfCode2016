package main

import "testing"

const testInputDay25 = `cpy a d
cpy 9 c
cpy 282 b
inc d
dec b
jnz b -2
dec c
jnz c -5
cpy d a
jnz 0 0
cpy a b
cpy 0 a
cpy 2 c
jnz b 2
jnz 1 6
dec b
dec c
jnz c -4
inc a
jnz 1 -7
cpy 2 b
jnz c 2
jnz 1 4
dec b
dec c
jnz 1 -4
jnz 0 0
out b
jnz a -19
jnz 1 -21
`

func TestGetLowestIV(t *testing.T) {
	input := openInput(testInputDay25)
	iv := getLowestIV(input)
	const expected = 192
	if iv != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", iv, expected)
	}
}
