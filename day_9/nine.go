package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/baking-code/godash/every"
)

type tuple struct {
	L string
	R string
}

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0

	nextTotal := 0
	prevTotal := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		seq := getNumbersFromString(line)
		stop := false
		nextDiffs := seq
		lasts := []int{}
		firsts := []int{}
		lastTotal := 0
		firstTotal := 0
		// fmt.Println("initial", seq)
		for !stop {
			newDiff := []int{}
			for i := range nextDiffs {
				if i+1 < len(nextDiffs) {
					newDiff = append(newDiff, nextDiffs[i+1]-nextDiffs[i])
				}
			}
			if every.Every[int](newDiff, func(e int, index int, collection []int) bool { return e == 0 }) {
				// when we're done, add each "lasts" up in sequence so we can get our extrapolate
				for _, last := range lasts {
					lastTotal += last
				}
				for _, first := range firsts {
					firstTotal += first
				}
				fmt.Println("last", seq[len(seq)-1], lastTotal)
				nextTotal += seq[len(seq)-1] + lastTotal
				fmt.Println("first", seq[len(seq)-1], firstTotal)
				prevTotal += seq[0] + firstTotal
				stop = true
			} else {
				lasts = append(lasts, newDiff[len(newDiff)-1])
				firsts = append(firsts, newDiff[0])
				nextDiffs = newDiff
			}
		}

	}

	fmt.Println("next", nextTotal, "previous", prevTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumbersFromString(s string) []int {
	strings := strings.Split(s, " ")
	numbers := make([]int, 0)
	for _, str := range strings {
		if str != " " && str != "" {
			num, err := strconv.ParseInt(str, 10, 0)
			if err != nil {
				fmt.Println(err)
			} else {
				numbers = append(numbers, int(num))
			}
		}
	}
	return numbers
}
