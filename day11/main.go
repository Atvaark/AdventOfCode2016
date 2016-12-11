package main

import (
	"bytes"
	"fmt"
	"sort"
)

const input = `The first floor contains a strontium generator, a strontium-compatible microchip, a plutonium generator, and a plutonium-compatible microchip.
The second floor contains a thulium generator, a ruthenium generator, a ruthenium-compatible microchip, a curium generator, and a curium-compatible microchip.
The third floor contains a thulium-compatible microchip.
The fourth floor contains nothing relevant.
`

const topFloorIndex = 4
const bottomFloorIndex = 1

func main() {
	arrangement := openInput()
	stepCount := run(arrangement)
	fmt.Printf("Step count : %d\n", stepCount)
}

type arrangement struct {
	items []item
}

func (a arrangement) String() string {
	var buffer bytes.Buffer
	for _, i := range a.items {
		buffer.WriteString(i.String())
		buffer.WriteString(" ")
	}
	return buffer.String()
}

type item struct {
	typeID    itemType
	elementID elementType
	floorID   int
}

func (i item) String() string {
	var element string
	switch i.elementID {
	case curium:
		element = "C"
	case dilithium:
		element = "D"
	case elerium:
		element = "E"
	case hydrogen:
		element = "H"
	case lithium:
		element = "L"
	case plutonium:
		element = "P"
	case ruthenium:
		element = "R"
	case strontium:
		element = "S"
	case thulium:
		element = "T"
	default:
		element = ""
	}

	var typeID string
	switch i.typeID {
	case elevator:
		typeID = "E"
	case generator:
		typeID = "G"
	case microchip:
		typeID = "M"
	default:
		typeID = ""

	}

	return fmt.Sprintf("%s%s@%d", element, typeID, i.floorID)
}

type itemType int

const (
	noItem itemType = iota
	elevator
	microchip
	generator
)

type elementType int

const (
	noElement elementType = iota
	curium
	dilithium
	elerium
	hydrogen
	lithium
	plutonium
	ruthenium
	strontium
	thulium
)

type arrangementNode struct {
	arr      arrangement
	distance int
}

func openInput() arrangement {
	// TODO: Parse input

	var input arrangement
	input.items = []item{
		item{elementID: noElement, typeID: elevator, floorID: 1},
		item{elementID: strontium, typeID: generator, floorID: 1},
		item{elementID: strontium, typeID: microchip, floorID: 1},
		item{elementID: plutonium, typeID: generator, floorID: 1},
		item{elementID: plutonium, typeID: microchip, floorID: 1},

		item{elementID: thulium, typeID: generator, floorID: 2},
		item{elementID: ruthenium, typeID: generator, floorID: 2},
		item{elementID: ruthenium, typeID: microchip, floorID: 2},
		item{elementID: curium, typeID: generator, floorID: 2},
		item{elementID: curium, typeID: microchip, floorID: 2},

		item{elementID: thulium, typeID: microchip, floorID: 3},

		// part2
		item{elementID: elerium, typeID: generator, floorID: 1},
		item{elementID: elerium, typeID: microchip, floorID: 1},
		item{elementID: dilithium, typeID: generator, floorID: 1},
		item{elementID: dilithium, typeID: microchip, floorID: 1},

		// Test
		// item{elementID: noElement, typeID: elevator, floorID: 1},
		// item{elementID: hydrogen, typeID: microchip, floorID: 1},
		// item{elementID: lithium, typeID: microchip, floorID: 1},

		// item{elementID: hydrogen, typeID: generator, floorID: 2},

		// item{elementID: lithium, typeID: generator, floorID: 3},
	}

	reorder(input)
	return input
}

type byTypes []item

func (s byTypes) Len() int {
	return len(s)
}

func (s byTypes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byTypes) Less(i, j int) bool {
	if s[i].floorID < s[j].floorID {
		return true
	}

	if s[i].floorID > s[j].floorID {
		return false
	}

	if s[i].elementID < s[j].elementID {
		return true
	}

	if s[i].elementID > s[j].elementID {
		return false
	}

	return s[i].typeID < s[j].typeID
}

func reorder(a arrangement) {
	a.items = reorderItems(a.items)
}

func reorderItems(items []item) []item {
	sort.Sort(byTypes(items))
	return items
}

func isDone(items []item) bool {
	for _, i := range items {
		if i.floorID != topFloorIndex {
			return false
		}
	}

	return true
}

func run(root arrangement) int {
	nodeQueues := make(map[int][]arrangementNode, 0)
	nodeQueues[0] = []arrangementNode{arrangementNode{arr: root}}

	checkedNodes := make(map[string]bool)
	exitNodes := make([]arrangementNode, 0)

	minDistance := -1
	currentDistance := 0
	for {
		nodeQueue, ok := nodeQueues[currentDistance]
		if !ok {
			break
		}

		if len(nodeQueue) == 0 {
			currentDistance++
			//fmt.Println("max distance increased to ", currentDistance)
			continue
		}

		var node arrangementNode
		node = nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		nodeQueues[currentDistance] = nodeQueue
		nextNodes := step(node)
		for _, nextNode := range nextNodes {
			if isDone(nextNode.arr.items) {
				exitNodes = append(exitNodes, nextNode)
			} else {
				nextNodeKey := nextNode.arr.String()
				_, ok := checkedNodes[nextNodeKey]
				if !ok {
					nextNodeQueue, ok := nodeQueues[nextNode.distance]
					if !ok {
						nextNodeQueue = []arrangementNode{}
					}

					nextNodeQueue = append(nextNodeQueue, nextNode)
					nodeQueues[nextNode.distance] = nextNodeQueue
					checkedNodes[nextNodeKey] = true
				}
			}

		}
	}
	for _, finalNode := range exitNodes {
		if minDistance == -1 || finalNode.distance < minDistance {
			minDistance = finalNode.distance
		}
	}

	return minDistance
}

func step(node arrangementNode) []arrangementNode {
	var nextNodes = make([]arrangementNode, 0)
	items := node.arr.items

	elevatorFloor := -1
	for _, item := range items {
		if item.typeID == elevator {
			elevatorFloor = item.floorID
			break
		}
	}

	minFloor := getMinFloor(items)
	floorItems := filterByFloor(items, elevatorFloor)
	floorPerms := getPermutations(floorItems)

	for _, floorPerm := range floorPerms {
		canGoUp := elevatorFloor < topFloorIndex
		canGoDown := elevatorFloor > minFloor

		if canGoUp {
			next := apply(items, floorPerm, 1)
			if validateItems(next) {
				nextArrangement := arrangementNode{arr: arrangement{items: next}, distance: node.distance + 1}
				nextNodes = append(nextNodes, nextArrangement)
			}
		}

		if canGoDown {
			next := apply(items, floorPerm, -1)
			if validateItems(next) {
				nextArrangement := arrangementNode{arr: arrangement{items: next}, distance: node.distance + 1}
				nextNodes = append(nextNodes, nextArrangement)
			}
		}
	}

	return nextNodes
}

func arrangementEquals(i arrangement, j arrangement) bool {
	return i.String() == j.String()
}

func validateItems(items []item) bool {
	for _, item1 := range items {
		if item1.typeID != microchip {
			continue
		}

		var hasShield bool
		var needsShield bool

		for _, item2 := range items {
			if item2.floorID != item1.floorID {
				continue
			}

			if item2.typeID != generator {
				continue
			}

			needsShield = true

			if item2.elementID == item1.elementID {
				hasShield = true
				break
			}
		}

		if needsShield && !hasShield {
			return false
		}
	}
	return true
}

func itemEquals(i item, j item) bool {
	return i.elementID == j.elementID && i.floorID == j.floorID && i.typeID == j.typeID
}

func getPermutations(items []item) [][]item {
	set := make(map[string][]item, 0)

	_ = set

	for _, item1 := range items {
		set[item1.String()] = []item{item1}

		for _, item2 := range items {
			if itemEquals(item1, item2) {
				continue
			}

			p1 := []item{item1, item2}
			p1 = reorderItems(p1)
			p1Key := fmt.Sprintf("%v", p1)

			_, ok := set[p1Key]
			if !ok {
				set[p1Key] = p1
			}
		}
	}

	result := make([][]item, 0)
	for _, p := range set {
		result = append(result, p)
	}

	return result
}

func filterByFloor(items []item, floodID int) []item {
	result := make([]item, 0)
	for _, i := range items {
		if i.floorID == floodID && i.typeID != elevator {
			result = append(result, i)
		}
	}

	return result
}

func getMinFloor(items []item) int {
	result := topFloorIndex
	for _, item := range items {
		if item.floorID < result {
			result = item.floorID
		}
	}
	return result
}

func apply(items []item, permItems []item, dir int) []item {
	next := make([]item, len(items))
	copy(next, items)

	for _, floorPermItem := range permItems {
		for nextItemIndex, nextItem := range next {
			if itemEquals(nextItem, floorPermItem) {
				nextItem.floorID = nextItem.floorID + dir
				next[nextItemIndex] = nextItem
			}
		}
	}

	for elevatorItemIndex, elevatorItem := range next {
		if elevatorItem.typeID == elevator {
			elevatorItem.floorID = elevatorItem.floorID + dir
			next[elevatorItemIndex] = elevatorItem
		}
	}

	next = reorderItems(next)
	return next
}
