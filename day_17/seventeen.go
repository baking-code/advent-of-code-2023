package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Matrix [][]int
type Dir struct{ x, y int }

type Move struct {
	x              int
	y              int
	dir            Dir
	numConsecutive int
}

var (
	UP    = Dir{x: 0, y: -1}
	RIGHT = Dir{x: 1, y: 0}
	DOWN  = Dir{x: 0, y: 1}
	LEFT  = Dir{x: -1, y: 0}
)

func (m Matrix) String() string {
	str := ""
	for _, row := range m {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}

func (m Matrix) RunPartOne() int {
	return 0
}

func (current Dir) isReverse(next Dir) bool {
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

func (m Matrix) visit(maxConsecutive, minConsecutive int) int {
	width := len(m)
	depth := len(m[0])

	startR := Move{x: 0, y: 0, dir: RIGHT, numConsecutive: 0}
	startD := Move{x: 0, y: 0, dir: DOWN, numConsecutive: 0}
	// use Priority Queue, with our heat loss as the priority indicator
	queue := PriorityQueue[Move]{
		&Item[Move]{priority: 0, value: startR, index: 0},
		&Item[Move]{priority: 0, value: startD, index: 1},
	}

	minHeatLoss := map[Move]int{startR: 0, startD: 0}

	heap.Init(&queue)
	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*Item[Move])
		currentMove := current.value
		currentHeatLoss := current.priority
		if minHeatLoss[currentMove] < currentHeatLoss {
			// we've already found a shorter route, so ignore this one
			continue
		}
		if currentMove.x == width-1 && currentMove.y == depth-1 && currentMove.numConsecutive >= minConsecutive {
			// we've reached our destination, exit
			return currentHeatLoss
		}
		for _, nextDirection := range [4]Dir{LEFT, RIGHT, UP, DOWN} {
			nextX, nextY := currentMove.x+nextDirection.x, currentMove.y+nextDirection.y
			if nextX < 0 || nextX >= width || nextY < 0 || nextY >= depth {
				// out of bounds, so ignore
				continue
			}
			consecutive := currentMove.numConsecutive
			if consecutive == maxConsecutive && nextDirection == currentMove.dir {
				// we've hit our max in one direction, we have to ignore
				continue
			}
			if nextDirection.isReverse(currentMove.dir) {
				// we can't go back on ourselves
				continue
			}
			if consecutive < minConsecutive && nextDirection != currentMove.dir {
				consecutive += 1
				// need at least minConsecutive in this direction
				continue
			}
			if nextDirection != currentMove.dir {
				consecutive = 1
			} else {
				consecutive = consecutive%maxConsecutive + 1
			}
			nextMove := Move{x: nextX, y: nextY, numConsecutive: consecutive, dir: nextDirection}
			nextHeatLoss := m[nextY][nextX]
			newCost, found := minHeatLoss[nextMove]
			newHeatLoss := currentHeatLoss + nextHeatLoss
			if found && newCost <= newHeatLoss {
				// if we've already made this next move and it isn't a new maximum, we don't need to do anything
				continue
			}
			// this is new or a new max, so we update the new minimum
			minHeatLoss[nextMove] = newHeatLoss
			// and push this onto our queue
			heap.Push(&queue, &Item[Move]{priority: newHeatLoss, value: nextMove})

		}
	}
	return -1
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
		row := []int{}
		// split string into chars so we can form a 2x2 matrix and inspect the columns
		for _, char := range line {
			num, _ := strconv.ParseInt(string(char), 10, 0)
			row = append(row, int(num))
		}
		matrix = append(matrix, row)
	}

	fmt.Println("Pt1", matrix.visit(3, 0))
	fmt.Println("Pt2", matrix.visit(10, 4))

}
