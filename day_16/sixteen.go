package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Matrix [][]string
type Dir int

// x,y and direction test
type Visited [][][4]bool

const (
	UP    = Dir(0)
	RIGHT = Dir(1)
	DOWN  = Dir(2)
	LEFT  = Dir(3)
)

func (m Matrix) String() string {
	str := ""
	for _, row := range m {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}

func (m Matrix) Energized(startX, startY int, d Dir) int {
	v := make(Visited, len(m))
	for y := range v {
		// set up each visitor
		v[y] = make([][4]bool, len(m[y]))
	}
	return m.visit(startX, startY, d, v)
}

func (m Matrix) visit(x, y int, d Dir, v Visited) (visited int) {
	width := len(m)
	// out of bounds?
	if x < 0 || y < 0 || y >= width || x >= len(m[y]) {
		return
	}
	// already visited?
	if v[y][x][d] {
		return
	}

	visited++
	// reset directions
	for _, dir := range v[y][x] {
		if dir {
			visited--
			break
		}
	}
	v[y][x][d] = true

	switch m[y][x] {
	case "-":
		if d == UP || d == DOWN {
			visited += m.visit(x-1, y, LEFT, v)
			visited += m.visit(x+1, y, RIGHT, v)
			return
		}
	case "|":
		if d == LEFT || d == RIGHT {
			visited += m.visit(x, y-1, UP, v)
			visited += m.visit(x, y+1, DOWN, v)
			return
		}
	case "\\":
		switch d {
		case UP:
			d = LEFT
		case RIGHT:
			d = DOWN
		case DOWN:
			d = RIGHT
		case LEFT:
			d = UP
		}
	case "/":
		switch d {
		case UP:
			d = RIGHT
		case RIGHT:
			d = UP
		case DOWN:
			d = LEFT
		case LEFT:
			d = DOWN
		}
	}

	switch d {
	case UP:
		y--
	case RIGHT:
		x++
	case DOWN:
		y++
	case LEFT:
		x--
	}

	// visit the next one
	return visited + m.visit(x, y, d, v)
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
	lines := readFile("./data.txt")
	matrix := Matrix{}
	for _, line := range lines {
		row := []string{}
		// split string into chars so we can form a 2x2 matrix and inspect the columns
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}

	fmt.Println("Pt1", matrix.Energized(0, 0, RIGHT))

	width := len(matrix)
	depth := len(matrix[0])
	// 2X so that we can test 2 directions
	totals := make([]int, width*2+depth*2)
	wg := sync.WaitGroup{}
	wg.Add(len(totals))

	count := func(x, y, i int, d Dir) {
		totals[i] = matrix.Energized(x, y, d)
		wg.Done()
	}

	i := 0
	// try each column with first (->) and last (<-) row
	for y := 0; y < width; y++ {
		go count(0, y, i, RIGHT)
		i++
		go count(len(matrix[y])-1, y, i, LEFT)
		i++
	}
	// try each row from first (V) and last (^) column
	for x := 0; x < depth; x++ {
		go count(x, 0, i, DOWN)
		i++
		go count(x, len(matrix)-1, i, UP)
		i++
	}
	// wait for all goroutines to finish
	wg.Wait()

	max := 0
	for _, total := range totals {
		if total > max {
			max = total
		}
	}

	fmt.Println("Pt2", max)
}
