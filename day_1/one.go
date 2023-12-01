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
	digitRegex := regexp.MustCompile("\\d")

	total := int64(0)
	for scanner.Scan() {
		input := scanner.Text()
		numbers := digitRegex.FindAllString(input, -1)
		if len(numbers) == 0 {
			continue
		}
		if len(numbers) == 1 {
			parsed, err := strconv.ParseInt(numbers[0]+numbers[0], 10, 0)
			if err == nil {
				total += parsed
			} else {
				fmt.Print(numbers, err)
			}
		} else {
			parsedFirst, err := strconv.ParseInt(numbers[0]+numbers[len(numbers)-1], 10, 0)
			if err == nil {
				total += parsedFirst
			} else {
				fmt.Print(numbers, err)
			}
		}
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
