package main

import "fmt"

func main() {
	const elfCount = 3014603
	finalElfNumber1 := runPart1(elfCount)
	fmt.Printf("part 1 elf who gets all presents: %d\n", finalElfNumber1) // ok: 1834903

	finalElfNumber2 := runPart2(elfCount)
	fmt.Printf("part 2 elf who gets all presents: %d\n", finalElfNumber2) // ok: 1420280

}

type elf struct {
	nr           int
	presentCount int
}

func runPart1(elfCount int) int {
	elves := make([]*elf, elfCount, elfCount)
	for i := 0; i < elfCount; i++ {
		elves[i] = &elf{nr: i + 1, presentCount: 1}
	}

	i := 0
	for {
		if i >= elfCount {
			i = 0
			continue
		}

		currentElf := elves[i]

		if currentElf.presentCount == 0 {
			i++
			continue
		}

		var nextElfIdx int
		var nextElf *elf
		nextElfStartIndex := i
		for {
			nextElfIdx = getNextElfPart1(elves, nextElfStartIndex)
			if nextElfIdx == i {
				return nextElfIdx + 1
			}

			nextElf = elves[nextElfIdx]
			if nextElf.presentCount == 0 {
				nextElfStartIndex++
				continue
			}

			break
		}

		currentElf.presentCount += nextElf.presentCount
		nextElf.presentCount = 0
		i++
	}
}

func getNextElfPart1(elves []*elf, startIndex int) int {
	for i := startIndex + 1; i != startIndex; i++ {
		if i >= len(elves) {
			i = -1
			continue
		}

		if elves[i].presentCount > 0 {
			return i
		}
	}

	return startIndex
}

type elfNode struct {
	elf
	next *elfNode
	prev *elfNode
}

func (n *elfNode) delete() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func runPart2(elfCount int) int {
	startNode, acrossNode := createElfNodes(elfCount)

	currentNode := startNode
	for i := 0; i < elfCount; i++ {
		if currentNode.next == currentNode {
			return currentNode.nr
		}

		currentNode.elf.presentCount += acrossNode.elf.presentCount
		acrossNode.presentCount = 0
		acrossNode.delete()

		acrossNode = acrossNode.next

		if (elfCount-i)%2 == 1 {
			acrossNode = acrossNode.next
		}

		currentNode = currentNode.next
	}

	return -1
}

func createElfNodes(elfCount int) (startNode *elfNode, acrossNode *elfNode) {
	acrossNodeIndex := elfCount / 2
	var prevNode *elfNode
	for i := 0; i < elfCount; i++ {
		currentNode := &elfNode{elf: elf{nr: i + 1, presentCount: 1}}
		currentNode.prev = prevNode

		if prevNode != nil {
			prevNode.next = currentNode
		}

		if i == 0 {
			startNode = currentNode
		}

		if i == acrossNodeIndex {
			acrossNode = currentNode
		}

		prevNode = currentNode
	}

	startNode.prev = prevNode
	prevNode.next = startNode

	return startNode, acrossNode
}
