package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

type key struct {
	hash       string
	triplet    string
	quintuplet string
	startIndex int
}

func getHash(knownHashes map[string]string, dataString string) string {
	hashString, ok := knownHashes[dataString]
	if !ok {
		hash := md5.New()
		io.WriteString(hash, dataString)
		hashString = fmt.Sprintf("%x", hash.Sum(nil))
		knownHashes[dataString] = hashString
	}
	return hashString
}

func getStretchedHash(knownHashes map[string]string, dataString string) string {
	const stretchBy = 2016

	hashString, ok := knownHashes[dataString]
	if !ok {
		hashString = dataString
		for i := 0; i < stretchBy+1; i++ {
			hash := md5.New()
			io.WriteString(hash, hashString)
			hashString = fmt.Sprintf("%x", hash.Sum(nil))
		}
		knownHashes[dataString] = hashString
	}
	return hashString
}

type hashfunc func(knownHashes map[string]string, dataString string) string

func main() {
	const saltTest = "abc"
	const salt = "jlmsuwbz"

	part1Result, _ := run(salt, getHash)
	fmt.Println("part 1: ", part1Result)
	part2Result, _ := run(salt, getStretchedHash)
	fmt.Println("part 2: ", part2Result)

}

func run(salt string, h hashfunc) (int, bool) {
	const keyRange = 1000
	const keysRequired = 64
	const maxI = 1000000
	knownHashes := make(map[string]string, 0)
	confirmedKeys := make([]*key, 0)
	for i := 0; i < maxI; i++ {
		dataString := fmt.Sprintf("%s%d", salt, i)
		hashString := h(knownHashes, dataString)
		tripletRune, ok := getTripletRune(hashString)
		if !ok {
			continue
		}

		k := &key{hash: hashString, triplet: getTripleString(tripletRune), quintuplet: getQuintupletString(tripletRune), startIndex: i}

		for j := i + 1; j < i+keyRange+1; j++ {
			dataString = fmt.Sprintf("%s%d", salt, j)
			hashString = h(knownHashes, dataString)
			if strings.Index(hashString, k.quintuplet) >= 0 {
				confirmedKeys = append(confirmedKeys, k)
				if len(confirmedKeys) == keysRequired {
					return k.startIndex, true
				}

				break
			}
		}
	}

	return -1, false
}

func getTripletRune(s string) (rune, bool) {
	c := 0
	r := byte(0)
	for i := 0; i < len(s); i++ {
		if c > 0 && s[i] == r {
			c++
			if c == 3 {
				return rune(r), true
			}
		} else {
			c = 1
			r = s[i]
		}
	}

	return ' ', false
}

func getTripleString(r rune) string {
	return strings.Repeat(string(r), 3)
}

func getQuintupletString(r rune) string {
	return strings.Repeat(string(r), 5)
}
