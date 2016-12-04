package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type room struct {
	name     string
	sectorID int
	checksum string
	isReal   bool
}

func main() {
	rooms, error := openRoomsFile("input.txt")
	if error != nil {
		fmt.Printf("unable to read input: %v", error)
		os.Exit(1)
	}

	part1Result := part1(rooms)
	fmt.Println("part1: ", part1Result)

	part2Result := part2(rooms)
	fmt.Println("part2: ", part2Result)
}

func part1(rooms []room) int {
	var sumSectorIDs int
	for _, room := range rooms {
		if room.isReal {
			sumSectorIDs = sumSectorIDs + room.sectorID
		}
	}
	return sumSectorIDs
}

func part2(rooms []room) int {
	for _, room := range rooms {
		if room.name == "northpole object storage" {
			return room.sectorID
		}
	}

	return -1
}

func openRoomsFile(name string) ([]room, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rooms []room
	for scanner.Scan() {
		line := scanner.Text()
		room := parseRoom(line)
		checkIsRealRoom(&room)
		decryptName(&room)
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func decryptName(r *room) {
	minAlph := int('a')
	maxAlph := int('z')
	diffAlph := maxAlph - minAlph + 1

	decrypted := make([]rune, len(r.name))
	for i, char := range r.name {
		if char == '-' {
			decrypted[i] = ' '
			continue
		}

		value := int(char)
		decryptedValue := minAlph + ((value - minAlph + r.sectorID) % diffAlph)
		decrypted[i] = rune(decryptedValue)
	}

	r.name = string(decrypted[:])
}

func checkIsRealRoom(r *room) {
	checksum := calculateChecksum(r.name)
	r.isReal = r.checksum == checksum
}

func calculateChecksum(name string) string {
	// count chars
	charCountMap := make(map[rune]int)
	for _, char := range name {
		if char == '-' {
			continue
		}

		charCountMap[char] = charCountMap[char] + 1
	}

	// swap char and count pairs
	countCharsMap := make(map[int][]rune)
	for char, count := range charCountMap {
		chars := countCharsMap[count]
		chars = append(chars, char)
		countCharsMap[count] = chars
	}

	// sort counts descending
	var counts []int
	for count := range countCharsMap {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	checksumIndex := 0
	var checksum [5]rune
	for _, count := range counts {
		// sort chars alphabetically
		var chars runes
		chars = countCharsMap[count]
		sort.Sort(chars)

		// build result
		for _, char := range chars {
			checksum[checksumIndex] = char
			checksumIndex = checksumIndex + 1
			if checksumIndex > 4 {
				return string(checksum[:])
			}
		}
	}

	return "invalid_checksum"
}

func parseRoom(line string) room {
	exp, error := regexp.Compile(`([-a-z]+)-(\d+)\[([a-z]+)\]`)
	if error != nil {
		fmt.Println("error compiling: ", error)
		return room{}
	}

	res := exp.FindAllStringSubmatch(line, -1)
	var result room
	result.name = res[0][1]
	result.sectorID, _ = strconv.Atoi(res[0][2])
	result.checksum = res[0][3]
	return result
}

// sort interface for a rune slice
type runes []rune

func (s runes) Len() int           { return len(s) }
func (s runes) Less(i, j int) bool { return s[i] < s[j] }
func (s runes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
