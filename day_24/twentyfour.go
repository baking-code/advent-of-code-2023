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

type Hailstones struct {
	x0, y0, x, y float64
}

func solver(h1, h2 Hailstones, minLim, maxLim float64) bool {
	// trying to solve
	// f1(x) = a1*x + b1 = y
	// f2(x) = a2*x + b2 = y

	// check for perpendicular lines
	if h1.x0 == 0 {
		//fmt.Println("perpendicular x - ", h1, h2)
		return false
	}
	if h2.x0 == 0 {
		//fmt.Println("perpendicular y - ", h1, h2)
		return false
	}
	grad1 := (h1.y) / h1.x
	grad2 := (h2.y) / h2.x

	if grad1 == grad2 {
		// lines are parallel so will never intercept
		//fmt.Println("parallel - ", h1, h2)
		return false
	}

	b1 := h1.y0 - grad1*h1.x0
	b2 := h2.y0 - grad2*h2.x0

	xa := (b2 - b1) / (grad1 - grad2)

	ya := grad1*xa + b1
	if xa < minLim || xa >= maxLim {
		//fmt.Println("out of bounds x - ", h1, h2)
		return false
	}
	if ya < minLim || ya >= maxLim {
		//fmt.Println("out of bounds y - ", h1, h2)
		return false
	}

	// test for past crossing

	if xa < h1.x0 != (h1.x <= 0) || ya < h1.y0 != (h1.y <= 0) {
		//fmt.Println("hailstone a crosses in past - ", h1, h2)
		return false
	}
	if xa < h2.x0 != (h2.x <= 0) && ya < h2.y0 != (h2.y <= 0) {
		//fmt.Println("hailstone b crosses in past - ", h1, h2)
		return false
	}

	//fmt.Println("crosses inside area - ", h1, h2)

	return true
}

func main() {
	// lines := readFile("./test.txt")
	lines := readFile("./data.txt")
	hailstones := []Hailstones{}
	for _, l := range lines {
		tuple := strings.Split(l, "@")
		starts := strsToNumbers(strings.Split(tuple[0], ","))
		speeds := strsToNumbers(strings.Split(tuple[1], ","))
		hailstones = append(hailstones, Hailstones{x0: float64(starts[0]), y0: float64(starts[1]), x: float64(speeds[0]), y: float64(speeds[1])})
	}
	total := 0
	var solve = func(h1, h2 Hailstones) bool {
		// return solver(h1, h2, 7, 27)
		return solver(h1, h2, 200000000000000, 400000000000000)
	}
	for i, _ := range hailstones {
		if i < len(hailstones)-1 {
			for j := i + 1; j < len(hailstones); j++ {
				res := solve(hailstones[i], hailstones[j])
				fmt.Println(hailstones[i], hailstones[j], res)
				if res {
					total++
				}
			}
		}
	}
	fmt.Println("PT1", total)
}

var strsToNumbers = func(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		i, _ := strconv.Atoi(strings.TrimSpace(str))
		ints = append(ints, i)
	}
	return ints
}
