package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	lines, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	for _, line := range lines {
		result1 := part1(line)
		fmt.Printf("Part 1: %d\n", result1)
		result2 := part2(line)
		fmt.Printf("Part 2: %d\n", result2)
	}

}

func part1(line string) int {
	var resultBuffer bytes.Buffer
	var markerBuffer bytes.Buffer
	var repeatDataBuffer bytes.Buffer

	var inMarker bool
	var repeatDataLen int
	var repeatDataCount int

	for _, char := range line {
		if repeatDataLen == 0 && char == '(' {
			inMarker = true
			continue
		}

		if inMarker && char == ')' {
			inMarker = false
			mark := markerBuffer.String()
			fmt.Sscanf(mark, "%dx%d", &repeatDataLen, &repeatDataCount)
			markerBuffer.Reset()
			continue
		}

		if inMarker {
			markerBuffer.WriteRune(char)
			continue
		}

		if repeatDataLen > 0 {
			repeatDataBuffer.WriteRune(char)
			repeatDataLen--
			if repeatDataLen == 0 {
				data := repeatDataBuffer.String()
				for i := 0; i < repeatDataCount; i++ {
					resultBuffer.WriteString(data)
				}

				repeatDataBuffer.Reset()
			}

			continue
		}

		resultBuffer.WriteRune(char)
	}

	result := resultBuffer.String()
	//fmt.Printf("raw: %v decomp: %v len: %d\n", line, result, len(result))
	return len(result)
}

func part2(line string) int {
	lineData := []byte(line)
	startIndex := 0
	lineLength := len(lineData)
	return sum(&lineData, lineLength, startIndex, lineLength)
}

func sum(line *[]byte, lineLength int, sectionStart int, sectionLength int) int {
	var count int
	for i := sectionStart; i < sectionStart+sectionLength && i < lineLength; {
		char := (*line)[i]
		if char == '(' {
			i++

			var marker []byte
			for i < sectionStart+sectionLength && i < lineLength {
				char = (*line)[i]
				if char == ')' {
					i++
					break
				}

				marker = append(marker, char)
				i++
			}

			markerString := string(marker)
			var repeatDataLen int
			var repeatDataCount int
			fmt.Sscanf(markerString, "%dx%d", &repeatDataLen, &repeatDataCount)

			count += sum(line, lineLength, i, repeatDataLen) * repeatDataCount

			i += repeatDataLen
		} else {
			count++
			i++
		}
	}
	return count
}

func openInput(name string) ([]string, error) {
	file, error := os.Open("input.txt")
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result, nil
}
