package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../day_6/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0
	time := 0
	distance := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == 1 {
			// gather the times
			time = getNumberFromString(strings.Split(line, ":")[1])

		} else if lineNo == 2 {
			distance = getNumberFromString(strings.Split(line, ":")[1])
		}
	}

	total := 1
	winningRaces := 0
	for timePressed := 0; timePressed < time; timePressed++ {
		speed := timePressed
		distanceReached := speed * (time - timePressed)
		// fmt.Println("race", race, "speed", speed, "distance", distanceReached, "raceLength", distance)
		if distanceReached > distance {
			winningRaces++
		}
	}
	total *= winningRaces

	fmt.Println("total", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumberFromString(s string) int {
	strings := strings.Split(s, " ")
	numberString := ""
	for _, str := range strings {
		if str != " " && str != "" {
			numberString += str
		}
	}
	num, err := strconv.ParseInt(numberString, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	return int(num)
}
