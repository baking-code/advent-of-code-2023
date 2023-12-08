package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type tuple struct {
	L string
	R string
}

func mainold() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0

	reg := regexp.MustCompile("[A-Z]+")

	directions := []string{}
	bigMap := map[string]tuple{}
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == 1 {
			directions = strings.Split(line, "")
		} else if line != "" {
			split := strings.Split(line, " = ")
			key := split[0]
			result := reg.FindAll([]byte(split[1]), -1)
			bigMap[key] = tuple{L: string(result[0]), R: string(result[1])}
		}
	}
	result := "AAA"
	count := 0
	for result != "ZZZ" {
		for _, char := range directions {
			fmt.Print("moving from   ", result, "    ->>> ")
			step := string(char) // R or L
			if step == "L" {
				result = bigMap[result].L
			} else if step == "R" {
				result = bigMap[result].R
			} else {
				panic("ehlp")
			}
			fmt.Println("to  ", result)
			count++
		}
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
