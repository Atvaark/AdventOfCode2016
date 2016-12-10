package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type bot struct {
	botID     int
	lowValue  int
	highValue int
	low       target
	high      target
}

type instruction struct {
	botID           int
	initialValue    int
	initialValueSet bool
	low             target
	high            target
}

type target struct {
	targetType targetType
	ID         int
}

type targetType int

const (
	noType targetType = iota
	botType
	outputType
)

type output struct {
	outputID int
	value    int
}

type comparison struct {
	botID     int
	lowValue  int
	highValue int
}

func main() {
	bots, err := openInput("input.txt")
	if err != nil {
		fmt.Printf("unable to open input: %v", err)
		os.Exit(1)
	}

	part1Result, part2Result := run(bots)
	fmt.Printf("part 1: %d\n", part1Result)
	fmt.Printf("part 2: %d\n", part2Result)
}

func run(bots map[int]*bot) (int, int) {
	outChan := make(chan output)
	cmpChan := make(chan comparison)
	doneChan := make(chan bool)
	go startBots(bots, outChan, cmpChan, doneChan)
	part1Result, part2Result := calcResults(outChan, cmpChan, doneChan)
	return part1Result, part2Result
}

func calcResults(outChan chan output, cmpChan chan comparison, doneChan chan bool) (int, int) {
	part1Result := -1
	part2Result := 1
	for {
		select {
		case o := <-outChan:
			if 0 <= o.outputID && o.outputID <= 2 {
				part2Result *= o.value
			}

		case c := <-cmpChan:
			if c.lowValue == 17 && c.highValue == 61 {
				part1Result = c.botID
			}
		case <-doneChan:
			return part1Result, part2Result
		}
	}
}

func openInput(name string) (map[int]*bot, error) {
	file, error := os.Open("input.txt")
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bots := make(map[int]*bot)
	for scanner.Scan() {
		line := scanner.Text()
		instr := parseInstruction(line)
		setupBot(&bots, instr)
	}

	return bots, nil
}

func addValue(b *bot, value int) {
	if value > b.highValue {
		if b.lowValue != -1 {
			fmt.Printf("error adding value to bot %d: no capacity", b.botID)
		}

		b.lowValue = b.highValue
		b.highValue = value
	} else {
		if b.lowValue != -1 {
			fmt.Printf("error adding value to bot %d: no capacity", b.botID)
		}

		b.lowValue = value
	}
}

func setupBot(bots *map[int]*bot, instr instruction) {
	b, ok := (*bots)[instr.botID]
	if !ok {
		b = &bot{botID: instr.botID, highValue: -1, lowValue: -1}
		(*bots)[b.botID] = b
	}

	if instr.initialValueSet {
		addValue(b, instr.initialValue)
	}

	if instr.low.targetType != noType {
		b.low = instr.low
	}

	if instr.high.targetType != noType {
		b.high = instr.high
	}

}

func parseInstruction(line string) instruction {
	var result instruction
	split := strings.Split(line, " ")
	switch split[0] {
	case "value":
		value, _ := strconv.Atoi(split[1])
		result.initialValue = value
		result.initialValueSet = true

		botID, _ := strconv.Atoi(split[5])
		result.botID = botID
	case "bot":
		botID, _ := strconv.Atoi(split[1])
		result.botID = botID

		lowTargetType := split[5]
		lowTargetID := split[6]

		highTargetType := split[10]
		highTargetID := split[11]

		parseTarget := func(targetType string, targetID string) target {
			var newTarget target
			switch targetType {
			case "output":
				newTarget.targetType = outputType
			case "bot":
				newTarget.targetType = botType
			default:
			}

			ID, _ := strconv.Atoi(targetID)
			newTarget.ID = ID
			return newTarget
		}

		result.low = parseTarget(lowTargetType, lowTargetID)
		result.high = parseTarget(highTargetType, highTargetID)
	default:
	}

	return result
}

func startBots(bots map[int]*bot, outChan chan output, cmpChan chan comparison, doneChan chan bool) {
	var botIDs []int
	for botID := range bots {
		botIDs = append(botIDs, botID)
	}

	sort.Ints(botIDs)

	for {
		botFound := false
		for botID := range botIDs {
			bot, _ := bots[botID]

			if bot.lowValue != -1 && bot.highValue != -1 {
				botFound = true

				sendChip := func(value int, t target) {
					switch t.targetType {
					case outputType:
						//fmt.Printf("bot %d sending %d to output %d\n", bot.botID, value, t.ID)
						outChan <- output{outputID: t.ID, value: value}
					case botType:
						//fmt.Printf("bot %d sending %d to bot %d\n", bot.botID, value, t.ID)
						targetBot, _ := bots[t.ID]
						addValue(targetBot, value)
					default:
					}
				}

				cmpChan <- comparison{botID: bot.botID, lowValue: bot.lowValue, highValue: bot.highValue}
				sendChip(bot.lowValue, bot.low)
				sendChip(bot.highValue, bot.high)

				bot.lowValue = -1
				bot.highValue = -1

				break
			}
		}

		if !botFound {
			break
		}

	}

	doneChan <- true
}
