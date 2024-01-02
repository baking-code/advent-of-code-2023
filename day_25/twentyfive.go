package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	lines := readFile("./test.txt")
	// lines := readFile("./data.txt")

	total := 0
	// groups := [][]string{}
	connections := map[string][]string{}
	for _, line := range lines {
		spl := strings.Split(line, ":")
		left := spl[0]
		right := strings.Split(spl[1], " ")
		for _, conn := range right {
			if conn != "" {
				item := strings.Trim(conn, " ")
				connections[left] = append(connections[left], item)
				connections[item] = append(connections[item], left)
			}
		}
	}
	inverse := map[string][]string{}
	for key, values := range connections {
		for _, v := range values {
			inverse[v] = append(inverse[v], key)
		}
	}
	fmt.Println("PT1", total)
	fmt.Println("conns", connections)
	// fmt.Println("inv", inverse)
}

var strsToNumbers = func(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		i, _ := strconv.Atoi(strings.TrimSpace(str))
		ints = append(ints, i)
	}
	return ints
}
