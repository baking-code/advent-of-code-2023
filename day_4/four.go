package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/baking-code/godash/some"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// digits := regexp.MustCompile("(\\d+)")

	totalScore := 0
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
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
		gameScore := int(math.Pow(2.0, float64(numberOfWinners)-1))
		totalScore += gameScore
		fmt.Println("line", lineNo, "winners", numberOfWinners, "score", gameScore)
	}

	fmt.Println("Total score", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
