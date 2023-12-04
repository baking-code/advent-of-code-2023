package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/baking-code/godash"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// digits := regexp.MustCompile("(\\d+)")

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		allNumbers := line[colonIndex+1:]
		digits := strings.Split(allNumbers, "|")
		winningString, playString := digits[0], digits[1]
		winningNumbers, playNumbers := splitNumbers(winningString), splitNumbers(playString)
		numberOfWinners := 0
		for _, play := range playNumbers {

			if godash.Some[int64](winningNumbers, func(e int64, index int, collection []int64) bool {
				return e == play
			}) {
				numberOfWinners++
			}
		}
		totalScore += int(math.Pow(2.0, float64(numberOfWinners)))
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
		if str != " " {
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
