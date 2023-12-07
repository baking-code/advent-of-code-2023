package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/baking-code/godash/some"
)

func main() {
	file, err := os.Open("../day_4/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameCopiesMap := map[int]int{}
	totalScore := 0
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		gameCopiesMap[lineNo]++
		numberOfWinners := getNumberOfWinningNumbers(line)
		// for each copy we've collected
		for j := 0; j < gameCopiesMap[lineNo]; j++ {
			// add the copies for any subsequent winning cards
			for i := 1; i < numberOfWinners+1; i++ {
				gameCopiesMap[lineNo+i]++
			}
		}
	}

	for _, copies := range gameCopiesMap {
		totalScore += copies
	}

	fmt.Println("Total score", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumberOfWinningNumbers(line string) int {
	colonIndex := strings.Index(line, ":")
	allNumbers := line[colonIndex+1:]
	digits := strings.Split(allNumbers, "|")
	winningString, playString := digits[0], digits[1]
	winningNumbers, playNumbers := splitNumbers(winningString), splitNumbers(playString)
	numberOfWinners := 0
	for _, play := range playNumbers {

		if some.Some[int64](winningNumbers, func(e int64, index int, collection []int64) bool {
			return e == play
		}) {
			numberOfWinners++
		}
	}
	return numberOfWinners
}

func splitNumbers(s string) []int64 {
	strings := strings.Split(s, " ")
	numbers := make([]int64, 0)
	for _, str := range strings {
		if str != " " && str != "" {
			num, err := strconv.ParseInt(str, 10, 0)
			if err != nil {
				fmt.Println(err)
			} else {
				numbers = append(numbers, num)
			}
		}
	}
	return numbers
}
