package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../day_3/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	symbolsRegex := regexp.MustCompile("[*]+")
	digits := regexp.MustCompile("(\\d+)")
	lineNumber := 0
	var lineDigitIndexMap = map[int][][]int{}
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		digitsIndexes := digits.FindAllIndex([]byte(line), -1)
		var arr = make([][]int, 0)
		for _, hit := range digitsIndexes {
			arr = append(arr, hit)
		}
		lineDigitIndexMap[lineNumber] = arr
		lineNumber++
	}
	var arr = make([]int64, 0)
	for lineNumber, line := range lines {
		symbols := symbolsRegex.FindAllIndex([]byte(line), -1)
		for _, hit := range symbols {
			if len(hit) > 0 {
				fmt.Println("WWWW", hit)
			}
			symbolIndex := hit[0]
			//check prev line
			isHitPrevious := isSymbolHit(lineDigitIndexMap[lineNumber-1], symbolIndex, false)
			// check next line
			isHitNext := isSymbolHit(lineDigitIndexMap[lineNumber+1], symbolIndex, false)
			// check this line
			isHit := isSymbolHit(lineDigitIndexMap[lineNumber], symbolIndex, true)
			if len(isHit) > 0 {
				// fmt.Println("HEY", len(isHitPrevious) > 0, len(isHitNext) > 0)
			}
			if len(isHitNext) > 0 && len(isHitPrevious) > 0 {
				digitOne, err := strconv.ParseInt(lines[lineNumber-1][isHitPrevious[0]:isHitPrevious[1]], 10, 0)
				digitTwo, err := strconv.ParseInt(lines[lineNumber+1][isHitNext[0]:isHitNext[1]], 10, 0)
				if err == nil {
					arr = append(arr, digitOne*digitTwo)
				} else {
					fmt.Println(err)
				}
			} else if len(isHit) > 0 && len(isHitPrevious) > 0 {
				digitOne, err := strconv.ParseInt(lines[lineNumber-1][isHitPrevious[0]:isHitPrevious[1]], 10, 0)
				digitTwo, err := strconv.ParseInt(lines[lineNumber][isHit[0]:isHit[1]], 10, 0)
				if err == nil {
					arr = append(arr, digitOne*digitTwo)
				} else {
					fmt.Println(err)
				}
			} else if len(isHit) > 0 && len(isHitNext) > 0 {
				digitOne, err := strconv.ParseInt(lines[lineNumber+1][isHitNext[0]:isHitNext[1]], 10, 0)
				digitTwo, err := strconv.ParseInt(lines[lineNumber][isHit[0]:isHit[1]], 10, 0)
				if err == nil {
					arr = append(arr, digitOne*digitTwo)
				} else {
					fmt.Println(err)
				}
			}
		}

	}
	total := int64(0)
	for _, num := range arr {
		total += num
	}
	fmt.Println("game total", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isSymbolHit(digitIndexes2d [][]int, symbolIndex int, sameLine bool) []int {
	if digitIndexes2d != nil {
		for _, digitIndexes := range digitIndexes2d {
			if !sameLine {
				if digitIndexes[0]-1 <= symbolIndex && symbolIndex <= digitIndexes[1] {
					return digitIndexes
				}
			} else {
				if digitIndexes[0]-1 == symbolIndex || symbolIndex == digitIndexes[1] {
					return digitIndexes
				}
			}
		}
	}

	return []int{}
}
