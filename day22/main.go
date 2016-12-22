package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	key  string
	x    int
	y    int
	size int
	used int
}

func main() {
	nodes, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	fmt.Println("grid:")
	printNodes(nodes)

	viablePairCount := getViablePairCount(nodes)
	fmt.Println("viable pair count: ", viablePairCount)
}

func printNodes(nodes []node) {
	var maxX int
	var maxY int
	for _, n := range nodes {
		if n.x > maxX {
			maxX = n.x
		}

		if n.y > maxY {
			maxY = n.y
		}
	}

	grid := make([][]node, maxX+1, maxX+1)
	for x := 0; x <= maxX; x++ {
		column := make([]node, maxY+1, maxY+1)
		grid[x] = column
	}

	for _, n := range nodes {
		grid[n.x][n.y] = n
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			// fmt.Printf("%3d/%3d ", grid[x][y].used, grid[x][y].size)
			n := grid[x][y]

			switch {
			case n.x == 0 && n.y == 0:
				fmt.Print("S")
			case n.x == maxX && n.y == 0:
				fmt.Print("G")
			case n.used == 0:
				fmt.Print("E")
			case n.used > 400:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func getViablePairCount(nodes []node) int {
	var viablePairCount int
	for i := 0; i < len(nodes); i++ {
		nodeA := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			nodeB := nodes[j]
			if nodeA.used > 0 && nodeA.used <= (nodeB.size-nodeB.used) {
				viablePairCount++
			}

			if nodeB.used > 0 && nodeB.used <= (nodeA.size-nodeA.used) {
				viablePairCount++
			}
		}
	}

	return viablePairCount
}

func openInput(name string) ([]node, error) {
	file, error := os.Open(name)
	if error != nil {
		return nil, error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nodes []node
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "/dev/grid/") == -1 {
			continue
		}

		n := parseNode(line)
		nodes = append(nodes, n)
	}

	return nodes, nil
}

func parseNode(line string) node {
	var n node
	fields := strings.Fields(line)
	n.key = fields[0]
	fmt.Sscanf(fields[0], "/dev/grid/node-x%d-y%d", &n.x, &n.y)
	fmt.Sscanf(fields[1], "%dT", &n.size)
	fmt.Sscanf(fields[2], "%dT", &n.used)
	return n
}
