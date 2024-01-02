package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Matrix [][]rune
type Direction struct{ x, y int }
type Coordinate struct{ x, y int }

type Move struct {
	x              int
	y              int
	dir            Direction
	numConsecutive int
}

var (
	UP    = Direction{x: 0, y: -1}
	RIGHT = Direction{x: 1, y: 0}
	DOWN  = Direction{x: 0, y: 1}
	LEFT  = Direction{x: -1, y: 0}
)

var directions = [4]Direction{UP, DOWN, LEFT, RIGHT}

func (m Matrix) String() string {
	str := ""
	for _, row := range m {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}
func (m Matrix) InBounds(c Coordinate) bool {
	// width := len(m)
	// depth := len(m[0])
	// return Abs(c.x%width) >= 0 && Abs(c.x%width) < width && Abs(c.y%depth) >= 0 && Abs(c.y%depth) < depth
	return true
}
func (m Matrix) IsRock(c Coordinate) bool {
	width := len(m)
	depth := len(m[0])
	y := c.y % depth
	if c.y < 0 {
		y = depth + c.y%depth - 1
	}
	x := c.x % width
	if c.x < 0 {
		x = width + c.x%width - 1
	}
	return m[y][x] == '#'
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func move(c Coordinate, d Direction) Coordinate {
	return Coordinate{x: c.x + d.x, y: c.y + d.y}
}

func (current Direction) isReverse(next Direction) bool {
	switch next {
	case RIGHT:
		return current == LEFT
	case LEFT:
		return current == RIGHT
	case UP:
		return current == DOWN
	case DOWN:
		return current == UP
	default:
		return false
	}
}

func (m Matrix) walk(startingPoint Coordinate, numSteps int) int {

	// track with map to behave like a set
	steps := map[Coordinate]bool{startingPoint: true}
	for i := 0; i < numSteps; i++ {
		newSteps := map[Coordinate]bool{}
		for step := range steps {
			for _, d := range directions {
				next := move(step, d)
				isNextRock := m.IsRock(next)
				isNextInBounds := m.InBounds(next)
				if !isNextRock && isNextInBounds {
					newSteps[next] = true
				}
			}
		}
		// fmt.Println(newSteps)
		steps = newSteps
	}
	return len(steps)
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
	matrix := Matrix{}
	start := Coordinate{}
	for _, line := range lines {
		row := []rune{}
		// split string into chars so we can form a 2x2 matrix and inspect the columns
		for i, char := range line {
			if char == 'S' {
				start = Coordinate{x: i, y: len(matrix)}
			}
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	fmt.Println("Pt1 6 steps", matrix.walk(start, 6))
	fmt.Println("Pt2 10 steps", matrix.walk(start, 10))
	fmt.Println("Pt2 50 steps", matrix.walk(start, 50))
	fmt.Println("Pt2 64 steps", matrix.walk(start, 64))
	fmt.Println("Pt2 100 steps", matrix.walk(start, 100))
	// fmt.Println("Pt2 26501365 steps", matrix.walk(start, 26501365))

}
