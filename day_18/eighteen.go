package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dirMap = map[string]image.Point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
	"0": {1, 0},
	"1": {0, 1},
	"2": {-1, 0},
	"3": {0, -1},
}

type Instruction struct {
	direction    string
	number       int64
	hexNumber    int64
	hexDirection string
}

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
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
	instructions := []Instruction{}
	hexCapture := regexp.MustCompile(`\(#(.*)\)`)
	run := func(useHex bool) int {
		base := 10
		if useHex {
			base = 16
		}
		for _, line := range lines {
			split := strings.Split(line, " ")
			number, _ := strconv.ParseInt(string(split[1]), base, strconv.IntSize)
			hex := hexCapture.FindStringSubmatch(split[2])
			hexNumber, _ := strconv.ParseInt(hex[1][:5], base, strconv.IntSize)
			hexDir := hex[1][5:]
			instructions = append(instructions, Instruction{number: number, direction: split[0], hexDirection: hexDir, hexNumber: hexNumber})
		}
		// image package gives us lovely cartesian manipulation
		oldPlot := image.Point{0, 0}
		totalArea := 0
		for _, ins := range instructions {
			len := ins.number
			move := dirMap[ins.direction]
			if useHex {
				len = ins.hexNumber
				move = dirMap[ins.hexDirection]
			}
			// add the direction (multiplied by the length) to get the next plot
			newPlot := oldPlot.Add(move.Mul(int(len)))
			// shoelace - cross-product of each sequential plots (thankfully just squares here)
			totalArea += oldPlot.X*newPlot.Y - oldPlot.Y*newPlot.X + int(len)
			oldPlot = newPlot
		}
		return totalArea/2 + 1
	}
	fmt.Println("Pt1", run(false))
	fmt.Println("Pt2", run(true))

}
