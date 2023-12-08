package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type round struct {
	hand []byte
	bid  int
	rank int
}

type cardCounter map[byte]int

func countCards(cards []byte) cardCounter {
	counts := map[byte]int{}
	for _, c := range cards {
		counts[c]++
	}
	return counts
}
func (cc cardCounter) getMostCommon() (b byte, c int) {
	for k, v := range cc {
		if v > c {
			b = k
			c = v
		}
	}
	return b, c
}

func handType(hand []byte, wildcard bool) int {
	if bytes.Count(hand, []byte{'1'}) == 5 {
		return 7 // five of a kind, joker early return
	}
	if wildcard {
		var handCopy []byte
		for _, b := range hand {
			if b == '1' {
				continue
			}
			handCopy = append(handCopy, b)
		}
		counts := countCards(handCopy)
		b, _ := counts.getMostCommon()
		// overwrite hand
		hand = append(handCopy, bytes.Repeat([]byte{b}, len(hand)-len(handCopy))...)
	}

	counts := countCards(hand)
	numGroups := len(counts)
	b, mostCommon := counts.getMostCommon()
	delete(counts, b)
	_, nextMostCommon := counts.getMostCommon()
	switch numGroups {
	case 1:
		// five of a kind (highest)
		return 7
	case 2:
		if mostCommon == 4 {
			// four of a kind
			return 6
		}
		// full house
		return 5
	case 3:
		if mostCommon == 3 && nextMostCommon == 1 {
			// three of a kind
			return 4
		}
		// two pair
		return 3
	case 4:
		// one pair
		return 2
	default:
		return 1
	}
}

func roundSortFunc(a, b round) int {
	// first sort by hand type
	if a.rank != b.rank {
		return a.rank - b.rank
	}

	// then sort by the card values, left->right
	return bytes.Compare(a.hand, b.hand)
}

func handToRound(hand, bid []byte, wildcard bool) round {
	hand = append(hand[:0:0], hand...)
	// when wildcards enabled, we change how joker gets mapped
	jsub := byte('B')
	if wildcard {
		jsub = '1'
	}
	// convert the cards to hex values according to their rank so we can sort by their bytecode
	for i, c := range hand {
		switch c {
		case 'T':
			hand[i] = 'A'
		case 'J':
			hand[i] = jsub
		case 'Q':
			hand[i] = 'C'
		case 'K':
			hand[i] = 'D'
		case 'A':
			hand[i] = 'E'
		default:
			// leave as-is - numbers will follow this ranking
		}
	}
	return round{hand, mustInt(bid, "bid should be an int"), handType(hand, wildcard)}
}

func parseInput(lines [][]byte, wildcard bool) []round {
	rounds := make([]round, len(lines))
	for i, line := range lines {
		parts := bytes.Split(line, []byte{' '})
		rounds[i] = handToRound(parts[0], parts[1], wildcard)
	}
	return rounds
}

func solve(rounds []round, wildcard bool) int {
	slices.SortFunc(rounds, roundSortFunc)
	answer := 0
	for i, r := range rounds {
		answer += (i + 1) * r.bid
	}
	return answer
}

func part1(data []byte) {
	rounds := parseInput(bytes.Split(data, []byte{'\n'}), false)
	fmt.Printf("Part 1: %d\n", solve(rounds, false))
}

func part2(data []byte) {
	rounds := parseInput(bytes.Split(data, []byte{'\n'}), true)
	fmt.Printf("Part 1: %d\n", solve(rounds, true))
}

func main() {
	start := time.Now()
	data, err := os.ReadFile("./data.txt")
	must(err, "unable to read input file")
	part1(data)
	part2(data)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func mustInt(b []byte, msg string) int {
	i, err := strconv.Atoi(string(b))
	must(err, msg)
	return i
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
