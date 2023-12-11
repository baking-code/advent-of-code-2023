package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	copies := 1000000
	rowsWithoutGalaxies := []int{}

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
			// matrix = append(matrix, row)
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, x)
			rowsInserted++
		}
		x++
	}
	colsWithoutGalaxies := []int{}
	// for y := 0; y < len(matrix)-rowsInserted; y++ {
	// 	hasGalaxies := false
	// 	for i := 0; i < len(matrix); i++ {
	// 		s := matrix[i][y]
	// 		if string(s) == "#" {
	// 			hasGalaxies = hasGalaxies || true
	// 		}
	// 	}
	// 	if !hasGalaxies {
	// 		colsWithoutGalaxies = append(colsWithoutGalaxies, y)
	// 	}
	// }

	//marker := 0
	// for _, v := range colsWithoutGalaxies {
	// 	for i, row := range matrix {
	// 		// fmt.Println("inserting item at col number", v+marker, "for row number", i)
	// 		matrix[i] = slices.Insert(row, v+marker, ".")
	// 	}
	// 	fmt.Println(matrix)
	// 	marker++
	// }
	fmt.Println(matrix)

	for y := 0; y < len(matrix); y++ {
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
	// fmt.Println(matrix, colsWithoutGalaxies)
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
	// fmt.Println(matrix)
	// fmt.Println(galaxies)

	total := 0
	pt2Total := 0
	for i, v := range galaxies {
		if i < len(galaxies) {
			for j := i; j < len(galaxies)-1; j++ {
				total += getShortestPathPt2(v, galaxies[j+1], rowsWithoutGalaxies, colsWithoutGalaxies, 1)
				pt2Total += getShortestPathPt2(v, galaxies[j+1], rowsWithoutGalaxies, colsWithoutGalaxies, copies)
			}
		} else {
			total += getShortestPathPt2(v, galaxies[0], rowsWithoutGalaxies, colsWithoutGalaxies, 1)
			pt2Total += getShortestPathPt2(v, galaxies[0], rowsWithoutGalaxies, colsWithoutGalaxies, copies)

		}
	}

	fmt.Println("total Pt1", total)    //374, 9799681
	fmt.Println("total Pt1", pt2Total) //

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
func getShortestPathPt2(a, b Coord, rowsWithoutGalaxies []int, colsWithoutGalaxies []int, multiple int) int {
	dx := b.x - a.x
	dy := b.y - a.y
	xx, yy := multiplier(a, b, rowsWithoutGalaxies, colsWithoutGalaxies)
	absx := int(math.Abs(float64(dx)))
	absy := int(math.Abs(float64(dy)))
	// we define the number of copies we want, we already have the original so we -1 from the desired number
	res := absx + (xx*multiple - 1) + absy + yy*multiple - 1
	// fmt.Println(a, b, absx, xx*multiple, absy, yy*multiple, res)
	return res
}

func multiplier(a, b Coord, listX, listY []int) (int, int) {
	resultX := 0
	resultY := 0
	for _, d := range listX {
		minX := int(math.Min(float64(a.x), float64(b.x)))
		maxX := int(math.Max(float64(a.x), float64(b.x)))
		if minX < d && d < int(maxX) {
			resultX++
		}
	}
	for _, d := range listY {
		minY := int(math.Min(float64(a.y), float64(b.y)))
		maxY := int(math.Max(float64(a.y), float64(b.y)))
		if minY < d && d < maxY {
			resultY++
		}
	}
	// fmt.Println("gaps between", a, b, "are: ", resultX, resultY)
	return resultX, resultY
}
