package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// digits := regexp.MustCompile("(\\d+)")

	// totalScore := 0
	type mapper struct {
		source int
		dest   int
		number int
	}
	mapNames := map[int]string{}
	maps := map[int][]mapper{}
	lineNo := 0
	seeds := []int{}
	currentMapIndex := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if lineNo == 1 {
			// gather the seeds
			numbers := strings.Split(line, ":")[1]
			seeds = getNumbersFromString(numbers)
		} else if strings.Contains(line, "map") {
			mapName := strings.Split(strings.Split(line, "-")[2], " ")[0]
			currentMapIndex++
			maps[currentMapIndex] = make([]mapper, 0)
			mapNames[currentMapIndex] = mapName
		} else if r := getNumbersFromString(line); len(r) > 0 {
			maps[currentMapIndex] = append(maps[currentMapIndex], mapper{dest: r[0], source: r[1], number: r[2]})
		}

	}
	// ranges := map[int]map[int]int{}
	// for ind, m := range maps {
	// 	rangeMap := map[int]int{}
	// 	for _, mapper := range m {
	// 		for i := 0; i < int(mapper.number); i++ {
	// 			rangeMap[mapper.source+i] = mapper.dest + i
	// 		}
	// 	}
	// 	ranges[ind] = rangeMap
	// }

	// seed 79 soil 81 fertilizer 81 water 81  light 74 temperature 78 humidity 78 location 82
	// seed 14 soil 14 fertilizer 53 water 49 light 42 temperature 42 humidity 43 location 43
	// seed 55 soil 57 fertilizer 57 water 53 light 46 temperature 82 humidity 82 location 86
	// seed 13 soil 13 fertilizer 52 water 41 light 34 temperature 34 humidity 35 location 35

	finalSeeds := make([]int, 0)
	for _, seed := range seeds {
		track := seed
		fmt.Print("seed ", seed, " ")
		for i := 1; i < len(maps)+1; i++ {
			var mapper = maps[i]
			for _, v := range mapper {
				// fmt.Println(mapNames[i])
				if track >= v.source && track < v.source+v.number {
					track = (track - v.source) + v.dest
					break
				}
			}
			fmt.Print(mapNames[i], " ", track, " ")
		}
		fmt.Println("")
		finalSeeds = append(finalSeeds, track)
	}

	// fmt.Println(mapNames)
	// fmt.Println(ranges)
	// fmt.Println("seeds", seeds, finalSeeds)
	lowest := math.MaxInt
	for _, s := range finalSeeds {
		if s < lowest {
			lowest = s
		}
	}
	fmt.Println("lowest", lowest)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumbersFromString(s string) []int {
	strings := strings.Split(s, " ")
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
