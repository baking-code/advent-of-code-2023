package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Matrix [][]string

func (m Matrix) String() string {
	str := ""
	for i, row := range m {
		str += fmt.Sprintf("%d%v\n", i+1, row)
	}
	return str
}

func main() {
	rowReflectionCount := 0
	columnReflectionCount := 0

	var processRows = func(picture []string, ind int) {
		matrix := Matrix{}
		// check for mirrored rows
		for _, line := range picture {
			row := []string{}
			// split string into chars so we can form a 2x2 matrix and inspect the columns
			for _, char := range line {
				row = append(row, string(char))
			}
			matrix = append(matrix, row)
		}
		rCount, ok := countMirrors(matrix, "rows", ind)
		rowReflectionCount += rCount
		if !ok {
			transposed := transpose(matrix)
			cCount, _ := countMirrors(transposed, "cols", ind)
			columnReflectionCount += cCount
		}

	}

	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 0
	picNo := 0
	rows := []string{}
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if line != "" {
			rows = append(rows, line)
		} else {
			picNo++
			processRows(rows, picNo)
			rows = []string{}
		}
	}
	picNo++
	processRows(rows, picNo)
	fmt.Println("pt1", rowReflectionCount, columnReflectionCount, columnReflectionCount+(100*rowReflectionCount))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func transpose(a [][]string) [][]string {
	depth := len(a)
	width := len(a[0])
	newArr := make([][]string, width)
	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

var areRowsEqual = func(rowA, rowB []string) bool {
	every := true
	for i := 0; i < len(rowA); i++ {
		if rowA[i] != rowB[i] {
			every = false
			break
		}
	}
	return every
}
var areRowsMirrored = func(rows [][]string, ind int) bool {
	low := ind - 1
	high := ind
	for low >= 0 && high < len(rows) {
		if areRowsEqual(rows[low], rows[high]) {
			low--
			high++
		} else {
			break
		}
	}
	if (low == 0 && high == len(rows)-1) || (low == 1 && high == len(rows)) || (low == 0 && high == len(rows)) {
		return true
	}
	return false
}

var countMirrors = func(matrix Matrix, desc string, index int) (int, bool) {
	count := 0
	justUnderHalf := (len(matrix))/2 - 1
	justOverHalf := (len(matrix)+1)/2 + 1
	fmt.Println("matrix:", index, desc)
	fmt.Println(matrix)
	for i := justUnderHalf; i < justOverHalf; i++ {
		if areRowsMirrored(matrix, i) {
			fmt.Println("mirror found at", i, desc, index)
			count += i
			break
		}
	}
	return count, count != 0
}
