package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameRegex := regexp.MustCompile("Game (\\d+)")
	cubeRegex := regexp.MustCompile("(\\d+) (red|blue|green)")
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameMatch := gameRegex.FindStringSubmatch(line)
		gameNumber, err := strconv.ParseInt(string(gameMatch[1]), 10, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		cubeMatch := cubeRegex.FindAllStringSubmatch(line, -1)
		isValid := true
		for _, match := range cubeMatch {
			n, colour := match[1], match[2]
			num, err := strconv.ParseInt(string(n), 10, 0)
			// fmt.Println(num, colour, int64(numMap[colour]))
			if err != nil || num > int64(numMap[colour]) {
				isValid = false
				break
			}
		}
		if isValid {
			total += int(gameNumber)
		}
	}
	fmt.Println("game total", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
