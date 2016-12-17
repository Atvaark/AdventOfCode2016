package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

type doorState struct {
	upOpen    bool
	downOpen  bool
	leftOpen  bool
	rightOpen bool
}

type position struct {
	x int
	y int
}

func main() {
	const passcode = "gdjjyniy"
	path, dist := run(passcode)
	fmt.Printf("part 1: %s\n", path)
	fmt.Printf("part 2: %d\n", dist)
}

const gridWidth = 4
const gridHeight = 4

type pathNode struct {
	path     string
	pos      position
	distance int
}

type nodeQueue []*pathNode

func (s *nodeQueue) dequeue() *pathNode {
	index := 0
	e := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return e
}

func run(passcode string) (shortest string, longestPathDistance int) {
	inputNodes := make(nodeQueue, 0)
	outputNodes := make(nodeQueue, 0)
	inputNodes = append(inputNodes, &pathNode{})

	currentDistance := 0
	longestPath := &pathNode{}
	shortestPath := &pathNode{}
	shortestPathFound := false

	for {
		if len(inputNodes) == 0 {
			if len(outputNodes) == 0 {
				break
			}
			tmp := inputNodes
			inputNodes = outputNodes
			outputNodes = tmp
			currentDistance++

			continue
		}

		node := inputNodes.dequeue()
		doorstate := calcDoorState(passcode, node.path, node.pos)
		neighbors := getNeighbors(node, doorstate)

		for _, neighbor := range neighbors {
			if neighbor.pos.x == (gridWidth-1) && neighbor.pos.y == (gridHeight-1) {
				if !shortestPathFound {
					shortestPath = neighbor
					shortestPathFound = true
				}

				if neighbor.distance > longestPath.distance {
					longestPath = neighbor
				}

				continue
			}

			outputNodes = append(outputNodes, neighbor)
		}
	}

	return shortestPath.path, longestPath.distance
}

func getNeighbors(startNode *pathNode, doorstate doorState) []*pathNode {
	neighbors := make([]*pathNode, 0)
	if doorstate.upOpen {
		neighbors = append(neighbors, &pathNode{path: startNode.path + "U", pos: position{x: startNode.pos.x, y: startNode.pos.y - 1}, distance: startNode.distance + 1})
	}
	if doorstate.downOpen {
		neighbors = append(neighbors, &pathNode{path: startNode.path + "D", pos: position{x: startNode.pos.x, y: startNode.pos.y + 1}, distance: startNode.distance + 1})
	}
	if doorstate.leftOpen {
		neighbors = append(neighbors, &pathNode{path: startNode.path + "L", pos: position{x: startNode.pos.x - 1, y: startNode.pos.y}, distance: startNode.distance + 1})
	}
	if doorstate.rightOpen {
		neighbors = append(neighbors, &pathNode{path: startNode.path + "R", pos: position{x: startNode.pos.x + 1, y: startNode.pos.y}, distance: startNode.distance + 1})
	}

	return neighbors
}

func calcDoorState(passcode string, path string, pos position) doorState {
	hash := md5.New()
	io.WriteString(hash, passcode)
	io.WriteString(hash, path)
	doorStateCodes := fmt.Sprintf("%x", hash.Sum(nil))
	doorStateCodes = string(doorStateCodes[:4])

	var state doorState
	for i := 0; i < 4; i++ {
		doorStateCode := doorStateCodes[i]
		if doorStateCode >= 'b' && doorStateCode <= 'f' {
			switch i {
			case 0:
				state.upOpen = pos.y > 0
			case 1:
				state.downOpen = pos.y < (gridHeight - 1)
			case 2:
				state.leftOpen = pos.x > 0
			case 3:
				state.rightOpen = pos.x < (gridWidth - 1)
			}
		}
	}

	return state
}
