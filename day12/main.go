package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type instruction interface {
}

type cpyVarInstr struct {
	leftReg  string
	rightReg string
}

type cpyConstInstr struct {
	left     int
	rightReg string
}

type incInstr struct {
	reg string
}

type decInstr struct {
	reg string
}

type jnzVarInstr struct {
	leftReg string
	right   int
}

type jnzConstInstr struct {
	left  int
	right int
}

type cpu struct {
	registers map[string]int
}

func newCPU1() *cpu {
	var cpu cpu
	cpu.registers = make(map[string]int, 0)
	cpu.registers["a"] = 0
	cpu.registers["b"] = 0
	cpu.registers["c"] = 0
	cpu.registers["d"] = 0
	return &cpu
}

func newCPU2() *cpu {
	var cpu cpu
	cpu.registers = make(map[string]int, 0)
	cpu.registers["a"] = 0
	cpu.registers["b"] = 0
	cpu.registers["c"] = 1
	cpu.registers["d"] = 0
	return &cpu
}

func main() {
	input := openInput()

	cpu1 := newCPU1()
	result1 := run(cpu1, input)
	fmt.Println("part 1: ", result1)

	cpu2 := newCPU2()
	result2 := run(cpu2, input)
	fmt.Println("part 2: ", result2)
}

func run(c *cpu, is []instruction) int {
	pc := 0
	end := len(is)
	for pc < end {
		i := is[pc]
		switch t := i.(type) {
		case cpyVarInstr:
			c.registers[t.rightReg] = c.registers[t.leftReg]
		case cpyConstInstr:
			c.registers[t.rightReg] = t.left
		case incInstr:
			c.registers[t.reg]++
		case decInstr:
			c.registers[t.reg]--
		case jnzVarInstr:
			if c.registers[t.leftReg] != 0 {
				pc += t.right
				continue
			}
		case jnzConstInstr:
			if t.left != 0 {
				pc += t.right
				continue
			}
		}
		pc++
	}

	result, _ := c.registers["a"]
	return result
}

func openInput() []instruction {
	const testInput = `cpy 1 a
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

	scanner := bufio.NewScanner(strings.NewReader(testInput))
	scanner.Split(bufio.ScanLines)
	is := make([]instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instr := parseInstruction(line)
		is = append(is, instr)
	}

	return is
}

func parseInstruction(line string) instruction {
	split := strings.Split(line, " ")
	var i instruction
	op := split[0]
	switch op {
	case "cpy":
		left := split[1]
		leftConst, err := strconv.Atoi(left)
		right := split[2]
		if err == nil {
			i = cpyConstInstr{left: leftConst, rightReg: right}
		} else {
			i = cpyVarInstr{leftReg: left, rightReg: right}
		}
	case "inc":
		i = incInstr{reg: split[1]}
	case "dec":
		i = decInstr{reg: split[1]}
	case "jnz":
		left := split[1]
		leftConst, err := strconv.Atoi(left)
		offset, _ := strconv.Atoi(split[2])
		if err == nil {
			i = jnzConstInstr{left: leftConst, right: offset}
		} else {
			i = jnzVarInstr{leftReg: split[1], right: offset}
		}
	default:
	}

	return i
}
