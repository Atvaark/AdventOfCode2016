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
	type state struct {
		resultBuffer     bytes.Buffer
		markerBuffer     bytes.Buffer
		repeatDataBuffer bytes.Buffer
		inMarker         bool
		repeatDataLen    int
		repeatDataCount  int
		decompressedLen  int
	}

	var s state
	var handle func(char rune, s *state)
	handle = func(char rune, s *state) {
		if s.repeatDataLen == 0 && char == '(' {
			s.inMarker = true
			return
		}

		if s.inMarker && char == ')' {
			s.inMarker = false
			mark := s.markerBuffer.String()
			fmt.Sscanf(mark, "%dx%d", &s.repeatDataLen, &s.repeatDataCount)
			s.markerBuffer.Reset()
			return
		}

		if s.inMarker {
			s.markerBuffer.WriteRune(char)
			return
		}

		if s.repeatDataLen > 0 {
			s.repeatDataBuffer.WriteRune(char)
			s.repeatDataLen--
			if s.repeatDataLen == 0 {
				data := s.repeatDataBuffer.String()
				s.repeatDataBuffer.Reset()

				var repeatedDataBuffer bytes.Buffer
				for i := 0; i < s.repeatDataCount; i++ {
					repeatedDataBuffer.WriteString(data)
				}
				repeatedData := repeatedDataBuffer.String()
				repeatedDataBuffer.Reset()

				for _, repeatedChar := range repeatedData {
					handle(repeatedChar, s)
				}
			}

			return
		}

		s.decompressedLen++
	}

	for _, char := range line {
		handle(char, &s)
	}

	//result := s.resultBuffer.String()
	//fmt.Printf("raw: %v decomp: %v len: %d %d\n", line, result, len(result), s.decompressedLength)
	return s.decompressedLen
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
