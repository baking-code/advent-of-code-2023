package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Matrix [][]string

type Coord struct {
	x, y int
}

func (m Matrix) String() string {
	str := ""
	for _, row := range m {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}
func (m Matrix) getSymbol(c Coord) string {
	return m[c.x][c.y]
}

func main() {
	matrix := Matrix{}
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0

	var startingPoint = Coord{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for x, s := range line {
			if string(s) == "S" {
				startingPoint = Coord{x, y}
			}
			row = append(row, string(s))
		}
		matrix = append(matrix, row)
		y++
	}

	fmt.Println(matrix)
	fmt.Println("starting at", startingPoint)
	loop, err := findLoop(matrix, startingPoint)
	// answer is the largest length of loop / 2
	fmt.Println("PT1", len(loop)/2)
	enclosed := findArea(matrix, loop)
	fmt.Println("PT2", enclosed)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findArea(m Matrix, path []Coord) int {
	num := 0
	size := len(m)
	isPath := func(c Coord) bool {
		for _, pathEl := range path {
			if pathEl.x == c.x && pathEl.y == c.y {
				return true
			}
		}
		return false
	}

	isContained := func(c Coord) bool {
		// extrapolate to each edge to see if we hit our path
		// if it's got a path around it in each direction, then it's contained
		xp := c.x
		inPathA := false
		inPathB := false
		inPathC := false
		inPathD := false
		for {
			xp++
			if isPath(Coord{xp, c.y}) {
				inPathA = true
				break
			} else if xp > size {
				break
			}
			// fmt.Println("xp", xp)
		}
		xn := c.x
		for {
			xn--
			if isPath(Coord{xn, c.y}) {
				inPathB = true
				break
			} else if xn < 0 {
				break
			}
			// fmt.Println("xn", xn)
		}
		yp := c.y
		for {
			yp++
			if isPath(Coord{c.x, yp}) {
				inPathC = true
				break
			} else if yp > size {
				break
			}
			// fmt.Println("yp", yp)
		}
		yn := c.y
		for {
			yn--
			if isPath(Coord{c.x, yn}) {
				inPathD = true
				break
			} else if yn < 0 {
				break
			}
			// fmt.Println("yn", yn)
		}
		return inPathA && inPathB && inPathC && inPathD
	}
	for dx := range m {
		for dy := range m[dx] {
			coord := Coord{x: dx, y: dy}
			if !isPath(coord) {
				if isContained(coord) {
					num++
				}
			}
		}
	}
	return num
}

func findLoop(m Matrix, startingPosition Coord) ([]Coord, error) {
	visited := make([][]bool, len(m))
	for i := range visited {
		visited[i] = make([]bool, len(m[i]))
	}

	loop := []Coord{startingPosition}
	visited[startingPosition.x][startingPosition.y] = true

	for {
		// find valid surrounding coords
		surrounding := findSurrounding(m, loop[len(loop)-1])
		// disregard ones we've already visited
		for len(surrounding) > 0 && visited[surrounding[0].x][surrounding[0].y] {
			surrounding = surrounding[1:]
		}

		// if we've seen everything, we're done
		if len(surrounding) == 0 {
			fmt.Println("Back at pos", startingPosition, loop[len(loop)-1])
			loop = append(loop, startingPosition)
			break
		}

		// for a valid loop there's only one possible outcome
		loop = append(loop, surrounding[0])
		// mark this as visited so we don't go back on ourselves
		visited[surrounding[0].x][surrounding[0].y] = true
	}

	return loop, nil
}

func findSurrounding(m Matrix, c Coord) []Coord {
	surrounding := []Coord{}

	shape := m.getSymbol(c)

	switch shape {
	case "S":
		surrounding = findStartingPositionPossibles(m, c)
	case "|":
		// for each valid symbol, check the coordinate is within the matrix bounds
		if c.x > 0 {
			surrounding = append(surrounding, Coord{c.x - 1, c.y})
		}
		if c.x < len(m)-1 {
			surrounding = append(surrounding, Coord{c.x + 1, c.y})
		}
	case "-":
		if c.y > 0 {
			surrounding = append(surrounding, Coord{c.x, c.y - 1})
		}
		if c.y < len(m[c.x])-1 {
			surrounding = append(surrounding, Coord{c.x, c.y + 1})
		}
	case "L":
		if c.x > 0 {
			surrounding = append(surrounding, Coord{c.x - 1, c.y})
		}
		if c.y < len(m[c.x])-1 {
			surrounding = append(surrounding, Coord{c.x, c.y + 1})
		}
	case "J":
		if c.x > 0 {
			surrounding = append(surrounding, Coord{c.x - 1, c.y})
		}
		if c.y > 0 {
			surrounding = append(surrounding, Coord{c.x, c.y - 1})
		}
	case "7":
		if c.x < len(m)-1 {
			surrounding = append(surrounding, Coord{c.x + 1, c.y})
		}
		if c.y > 0 {
			surrounding = append(surrounding, Coord{c.x, c.y - 1})
		}
	case "F":
		if c.x < len(m)-1 {
			surrounding = append(surrounding, Coord{c.x + 1, c.y})
		}
		if c.y < len(m[c.x])-1 {
			surrounding = append(surrounding, Coord{c.x, c.y + 1})
		}
	}

	return surrounding
}

func findStartingPositionPossibles(m Matrix, pos Coord) []Coord {
	surrounding := []Coord{}

	if pos.x > 0 && slices.Contains([]string{"|", "F", "7"}, m[pos.x-1][pos.y]) {
		surrounding = append(surrounding, Coord{pos.x - 1, pos.y})
	}
	if pos.x < len(m)-1 && slices.Contains([]string{"|", "L", "J"}, m[pos.x+1][pos.y]) {
		surrounding = append(surrounding, Coord{pos.x + 1, pos.y})
	}
	if pos.y > 0 && slices.Contains([]string{"-", "L", "F"}, m[pos.x][pos.y-1]) {
		surrounding = append(surrounding, Coord{pos.x, pos.y - 1})
	}
	if pos.y < len(m[pos.x])-1 && slices.Contains([]string{"-", "J", "7"}, m[pos.x][pos.y+1]) {
		surrounding = append(surrounding, Coord{pos.x, pos.y + 1})
	}

	return surrounding
}
