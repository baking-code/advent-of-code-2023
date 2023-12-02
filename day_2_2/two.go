package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
		rounds := strings.Split(line, ";")
		gameMatch := gameRegex.FindStringSubmatch(line)
		gameNumber, err := strconv.ParseInt(string(gameMatch[1]), 10, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		isValid := true
		var minMap = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, round := range rounds {
			cubeMatch := cubeRegex.FindAllStringSubmatch(round, -1)
			for _, match := range cubeMatch {
				n, colour := match[1], match[2]
				num, err := strconv.ParseInt(string(n), 10, 0)

				// fmt.Println(num, colour, int64(numMap[colour]))
				minMap[colour] = int(math.Max(float64(minMap[colour]), float64(num)))
				if err != nil {
					isValid = false
					break
				}
			}
		}
		lineTotal := 1
		for _, v := range minMap {
			lineTotal *= v
		}
		fmt.Println("Game", gameNumber, "total", lineTotal, "map", minMap)
		if isValid {
			total += int(lineTotal)
		}
	}
	fmt.Println("game total", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
