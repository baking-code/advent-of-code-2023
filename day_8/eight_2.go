package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func endsWithA(s string) bool {
	return s[len(s)-1:] == "A"
}

func endsWithZ(s string) bool {
	return s[len(s)-1:] == "Z"
}

func allEndWithZ(strs []string) bool {
	for _, v := range strs {
		if !endsWithZ(v) {
			return false
		}
	}
	return true
}

type tuple struct {
	L string
	R string
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0

	reg := regexp.MustCompile("[0-9A-Z]+")

	directions := []string{}
	bigMap := map[string]tuple{}
	result := []string{}
	ends := []string{}
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == 1 {
			directions = strings.Split(line, "")
		} else if line != "" {
			split := strings.Split(line, " = ")
			key := split[0]
			if endsWithA(key) {
				result = append(result, key)
			} else if endsWithZ(key) {
				ends = append(ends, key)
			}
			steps := reg.FindAll([]byte(split[1]), -1)
			bigMap[key] = tuple{L: string(steps[0]), R: string(steps[1])}
		}
	}
	fmt.Println(result)
	fmt.Println(ends)
	// fmt.Println(bigMap)

	// result := "AAA"
	count := 0
	lowestCounts := make([]int, len(result))
	gotEmAll := func() bool {
		for _, v := range lowestCounts {
			if v == 0 {
				return true
			}
		}
		return false
	}
	for gotEmAll() {
		for _, char := range directions {
			for i, each := range result {
				if endsWithZ(each) {
					// capture when we hit the desired destination
					lowestCounts[i] = count
				}
				// fmt.Print("moving from   ", each, "    ->>> ")
				step := string(char) // R or L
				if step == "L" {
					each = bigMap[each].L
				} else if step == "R" {
					each = bigMap[each].R
				} else {
					panic("ehlp")
				}
				// fmt.Println("to  ", each)
				result[i] = each
			}
			count++
		}
	}

	lcm := getLeastCommonMultiple(lowestCounts)
	fmt.Println(count, lowestCounts, lcm)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func greatestCommonDivisor(i, j int) int {
	// keep looking for the remainder until we get an exact division
	for j != 0 {
		temp := j
		j = i % j
		i = temp
	}
	return i
}

func getLeastCommonMultiple(numbers []int) int {
	first, second := numbers[0], numbers[1]
	result := leastCommonMultiple(first, second)
	rest := numbers[2:]
	for i := 0; i < len(rest); i++ {
		result = leastCommonMultiple(result, rest[i])
	}

	return result
}

func leastCommonMultiple(first, second int) int {
	// find the LCM between two numbers
	return first * second / greatestCommonDivisor(first, second)
}
