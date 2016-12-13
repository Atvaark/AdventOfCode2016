package main

import "fmt"

//const part1Seed = 1352
//const testSeed = 10
const testSeed = 1352

type pos struct {
	x int
	y int
}

func (p *pos) equals(other *pos) bool {
	return p.x == other.x && p.y == other.y
}

type pathNode struct {
	p        *pos
	distance int
}

func removeAtIndex(s *[]*pathNode, index int) *pathNode {
	e := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return e
}

func wasVisited(visited *[]*pos, p *pos) bool {
	for _, v := range *visited {
		if v.equals(p) {
			return true
		}
	}

	return false
}

func main() {
	start := &pos{1, 1}
	end := &pos{31, 39}
	seed := testSeed

	part1Result, part2Result := run(start, end, seed)

	fmt.Println("part 1:", part1Result)
	fmt.Println("part 2:", part2Result)
}

func run(start *pos, end *pos, seed int) (int, int) {
	// printMaze(start, end, seed)
	visited := make([]*pos, 0)
	inputQueue := make([]*pathNode, 0)
	inputQueue = append(inputQueue, &pathNode{start, 0})
	outputQueue := make([]*pathNode, 0)

	minToEnd := -1
	maxDistance := 50
	maxDistinctResult := -1
	currentDistance := 0

	for {
		if len(inputQueue) == 0 {
			if len(outputQueue) == 0 {
				break
			}

			tmp := inputQueue
			inputQueue = outputQueue
			outputQueue = tmp
			currentDistance++

			if currentDistance == maxDistance {
				maxDistinctResult = len(visited)
			}

			continue
		}

		p := removeAtIndex(&inputQueue, 0)

		neighbors := getNeighbors(p, seed)

		for _, n := range neighbors {
			if wasVisited(&visited, n.p) {
				continue
			}

			if n.p.equals(end) {
				if minToEnd == -1 || n.distance < minToEnd {
					minToEnd = n.distance
				}
			}

			visited = append(visited, n.p)
			outputQueue = append(outputQueue, n)
		}
	}

	//printVisited(visited, seed)
	return minToEnd, maxDistinctResult
}

func printVisited(visited []*pos, seed int) {
	maxX := 0
	maxY := 0
	for _, v := range visited {
		if v.x > maxX {
			maxX = v.x
		}

		if v.y > maxY {
			maxY = v.y
		}
	}

	maxX++
	maxY++
	maze := make([][]rune, 0)
	for y := 0; y < maxY; y++ {
		cy := make([]rune, maxX)
		for x := 0; x < maxX; x++ {
			if isWall(x, y, seed) {
				cy[x] = '#'
			} else {
				cy[x] = ' '
			}
		}
		maze = append(maze, cy)
	}

	for _, v := range visited {
		r := maze[v.y]
		r[v.x] = '.'
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%c", maze[y][x])
		}
		fmt.Println()
	}
}

func getNeighbors(n *pathNode, seed int) []*pathNode {
	top := pos{x: n.p.x, y: n.p.y - 1}
	right := pos{x: n.p.x + 1, y: n.p.y}
	bottom := pos{x: n.p.x, y: n.p.y + 1}
	left := pos{x: n.p.x - 1, y: n.p.y}

	neighbors := make([]*pathNode, 0)
	if top.isValid(seed) {
		neighbors = append(neighbors, &pathNode{p: &top, distance: n.distance + 1})
	}
	if right.isValid(seed) {
		neighbors = append(neighbors, &pathNode{p: &right, distance: n.distance + 1})
	}
	if bottom.isValid(seed) {
		neighbors = append(neighbors, &pathNode{p: &bottom, distance: n.distance + 1})
	}
	if left.isValid(seed) {
		neighbors = append(neighbors, &pathNode{p: &left, distance: n.distance + 1})
	}

	return neighbors
}

func printMaze(s *pos, e *pos, seed int) {
	for y := 0; y < 7; y++ {
		for x := 0; x < 10; x++ {
			p := pos{x, y}
			switch {
			case p.equals(s):
				fmt.Print("S")
			case p.equals(e):
				fmt.Print("E")
			case isWall(x, y, seed):
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (p *pos) isValid(seed int) bool {
	return p.x >= 0 && p.y >= 0 && !isWall(p.x, p.y, seed)
}

func isWall(x int, y int, seed int) bool {
	tmp := x*x + 3*x + 2*x*y + y + y*y
	tmp += seed

	var sumBits int
	for tmp > 0 {
		if tmp&1 > 0 {
			sumBits++
		}

		tmp = tmp >> 1
	}

	return sumBits%2 == 1
}
