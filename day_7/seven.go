package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type collect struct {
	hand  string
	score int
	rank  int
}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0

	hands := []collect{}

	for scanner.Scan() {
		lineNo++
		line := strings.Split(scanner.Text(), " ")
		hand, score := line[0], getNumberFromString(line[1])
		rank := computeHand(hand)
		hands = append(hands, collect{hand: hand, score: score, rank: rank})
	}

	total := 0
	// fmt.Println(hands)
	sortHands(hands)
	for _, h := range hands {
		// fmt.Println(h.score, lineNo)
		total += h.score * lineNo
		lineNo--
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumberFromString(s string) int {
	strings := strings.Split(s, " ")
	numberString := ""
	for _, str := range strings {
		if str != " " && str != "" {
			numberString += str
		}
	}
	num, err := strconv.ParseInt(numberString, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	return int(num)
}

func computeHand(s string) int {
	countMap := map[string]int{}
	for _, v := range s {
		countMap[string(v)]++
	}
	max := 0
	for _, v := range countMap {
		if v > max {
			max = v
		}
	}
	length := len(countMap)
	return max - length
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
var handMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func sortHand(hand string) string {
	bytes := strings.Split(hand, "")
	sort.SliceStable(bytes, func(i, j int) bool {
		return handMap[bytes[j]] < handMap[bytes[i]]
	})
	return strings.Join(bytes, "")
}

func sortHands(hands []collect) {
	sort.Slice(hands, func(i, j int) bool {
		// sort by rank
		rankA, rankB := hands[j].rank, hands[i].rank
		if rankA != rankB {
			return rankA < rankB
		}

		// then sort by each card in order
		a := []byte(hands[j].hand)
		b := []byte(hands[i].hand)
		// for k := 0; k < 5; k++ {
		// 	charA := string(a[k])
		// 	charB := string(b[k])
		// 	return handMap[charA] < handMap[charB]
		// }
		// return false
		return bytes.Compare(a, b) >= 0
	})
}
