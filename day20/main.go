package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type blockedRange struct {
	from int
	to   int
}

func main() {
	const ipLength = 0

	blacklist, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	_ = blacklist

	const maxIP = 4294967295
	minIP := -1
	var allowedCount int
Loop:
	for i := 0; i <= maxIP; i++ {
		for _, r := range blacklist {
			if r.from <= i && i <= r.to {
				i = r.to
				continue Loop
			}
		}

		if minIP == -1 {
			minIP = i
		}
		allowedCount++
	}

	fmt.Println("min: ", minIP)
	fmt.Println("allowed: ", allowedCount)
}

func isBlocked(IP int, blacklist []blockedRange) bool {
	for _, blockedIPs := range blacklist {
		if blockedIPs.from <= IP && IP <= blockedIPs.to {
			return true
		}
	}

	return false
}

func openInput(name string) ([]blockedRange, error) {

	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var blacklist []blockedRange
	for scanner.Scan() {
		line := scanner.Text()
		var b blockedRange
		fmt.Sscanf(line, "%d-%d", &b.from, &b.to)
		blacklist = append(blacklist, b)
	}

	sort.Sort(blockedRanges(blacklist))

	return blacklist, nil
}

type blockedRanges []blockedRange

func (slice blockedRanges) Len() int {
	return len(slice)
}

func (slice blockedRanges) Less(i, j int) bool {
	return slice[i].from < slice[j].from
}

func (slice blockedRanges) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
