package main

import (
	"bufio"
	"fmt"
	"strings"
)

const input = `Disc #1 has 17 positions; at time=0, it is at position 5.
Disc #2 has 19 positions; at time=0, it is at position 8.
Disc #3 has 7 positions; at time=0, it is at position 1.
Disc #4 has 13 positions; at time=0, it is at position 7.
Disc #5 has 5 positions; at time=0, it is at position 1.
Disc #6 has 3 positions; at time=0, it is at position 0.
`

type disc struct {
	ID     int
	maxPos int
	pos    int
}

func main() {
	discs := openInput()

	firstOkTime, secondOkTime := run(discs)
	fmt.Println("it's okay to press at (1)", firstOkTime)
	fmt.Println("it's okay to press at (2) ", secondOkTime)
}

func run(discs []*disc) (int, int) {
	firstOkTime := -1
	secondOkTime := -1
	const maxTime = 100000000

	for time := 0; time < maxTime; time++ {
		ok := true
		for timeOffset := 0; timeOffset < len(discs); timeOffset++ {
			d := discs[timeOffset]
			dp := (d.pos + time + timeOffset + 1) % d.maxPos
			if dp != 0 {
				ok = false
				break
			}
		}

		if ok {
			if firstOkTime == -1 {
				firstOkTime = time
				discs = append(discs, &disc{ID: len(discs) + 1, maxPos: 11, pos: 0})
				time = 0
			} else {
				secondOkTime = time
				break
			}
		}
	}

	return firstOkTime, secondOkTime
}

func openInput() []*disc {
	discs := make([]*disc, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		d := parseDisk(line)
		discs = append(discs, d)
	}

	return discs
}

func parseDisk(line string) *disc {
	var d disc
	s := strings.Split(line, " ")
	fmt.Sscanf(s[1], "#%d", &d.ID)
	fmt.Sscanf(s[3], "%d", &d.maxPos)
	fmt.Sscanf(s[11], "%d", &d.pos)
	return &d
}
