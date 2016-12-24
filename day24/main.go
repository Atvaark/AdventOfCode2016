package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type position struct {
	x int
	y int
}

type waypoint struct {
	ID  int
	pos position
}

type waypointRoute struct {
	w1       waypoint
	w2       waypoint
	distance int
}

type node struct {
	pos      position
	distance int
}

type grid struct {
	width       int
	height      int
	data        [][]rune
	start       waypoint
	waypoints   map[int]waypoint
	waypointIDs []int
}

func main() {
	grid, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	stepsPart1 := getRequiredSteps(grid, false)
	fmt.Printf("fewest steps part 1 : %d\n", stepsPart1)

	stepsPart2 := getRequiredSteps(grid, true)
	fmt.Printf("fewest steps part 2 : %d\n", stepsPart2)
}

func getRequiredSteps(g *grid, returnToStart bool) int {
	routePermutations := permutateRoutes(g)

	toVisitWaypointIDs := make([]int, 0, len(g.waypointIDs)-1)
	for _, waypointID := range g.waypointIDs {
		if waypointID != g.start.ID {
			toVisitWaypointIDs = append(toVisitWaypointIDs, waypointID)
		}
	}

	waypointPermutations := permutateWaypointIDs(g.start.ID, toVisitWaypointIDs, returnToStart)

	minDistance := -1
	for _, waypointPermutation := range waypointPermutations {
		distance := 0
		for i := 0; i < len(waypointPermutation)-1; i++ {
			startID := waypointPermutation[i]
			endID := waypointPermutation[i+1]

			var lowID int
			var highID int
			if startID < endID {
				lowID = startID
				highID = endID
			} else {
				lowID = endID
				highID = startID
			}

			for _, perm := range routePermutations {
				if perm.w1.ID == lowID && perm.w2.ID == highID {
					distance += perm.distance
					break
				}
			}
		}

		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}
func permutateWaypointIDs(startID int, visitIDs []int, returnToStart bool) [][]int {
	visitPermutations := make([][]int, 0)

	iteratorSlice := make([]int, len(visitIDs), len(visitIDs))
	copy(iteratorSlice, visitIDs)

	iterator := iteratePerms(iteratorSlice)
	for sign := iterator(); sign != 0; sign = iterator() {
		tmpIDs := make([]int, len(iteratorSlice), len(iteratorSlice))
		copy(tmpIDs, iteratorSlice)
		visitPermutations = append(visitPermutations, tmpIDs)
	}

	resultPermutations := make([][]int, 0, len(visitPermutations))
	for _, visitPermutation := range visitPermutations {
		resultPermutation := make([]int, 0, len(visitPermutation)+1)
		resultPermutation = append(resultPermutation, startID)
		resultPermutation = append(resultPermutation, visitPermutation...)

		if returnToStart {
			resultPermutation = append(resultPermutation, startID)
		}

		resultPermutations = append(resultPermutations, resultPermutation)
	}

	return resultPermutations
}


// https://rosettacode.org/wiki/Permutations_by_swapping#Go
func iteratePerms(p []int) func() int {
	f := pf(len(p))
	return func() int {
		return f(p)
	}
}

func pf(n int) func([]int) int {
	sign := 1
	switch n {
	case 0, 1:
		return func([]int) (s int) {
			s = sign
			sign = 0
			return
		}
	default:
		p0 := pf(n - 1)
		i := n
		var d int
		return func(p []int) int {
			switch {
			case sign == 0:
			case i == n:
				i--
				sign = p0(p[:i])
				d = -1
			case i == 0:
				i++
				sign *= p0(p[1:])
				d = 1
				if sign == 0 {
					p[0], p[1] = p[1], p[0]
				}
			default:
				p[i], p[i-1] = p[i-1], p[i]
				sign = -sign
				i += d
			}
			return sign
		}
	}
}

func isVisitedPosition(visited []position, p position) bool {
	for _, v := range visited {
		if v == p {
			return true
		}
	}

	return false
}

func (t *waypointRoute) calculateRouteDistance(g *grid) {
	visited := make([]position, 0)

	const queueLength = 1000000
	nodeQueue := make(chan node, queueLength)
	defer close(nodeQueue)
	nodeQueue <- node{pos: t.w1.pos, distance: 0}
	for {
		select {
		case n := <-nodeQueue:
			adjacentNodes := g.getAdjacentNodes(n)
			for _, adjacentNode := range adjacentNodes {
				if adjacentNode.pos == t.w2.pos {
					t.distance = adjacentNode.distance
					return
				}

				if isVisitedPosition(visited, adjacentNode.pos) {
					continue
				}

				visited = append(visited, adjacentNode.pos)
				nodeQueue <- adjacentNode
			}
		default:
			return
		}
	}
}
func (grid *grid) getAdjacentNodes(n node) []node {
	all := make([]position, 4, 4)

	// top
	all[0] = position{n.pos.x, n.pos.y - 1}
	// right
	all[1] = position{n.pos.x + 1, n.pos.y}
	// bottom
	all[2] = position{n.pos.x, n.pos.y + 1}
	// left
	all[3] = position{n.pos.x - 1, n.pos.y}

	adj := make([]node, 0)
	for _, a := range all {
		if a.x < 0 || a.y < 0 || a.x >= grid.width || a.y >= grid.height {
			continue
		}

		s := grid.data[a.y][a.x]
		const wallSymbol = '#'
		if s == wallSymbol {
			continue
		}

		adj = append(adj, node{pos: a, distance: n.distance + 1})
	}

	return adj
}

func permutateRoutes(g *grid) []*waypointRoute {
	permutations := make([]*waypointRoute, 0)
	for i := 0; i < len(g.waypoints); i++ {
		w1 := g.waypoints[i]
		for j := i + 1; j < len(g.waypoints); j++ {
			w2 := g.waypoints[j]

			t := waypointRoute{w1, w2, -1}
			t.calculateRouteDistance(g)
			permutations = append(permutations, &t)
		}
	}
	return permutations
}

func (grid *grid) printGrid() {
	for _, row := range grid.data {
		for _, col := range row {
			fmt.Printf("%c", col)
		}

		fmt.Println()
	}

	fmt.Println("start")
	fmt.Println(grid.start)

	fmt.Println("waypoints:")
	for _, waypointID := range grid.waypointIDs {
		fmt.Printf("%d:[%d,%d] ", waypointID, grid.waypoints[waypointID].pos.x, grid.waypoints[waypointID].pos.y)
	}

	fmt.Println()
}

func openInput(name string) (*grid, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := make([][]rune, 0)

	width := -1
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)

		switch {
		case width == -1:
			width = len(row)
		case width != len(row):
			return nil, fmt.Errorf("length mismatch. expected %d got %d", width, len(row))
		}

		rows = append(rows, row)
	}
	height := len(rows)

	var startWaypoint waypoint
	var startWaypointFound bool
	waypoints := make(map[int]waypoint)
	for y, row := range rows {
		for x, col := range row {
			var waypointID int
			_, err := fmt.Sscanf(string(col), "%d", &waypointID)
			if err == nil {
				p := waypoint{ID: waypointID, pos: position{x, y}}
				if waypointID == 0 {
					startWaypoint = p
					startWaypointFound = true
				}

				waypoints[waypointID] = p
			}
		}
	}

	if !startWaypointFound {
		return nil, errors.New("no start waypoint found")
	}

	waypointIDs := make([]int, 0, len(waypoints))
	for waypointID := range waypoints {
		waypointIDs = append(waypointIDs, waypointID)
	}
	sort.Ints(waypointIDs)

	return &grid{width: width, height: height, data: rows, start: startWaypoint, waypoints: waypoints, waypointIDs: waypointIDs}, nil
}
