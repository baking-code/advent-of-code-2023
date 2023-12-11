package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := 0
	rowsInserted := 0
	multiple := 1

	// rowsWithoutGalaxies := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		hasGalaxies := false
		for _, s := range line {
			row = append(row, string(s))
			if string(s) == "#" {
				hasGalaxies = hasGalaxies || true
			}
		}
		matrix = append(matrix, row)
		if !hasGalaxies {
			for i := 0; i < multiple; i++ {
				matrix = append(matrix, row)
				rowsInserted++
			}
			// rowsWithoutGalaxies = append(rowsWithoutGalaxies, )
		}
		x++
	}
	colsWithoutGalaxies := []int{}
	for y := 0; y < len(matrix)-rowsInserted; y++ {
		hasGalaxies := false
		for i := 0; i < len(matrix); i++ {
			s := matrix[i][y]
			if string(s) == "#" {
				hasGalaxies = hasGalaxies || true
			}
		}
		if !hasGalaxies {
			colsWithoutGalaxies = append(colsWithoutGalaxies, y)
		}
	}

	fmt.Println(matrix, colsWithoutGalaxies)
	marker := 0
	for _, v := range colsWithoutGalaxies {
		for i, row := range matrix {
			// fmt.Println("inserting item at col number", v+marker, "for row number", i)
			for k := 0; k < multiple; i++ {
				matrix[i] = slices.Insert(row, v+marker, ".")
			}
		}
		fmt.Println(matrix)
		marker += multiple
	}
	fmt.Println(matrix)

	for y := 0; y < len(matrix)-rowsInserted; y++ {
		hasGalaxies := false
		for i := 0; i < len(matrix); i++ {
			s := matrix[i][y]
			if string(s) == "#" {
				hasGalaxies = hasGalaxies || true
			}
		}
		if !hasGalaxies {
			colsWithoutGalaxies = append(colsWithoutGalaxies, y)
		}
	}
	fmt.Println(matrix)
	numRows := len(matrix)
	numCols := len(matrix[0])
	fmt.Println("rows:", numRows, "; cols:", numCols)

	galaxies := []Coord{}
	for x := 0; x < numRows; x++ {
		for y := 0; y < numCols; y++ {
			s := matrix[x][y]
			if string(s) == "#" {
				galaxies = append(galaxies, Coord{x, y})
				matrix[x][y] = fmt.Sprint(len(galaxies))
			}
		}
	}
	fmt.Println(galaxies)

	total := 0
	for i, v := range galaxies {
		if i < len(galaxies) {
			for j := i; j < len(galaxies)-1; j++ {
				total += getShortestPath(v, galaxies[j+1])
			}
		} else {
			total += getShortestPath(v, galaxies[0])
		}
	}

	fmt.Println("total", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getShortestPath(a, b Coord) int {
	dx := b.x - a.x
	dy := b.y - a.y
	res := int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
	// fmt.Println(a, b, res)
	return res
}
