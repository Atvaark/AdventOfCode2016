package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const displayWidth = 50
const displayHeight = 6

type instr struct {
	operator    string
	subOperator string
	operand1    int
	operand2    int
}

func main() {
	instructions, error := openInput("input.txt")
	if error != nil {
		fmt.Printf("unable to open input: %v", error)
		os.Exit(1)
	}

	pixelCount, text := run(instructions)
	fmt.Println("Count: ", pixelCount)
	fmt.Println("Text: ", text)
}

func run(instructions []instr) (pixelCount int, text string) {
	var display [displayHeight][displayWidth]rune
	clearDisplay(&display)

	fmt.Println()
	for _, instruction := range instructions {
		expecute(instruction, &display)
	}

	printDisplay(&display)
	fmt.Println()

	pixelCount = countPixels(&display)
	//fmt.Println("Count: ", pixelCount)

	text = readDisplay(&display)
	//fmt.Println("Text: ", text)

	return pixelCount, text
}

func expecute(in instr, display *[displayHeight][displayWidth]rune) {
	switch in.operator {
	case "rect":
		for y := 0; y < in.operand2; y++ {
			for x := 0; x < in.operand1; x++ {
				display[y][x] = '#'
			}
		}
	case "rotate":
		switch in.subOperator {
		case "column":
			columnIndex := in.operand1
			var tmpColumn [displayHeight]rune
			for i := 0; i < displayHeight; i++ {
				tmpColumn[(i+in.operand2)%displayHeight] = display[i][columnIndex]
			}
			for i := 0; i < displayHeight; i++ {
				display[i][columnIndex] = tmpColumn[i]
			}

		case "row":
			rowIndex := in.operand1
			var newRow [displayWidth]rune
			for i := 0; i < displayWidth; i++ {
				newRow[(i+in.operand2)%displayWidth] = display[rowIndex][i]
			}

			display[rowIndex] = newRow
		}
	}
}

func openInput(name string) ([]instr, error) {
	file, error := os.Open("input.txt")
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []instr
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		splitLen := len(split)
		if splitLen == 0 {
			continue
		}

		var inst instr
		inst.operator = split[0]
		switch inst.operator {
		case "rect":
			dimensions := split[1]
			fmt.Sscanf(dimensions, "%dx%d", &inst.operand1, &inst.operand2)

		case "rotate":
			spec := split[1] // column / row
			inst.subOperator = spec

			param1 := split[2]
			// by := [split[3] // by
			param2 := split[4] // 1
			var test rune
			fmt.Sscanf(param1, "%c=%d", &test, &inst.operand1)
			fmt.Sscanf(param2, "%d", &inst.operand2)

		default:
			fmt.Println("unknown")
		}

		result = append(result, inst)
	}

	return result, nil
}

func printDisplay(display *[displayHeight][displayWidth]rune) {
	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {
			fmt.Printf("%c", (*display)[y][x])
		}
		fmt.Print("\n")
	}
}

func countPixels(display *[displayHeight][displayWidth]rune) int {
	var count int
	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {
			if (*display)[y][x] == '#' {
				count++
			}
		}
	}

	return count
}

func clearDisplay(display *[displayHeight][displayWidth]rune) {
	for y := 0; y < displayHeight; y++ {
		for x := 0; x < displayWidth; x++ {
			display[y][x] = ' '
		}
	}
}

func readDisplay(display *[displayHeight][displayWidth]rune) string {
	const letterWidth = 5
	const letterHeight = 6

	letterCount := displayWidth / letterWidth
	result := make([]rune, letterCount, letterCount)
	for letterIndex := 0; letterIndex < letterCount; letterIndex++ {
		var letterBuffer bytes.Buffer
		for y := 0; y < letterHeight; y++ {
			for x := 0; x < letterWidth; x++ {
				letterBuffer.WriteRune(display[y][x+(letterIndex*letterWidth)])
			}
			letterBuffer.WriteRune('\n')
		}

		lett, ok := letterMap[letterBuffer.String()]
		if !ok {
			lett = '?'
		}

		result[letterIndex] = lett
	}

	return string(result)
}

var letterMap = map[string]rune{
	`#### 
   # 
  #  
 #   
#    
#### 
`: 'Z',
	`#### 
#    
###  
#    
#    
#    
`: 'F',
	`#  # 
#  # 
#### 
#  # 
#  # 
#  # 
`: 'H',
	` ### 
#    
#    
 ##  
   # 
###  
`: 'S',
	` ##  
#  # 
#  # 
#  # 
#  # 
 ##  
`: 'O',
	` ##  
#  # 
#    
# ## 
#  # 
 ### 
`: 'G',
	`###  
#  # 
#  # 
###  
#    
#    
`: 'P',
}
