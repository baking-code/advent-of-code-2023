package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	part    string
	symbol  string
	value   int
	outcome string
}

type WorkflowMap map[string][]Instruction

// min and max values for each letter of xmas
type Bracket map[string][2]int

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	lines := readFile("./data.txt")
	instructionRegex := regexp.MustCompile(`(([xmas])([<>])(\d+):)?([a-zA-Z]+)?`)
	workflowMap := map[string][]Instruction{}
	flip := false
	parts := []map[string]int{}
	for _, line := range lines {
		if line == "" {
			flip = true
			continue
		}
		if !flip {
			extractInstructions(line, instructionRegex, workflowMap)
		} else {
			parts = extractParts(line, parts)
		}
	}

	total := 0
	for _, part := range parts {
		workflow := "in"
		for {
			if workflow == "R" || workflow == "A" {
				break
			}
			instructions, ok := workflowMap[workflow]
			if ok {
				for _, instruction := range instructions {
					if instruction.symbol != "" {
						val := part[instruction.part]
						if instruction.symbol == "<" {
							if val < instruction.value {
								workflow = instruction.outcome
								break
							} else {
								continue
							}
						} else {
							if val > instruction.value {
								workflow = instruction.outcome
								break
							} else {
								continue
							}
						}
					} else {
						workflow = instruction.outcome
					}
				}
			} else {
				panic("not okay")
			}
		}
		if workflow == "A" {
			for _, value := range part {
				total += value
			}
		}
	}
	fmt.Println("PT1", total)
	fmt.Println("PT2", countBrackets(workflowMap, "in", Bracket{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}))

}

func extractParts(line string, parts []map[string]int) []map[string]int {
	partsString := strings.Split(line[1:len(line)-1], ",")
	partMap := map[string]int{}
	for _, part := range partsString {
		split := strings.Split(part, "=")
		value, _ := strconv.ParseInt(split[1], 10, strconv.IntSize)
		partMap[split[0]] = int(value)
	}
	parts = append(parts, partMap)
	return parts
}

func extractInstructions(line string, instructionRegex *regexp.Regexp, workflowMap WorkflowMap) {
	parts := strings.Split(line, "{")
	workflow := parts[0]
	instructionsStrings := strings.Split(parts[1], ",")
	instructions := []Instruction{}
	for _, ins := range instructionsStrings {
		in := instructionRegex.FindAllStringSubmatch(ins, -1)
		matches := in[0]
		value, _ := strconv.ParseInt(matches[4], 10, strconv.IntSize)
		instruction := Instruction{part: matches[2], symbol: matches[3], value: int(value), outcome: matches[5]}
		instructions = append(instructions, instruction)
	}
	workflowMap[workflow] = instructions
}

func cloneBracket(original Bracket) Bracket {
	c := make(Bracket, 4)
	for k, v := range original {
		c[k] = v
	}
	return c
}

func countBrackets(workflows WorkflowMap, currentPart string, bracket Bracket) int {
	if currentPart == "R" {
		// stop
		return 0
	} else if currentPart == "A" {
		// if workflow reaches end, sum all cross products for this valid bracket
		total := 1
		for _, bracket := range bracket {
			total *= (bracket[1] - bracket[0] + 1)
		}
		return total
	}

	total := 0

	for _, instruction := range workflows[currentPart] {
		value := bracket[instruction.part]

		// track each side of the workflow - these will form the next bracket to consider
		truePath := [2]int{}
		falsePath := [2]int{}
		if instruction.symbol == "<" {
			// take bracket up until the true value
			truePath = [2]int{value[0], instruction.value - 1}
			// bracket for the rest
			falsePath = [2]int{instruction.value, value[1]}
		} else if instruction.symbol == ">" {
			truePath = [2]int{instruction.value + 1, value[1]}
			falsePath = [2]int{value[0], instruction.value}
		} else {
			// move to the next workflow
			total += countBrackets(workflows, instruction.outcome, bracket)
			continue
		}

		// create a copy of the bracket reduced to the new part value
		if truePath[0] <= truePath[1] {
			nextBracket := cloneBracket(bracket)
			nextBracket[instruction.part] = truePath
			// we. go. again.
			total += countBrackets(workflows, instruction.outcome, nextBracket)
		}

		// we can break early if the false path is reduced to nothing
		if falsePath[0] > falsePath[1] {
			break
		}

		// otherwise update bracket and move onto the next workflow
		bracket[instruction.part] = falsePath
	}
	return total
}
