package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Lens struct {
	label  string
	number int64
}

type Box map[int][]Lens

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = strings.Split(line, ",")
	}
	return lines
}

var typeRegex = regexp.MustCompile("[=-]")

func main() {
	boxes := Box{}
	var indexOfLens = func(labelHash int, label string) int {
		index := -1
		for i, lens := range boxes[labelHash] {
			if lens.label == label {
				index = i
				break
			}
		}
		return index
	}
	sequence := readFile("./data.txt")
	pt1Total := 0
	pt2Total := 0
	for _, step := range sequence {
		pt1Total += Hash(step)
		// Part 2
		// split up the step
		loc := typeRegex.FindIndex([]byte(step))
		label, sign := string(step[:loc[0]]), string(step[loc[0]])
		labelHash := Hash(label)
		// fmt.Println(label, sign, labelHash)
		if sign == "=" {
			number, err := strconv.ParseInt(string(step[loc[1]:]), 10, 0)
			if err == nil {
				// find index of any existing lenses with this label
				index := indexOfLens(labelHash, label)
				if index > -1 {
					boxes[labelHash][index] = Lens{label, number}
				} else {
					boxes[labelHash] = append(boxes[labelHash], Lens{label, number})
				}
			}
		} else if sign == "-" {
			index := indexOfLens(labelHash, label)
			if index > -1 {
				boxes[labelHash] = remove(boxes[labelHash], index)
			}
		}
	}
	for k, box := range boxes {
		boxNumber := k + 1
		for ind, lens := range box {
			slotNumber := ind + 1
			slotTotal := boxNumber * slotNumber * int(lens.number)
			pt2Total += slotTotal
		}
	}
	fmt.Println("pt1:", pt1Total)
	fmt.Println("pt2:", pt2Total)
}

func Hash(s string) int {
	currentValue := 0
	for _, char := range s {
		code := int(char)
		currentValue += code
		currentValue *= 17
		currentValue = currentValue % 256
	}
	return currentValue
}

func remove(slice []Lens, index int) []Lens {
	return append(slice[:index], slice[index+1:]...)
}
