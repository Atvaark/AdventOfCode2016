package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	const input = "abbhdwsy"

	fmt.Println("input: ", input)
	result1 := part1(input)
	fmt.Println("result1: ", result1)
	result2 := part2(input)
	fmt.Println("result2: ", result2)
}

func part1(input string) string {
	const passwordLength = 8

	var password [passwordLength]rune
	var passwordIndex int
	var suffix int
	for ; passwordIndex < passwordLength; suffix++ {
		indexedInput := fmt.Sprintf("%s%d", input, suffix)
		hash := md5.New()
		io.WriteString(hash, indexedInput)
		hashString := fmt.Sprintf("%x", hash.Sum(nil))

		if strings.Index(hashString, "00000") == 0 {
			var r rune
			r = rune(hashString[5])
			password[passwordIndex] = r
			passwordIndex = passwordIndex + 1
			//fmt.Println("ok: ", passwordIndex, " hash: ", hashString, " rune: ", fmt.Sprintf("%c", r))
		}
	}

	return string(password[:])
}

func part2(input string) string {
	const passwordLength = 8

	var password [passwordLength]rune
	var passwordCount int
	var suffix int
	for ; passwordCount < passwordLength; suffix++ {
		indexedInput := fmt.Sprintf("%s%d", input, suffix)
		hash := md5.New()
		io.WriteString(hash, indexedInput)
		hashString := fmt.Sprintf("%x", hash.Sum(nil))

		if strings.Index(hashString, "00000") == 0 {
			index, err := strconv.Atoi(hashString[5:6])
			if err != nil || index >= passwordLength || password[index] != rune(0) {
				continue
			}

			passwordRune := rune(hashString[6])
			password[index] = passwordRune
			passwordCount++
			//fmt.Println("ok: ", passwordCount, " hash: ", hashString, " rune: ", fmt.Sprintf("%c", indexRune))
		}
	}

	return string(password[:])
}
