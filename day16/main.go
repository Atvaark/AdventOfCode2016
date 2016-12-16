package main

import (
	"fmt"
)

func main() {
	const part1Length = 272
	const part2Length = 35651584
	const input = "10010000000110000"

	part1Checksum := getChecksum(input, part1Length)
	fmt.Println("part 1: ", part1Checksum)
	part2Checksum := getChecksum(input, part2Length)
	fmt.Println("part 2: ", part2Checksum)
}

func getChecksum(input string, length int) string {
	// Dragon curve
	curve := input
	for len(curve) < length {
		curve = step(curve)
	}

	if len(curve) > length {
		curve = curve[:length]
	}

	// checksum
	checksum := curve
	for len(checksum)%2 == 0 {
		nextChecksumLength := len(checksum) / 2
		nextChecksum := make([]byte, nextChecksumLength)

		for i := 0; i < nextChecksumLength; i++ {
			a := checksum[i*2]
			b := checksum[(i*2)+1]
			c := '0'
			if a == b {
				c = '1'
			}
			nextChecksum[i] = byte(c)
		}

		checksum = string(nextChecksum)
	}

	return checksum
}

func step(data string) string {
	bb := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		aIndex := len(data) - i - 1
		tmp := data[aIndex]

		switch rune(tmp) {
		case '1':
			tmp = '0'
		case '0':
			tmp = '1'
		}

		bb[i] = tmp

	}

	b := string(bb)

	return data + "0" + b
}
