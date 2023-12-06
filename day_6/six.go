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
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0
	times := []int{}
	distances := []int{}
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == 1 {
			// gather the times
			times = getNumbersFromString(strings.Split(line, ":")[1])

		} else if lineNo == 2 {
			distances = getNumbersFromString(strings.Split(line, ":")[1])
		}
	}

	total := 1
	for race := range times {
		time, distance := times[race], distances[race]
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
	}

	fmt.Println("total", total)

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
