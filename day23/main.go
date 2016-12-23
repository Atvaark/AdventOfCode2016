package main

import (
	"bufio"
	"fmt"
	"strings"
)

const inputDay23 = `cpy a b
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

type instruction interface {
	print(c *cpu)
}

type cpyInstr struct {
	left  string
	right string
}

func (i cpyInstr) print(c *cpu) {
	fmt.Printf("cpy %s %s\n", i.left, i.right)
}

type incInstr struct {
	reg string
}

func (i incInstr) print(c *cpu) {
	fmt.Printf("inc %s\n", i.reg)
}

type decInstr struct {
	reg string
}

func (i decInstr) print(c *cpu) {
	fmt.Printf("dec %s\n", i.reg)
}

type jnzInstr struct {
	left  string
	right string
}

func (i jnzInstr) print(c *cpu) {
	fmt.Printf("jnz %s %s\n", i.left, i.right)
}

type tglInstr struct {
	reg string
}

func (i tglInstr) print(c *cpu) {
	fmt.Printf("tgl %s\n", i.reg)
}

type nopInstr struct {
}

func (i nopInstr) print(c *cpu) {
	fmt.Printf("nop\n")
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

func newCPU3() *cpu {
	var cpu cpu
	cpu.registers = make(map[string]int, 0)
	cpu.registers["a"] = 7
	cpu.registers["b"] = 0
	cpu.registers["c"] = 0
	cpu.registers["d"] = 0
	return &cpu
}

func newCPU4() *cpu {
	var cpu cpu
	cpu.registers = make(map[string]int, 0)
	cpu.registers["a"] = 12
	cpu.registers["b"] = 0
	cpu.registers["c"] = 0
	cpu.registers["d"] = 0
	return &cpu
}

func main() {
	input := openInput(inputDay23)

	cpu3 := newCPU3()
	result3 := run(cpu3, input)
	fmt.Println("part 3: ", result3)

	input = openInput(inputDay23)

	cpu4 := newCPU4()
	result4 := run(cpu4, input)
	fmt.Println("part 4: ", result4)
}

func (c *cpu) fetchValue(v string) int {
	value, ok := c.registers[v]
	if ok {
		return value
	}

	fmt.Sscanf(v, "%d", &value)
	return value
}

func (c *cpu) isRegister(v string) bool {
	_, ok := c.registers[v]
	return ok
}

func run(c *cpu, is []instruction) int {
	pc := 0
	end := len(is)
	for pc < end {
		i := is[pc]

		switch t := i.(type) {
		case cpyInstr:
			if !c.isRegister(t.right) {
				break
			}

			left := c.fetchValue(t.left)
			c.registers[t.right] = left
		case incInstr:
			c.registers[t.reg]++
		case decInstr:
			c.registers[t.reg]--
		case jnzInstr:
			left := c.fetchValue(t.left)
			if left != 0 {
				pc += c.fetchValue(t.right)
				continue
			}
		case tglInstr:
			pctg := pc + c.registers[t.reg]
			if pctg < end {
				ti := is[pctg]
				ti = toggle(c, ti)
				is[pctg] = ti
			}
		default:
			fmt.Printf("unknown instruction: %+v", i)
		}
		pc++
	}

	result, _ := c.registers["a"]
	return result
}

func toggle(c *cpu, i instruction) instruction {
	switch t := i.(type) {
	case cpyInstr:
		return jnzInstr{left: t.left, right: t.right}
	case incInstr:
		return decInstr{reg: t.reg}
	case decInstr:
		return incInstr{reg: t.reg}
	case jnzInstr:
		return cpyInstr{left: t.left, right: t.right}
	case tglInstr:
		return incInstr{reg: "a"}
	}

	fmt.Printf("unknown instruction to toggle: %+v", i)
	return nopInstr{}
}

func openInput(testInput string) []instruction {
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
		right := split[2]
		i = cpyInstr{left, right}
	case "inc":
		i = incInstr{reg: split[1]}
	case "dec":
		i = decInstr{reg: split[1]}
	case "jnz":
		left := split[1]
		right := split[2]
		i = jnzInstr{left, right}
	case "tgl":
		i = tglInstr{reg: split[1]}
	default:
		fmt.Printf("unknown instruction found: %s", line)
	}

	return i
}
