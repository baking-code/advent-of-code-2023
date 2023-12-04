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
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	symbolsRegex := regexp.MustCompile("[-!#$%@^&*()_+|~=`{}\\[\\]:\";'<>?,\\/]")
	digits := regexp.MustCompile("(\\d+)")
	lineNumber := 0
	var symbolMap = map[int][]int{}
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		symbolsIndexes := symbolsRegex.FindAllIndex([]byte(line), -1)
		var arr = make([]int, 0)
		for _, hit := range symbolsIndexes {
			arr = append(arr, hit[0])
		}
		symbolMap[lineNumber] = arr
		lineNumber++
	}
	var arr = make([]int64, 0)
	for lineNumber, line := range lines {
		ints := digits.FindAllIndex([]byte(line), -1)
		for _, hit := range ints {
			digitStart := hit[0]
			digitEnd := hit[1]
			digit := line[digitStart:digitEnd]
			// check current line
			isHitRow := isSymbolHit(symbolMap[lineNumber], digitStart, digitEnd, false)
			//check prev line
			isHitPrevious := isSymbolHit(symbolMap[lineNumber-1], digitStart, digitEnd, true)
			// check next line
			isHitNext := isSymbolHit(symbolMap[lineNumber+1], digitStart, digitEnd, true)

			if isHitRow || isHitNext || isHitPrevious {
				letsgo, err := strconv.ParseInt(digit, 10, 0)
				if err == nil {
					arr = append(arr, letsgo)
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

func isSymbolHit(symbolsIndexes []int, digitStart int, digitEnd int, checkSame bool) bool {
	fmt.Println(symbolsIndexes, digitStart, digitEnd)

	if symbolsIndexes != nil {
		for _, symbolIndex := range symbolsIndexes {
			if symbolIndex == digitStart-1 || symbolIndex == digitEnd {
				return true
			}
			if checkSame {
				if digitStart <= symbolIndex && symbolIndex <= digitEnd {
					return true
				}
			}
		}
	}
	return false
}
