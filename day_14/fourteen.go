package main

import (
	"bufio"
	"fmt"
	"os"
)

type Matrix [][]string

func (m Matrix) String() string {
	str := ""
	for _, row := range m {
		str += fmt.Sprintf("%v\n", row)
	}
	return str
}
func (m Matrix) isEqual(n Matrix) bool {
	depth := len(m)
	width := len(m[0])
	for x := 0; x < width; x++ {
		for y := 0; y < depth; y++ {
			if m[y][x] != n[y][x] {
				return false
			}
		}
	}
	return true
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
	depth := len(matrix)
	width := len(matrix[0])
	newMatrix := rollVert(depth, width, matrix, true)

	pt1Total := getTotalLoad(newMatrix, depth)
	fmt.Println("Part 1:", pt1Total)

	// Part 2
	newMatrix = matrix
	matrices := []Matrix{}
	cycleStarts := 0
	cycleEnds := 0
	bill := 1000000000
	for i := 1; i <= bill; i++ {
		newMatrix = rollVert(depth, width, newMatrix, true)   // north
		newMatrix = rollHoriz(depth, width, newMatrix, true)  // west
		newMatrix = rollVert(depth, width, newMatrix, false)  // south
		newMatrix = rollHoriz(depth, width, newMatrix, false) // east
		for j, v := range matrices {
			if v.isEqual(newMatrix) {
				fmt.Println("HAHA", i, j)
				cycleStarts = i
				cycleEnds = j + 1
			}
		}
		if cycleStarts > 0 {
			break
		} else {
			matrices = append(matrices, newMatrix)
		}
	}
	period := cycleStarts - cycleEnds
	which := cycleStarts + (((bill - cycleStarts) % period) - period) - 1
	fmt.Println("starts on", cycleEnds, "repeats on", cycleStarts, "cycle", period)
	pt2Total := getTotalLoad(matrices[which], depth)
	fmt.Println("Part 2:", pt2Total, which)

}

func getTotalLoad(newMatrix Matrix, depth int) int {
	pt1Total := 0
	for i, row := range newMatrix {
		multiplier := depth - i
		boulderCount := 0
		for _, char := range row {
			if char == "O" {
				boulderCount++
			}
		}
		pt1Total += (boulderCount * multiplier)
	}
	return pt1Total
}

func rollVert(depth int, width int, matrix Matrix, isNorth bool) Matrix {
	newMatrix := make(Matrix, depth)
	for x := 0; x < width; x++ {
		newMatrix[x] = make([]string, width)
	}
	for x := 0; x < width; x++ {
		availableSpaces := []int{}
		boulders := 0
		var doRoll = func(y int) {
			char := matrix[y][x]
			if char == "#" {
				fillAvailableVert(availableSpaces, boulders, newMatrix, x)
				newMatrix[y][x] = char
				availableSpaces = []int{}
				boulders = 0
			} else if char == "." {
				availableSpaces = append(availableSpaces, y)
			} else if char == "O" {
				availableSpaces = append(availableSpaces, y)
				boulders++
			}
		}
		if isNorth {
			for y := 0; y < depth; y++ {
				doRoll(y)
			}
		} else {
			for y := depth - 1; y >= 0; y-- {
				doRoll(y)
			}
		}
		fillAvailableVert(availableSpaces, boulders, newMatrix, x)
	}
	return newMatrix
}

func rollHoriz(depth int, width int, matrix Matrix, isWest bool) Matrix {
	newMatrix := make(Matrix, depth)
	for x := 0; x < width; x++ {
		newMatrix[x] = make([]string, width)
		availableSpaces := []int{}
		boulders := 0
		var doRoll = func(y int) {
			char := matrix[x][y]
			if char == "#" {
				fillAvailableHoriz(availableSpaces, boulders, newMatrix, x)
				newMatrix[x][y] = char
				availableSpaces = []int{}
				boulders = 0
			} else if char == "." {
				availableSpaces = append(availableSpaces, y)
			} else if char == "O" {
				availableSpaces = append(availableSpaces, y)
				boulders++
			}
		}
		if isWest {
			for y := 0; y < width; y++ {
				doRoll(y)
			}
		} else {
			for y := width - 1; y >= 0; y-- {
				doRoll(y)
			}
		}
		fillAvailableHoriz(availableSpaces, boulders, newMatrix, x)
	}
	return newMatrix
}

func fillAvailableVert(availableSpaces []int, boulders int, newMatrix Matrix, x int) {
	for _, space := range availableSpaces {
		if boulders > 0 {
			newMatrix[space][x] = "O"
			boulders--
		} else {
			newMatrix[space][x] = "."
		}
	}
}

func fillAvailableHoriz(availableSpaces []int, boulders int, newMatrix Matrix, y int) {
	for _, space := range availableSpaces {
		if boulders > 0 {
			newMatrix[y][space] = "O"
			boulders--
		} else {
			newMatrix[y][space] = "."
		}
	}
}
