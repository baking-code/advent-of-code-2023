package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 0
	pt1Count := 0
	pt2Count := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		spl := strings.Split(line, " ")
		condString := spl[1]
		springs, condition := spl[0], getNumbersFromString(condString)
		// keep track of the substrings we test and their number of permutations
		// eventually this
		visited := map[string]int{}
		pt1Count += countPermutations(springs, condition, visited)
		fmt.Println(visited)
		// pt2 :'(
		unfoldedSprings := ""
		unfoldedConditions := []int{}
		for i := 0; i < 5; i++ {
			unfoldedSprings += springs
			if i != 4 {
				unfoldedSprings += "?"
			}
			unfoldedConditions = append(unfoldedConditions, condition...)
		}
		visited = map[string]int{}
		// pt2Count += countPermutations(unfoldedSprings, unfoldedConditions, visited)
	}

	fmt.Println("pt1", pt1Count)
	fmt.Println("pt2", pt2Count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func countPermutations(s string, condition []int, visited map[string]int) int {
	// key is spring + condition
	key := s + fmt.Sprintf("%v", condition)
	if n, ok := visited[key]; ok {
		// return early if combination already checked
		return n
	}

	visited[key] = 0
	// if subset is valid, return
	if valid, end := validate(s, condition); end {
		if valid {
			visited[key] = 1
			return 1
		}
		return 0
	}

	// LETS START RECURSIN'

	// if first element is operational, then test the rest against the condition
	if s[0] == '.' {
		n := countPermutations(s[1:], condition, visited)
		visited[key] = n
		return n
	}

	n := 0
	// if first element is unknown, test assuming it is a goodun .
	if s[0] == '?' {
		n += countPermutations(s[1:], condition, visited)
	}

	// test whole spring against single condition
	count, ok := testCondition(s, condition[0])
	if !ok {
		visited[key] = n
		return n
	}
	// s[0:count] must be broken, so we move onto the next condition
	n += countPermutations(s[count:], condition[1:], visited)
	visited[key] = n
	return n
}

func testCondition(s string, condition int) (int, bool) {
	// if spring length is under the condition, then exit early!
	if len(s) < condition {
		return 0, false
	}

	// test each number of the condition
	for i := 0; i < condition; i++ {
		// if the element is already known to work, then return 0
		if s[i] == '.' {
			return 0, false
		}
	}

	// if we have equality, then it's a perfect match :)
	if len(s) == condition {
		return condition, true
	}

	// if last element is either working or unknown, then we return so that can be tested again
	if s[condition] == '.' {
		return condition + 1, true
	}
	if s[condition] == '?' {
		return condition + 1, true
	}
	return 0, false
}

// test a given spring subset by turning each unknown into a known to see if it matches the condition
func validate(s string, cond []int) (bool, bool) {
	numUnknowns := strings.Count(s, "?")
	numKnowns := strings.Count(s, "#")

	// if there's nothing to compare, return false for this comparison
	if numKnowns > 0 && len(cond) == 0 {
		return false, true
	}
	// replace all unknowns so we can start testing
	replaced := strings.ReplaceAll(s, "?", ".")
	end := numUnknowns == 0

	// split the springs by operational groups
	groups := strings.FieldsFunc(replaced, func(r rune) bool {
		return r == '.'
	})
	fmt.Println(replaced, groups)

	// if there's a length mismatch, then return early
	if len(groups) != len(cond) {
		return false, end
	}
	// for each condition, test against our groups
	for i, c := range cond {
		if c != len(groups[i]) {
			return false, end
		}
	}
	// if all groups match the conds, then it's a valid combo!
	return true, true
}

func getNumbersFromString(s string) []int {
	strings := strings.Split(s, ",")
	numbers := make([]int, 0)
	for _, str := range strings {
		if str != " " && str != "" {
			num, err := strconv.ParseInt(str, 10, 0)
			if err != nil {
				fmt.Println(err)
			} else {
				numbers = append(numbers, int(num))
			}
		}
	}
	return numbers
}
