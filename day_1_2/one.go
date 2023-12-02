package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parser(input string) int64 {
	parsed, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		parsed = int64(numMap[input])
	}
	return parsed
}

func concatNumbers(left int64, right int64) int64 {
	leftString, rightString := fmt.Sprintf("%d", left), fmt.Sprintf("%d", right)
	concatString := leftString + rightString
	parsed, err := strconv.ParseInt(concatString, 10, 0)
	if err == nil {
		return parsed
	} else {
		fmt.Print(err)
		return 0
	}
}

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	file, err := os.Open("../day_1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digitStrings := strings.Join(maps.Keys(numMap), "|")
	digitRegex := regexp.MustCompile(fmt.Sprintf("\\d|%s", digitStrings))
	// sevennine should become 79, so we need to either overlap matches
	// or in this case, reverse the string so we can capture the first (last)
	// instance of the reversed digit
	digitRegexRevserse := regexp.MustCompile(fmt.Sprintf("\\d|%s", ReverseString(digitStrings)))
	total := int64(0)
	count := 0
	for scanner.Scan() {
		input := scanner.Text()
		numberFirst := digitRegex.FindString(input)
		numberLast := ReverseString(digitRegexRevserse.FindString(ReverseString(input)))
		if numberLast == "" {
			parsed := parser(numberFirst)
			if err == nil {
				code := concatNumbers(parsed, parsed)
				total += code
				count++
			} else {
				fmt.Print(numberFirst, err)
			}
		} else {
			parsedFirst := parser(numberFirst)
			parsedLast := parser(numberLast)
			if err == nil {
				code := concatNumbers(parsedFirst, parsedLast)
				total += code
				count++
			} else {
				fmt.Print(numberFirst, err)
			}
		}
	}
	fmt.Println("Total", total)
	fmt.Println("Count", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
