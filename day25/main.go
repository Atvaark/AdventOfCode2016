package main

import (
	"bufio"
	"fmt"
	"strings"
)

const inputDay25 = `cpy a d
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

type outInstr struct {
	x string
}

func (i outInstr) print(c *cpu) {
	fmt.Printf("out %s\n", i.x)
}

type nopInstr struct {
}

func (i nopInstr) print(c *cpu) {
	fmt.Printf("nop\n")
}

type cpu struct {
	registers map[string]int
}

func newCPU(iv int) *cpu {
	var cpu cpu
	cpu.registers = make(map[string]int, 0)
	cpu.registers["a"] = iv
	cpu.registers["b"] = 0
	cpu.registers["c"] = 0
	cpu.registers["d"] = 0
	return &cpu
}

func main() {
	input := openInput(inputDay25)
	iv := getLowestIV(input)
	fmt.Println("part 1: ", iv)
}

func getLowestIV(input []instruction) int {
	const maxIV = 100000000
	for iv := 0; iv < maxIV; iv++ {
		cpu := newCPU(iv)
		isClock := run(cpu, input)
		if isClock {
			return iv
		}
	}

	return -1
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

func run(c *cpu, is []instruction) bool {
	maxSignalCheckCount := 100
	signalCheckCount := 0
	expectedSignal := 0

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
			if c.isRegister(t.reg) {
				c.registers[t.reg]++
			}
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
		case outInstr:
			signal := c.fetchValue(t.x)
			if signal != expectedSignal {
				return false
			}

			switch signal {
			case 0:
				expectedSignal = 1
			case 1:
				expectedSignal = 0
			}

			signalCheckCount++
			if signalCheckCount >= maxSignalCheckCount {
				return true
			}
		default:
			fmt.Printf("unknown instruction: %+v", i)
		}
		pc++
	}

	return false
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
	case outInstr:
		return incInstr{reg: t.x}
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
	case "out":
		i = outInstr{x: split[1]}
	default:
		fmt.Printf("unknown instruction found: %s", line)
	}

	return i
}
