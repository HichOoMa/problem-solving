package main

import (
	"fmt"
	"strconv"
	"strings"

	pkg "github.com/HaithamBH/problem-solving/advent-of-code"
)

func main() {
	lines := pkg.FileIntoLines(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day5/puzzle.txt",
	)

	part1(lines)
}

type rangeMap struct {
	startRange int
	startIndex int
	lastIndex  int
}

var rgMap []rangeMap

func part1(lines []string) {
	seeds := strings.Split(lines[0], ":")

	trackList := []int{}
	charIndex := 1
	for charIndex < len(seeds[1]) {
		if string(seeds[1][charIndex]) != " " {
			seed, lastIndex := getFullNumber(seeds[1], charIndex)
			trackList = append(trackList, seed)
			charIndex = lastIndex
			continue
		}
		charIndex++
	}
	mapNum := 1
	for _, line := range lines[2:] {
		if line == "" {
			for trackedIndex, tracked := range trackList {
				for _, rg := range rgMap {
					if tracked >= rg.startIndex && tracked <= rg.lastIndex {
						trackList[trackedIndex] = rg.startRange + tracked - rg.startIndex
					}
				}
			}
			mapNum++
			rgMap = []rangeMap{}
			continue
		}
		if _, b := isInt(string(line[0])); !b {
			continue
		}
		index := 0
		var numList []int
		for index < len(line) {
			char := line[index]
			if _, b := isInt(string(char)); b {
				num, newIndex := getFullNumber(line, index)
				index = newIndex
				numList = append(numList, num)
				continue
			}
			index++
		}
		rgMap = append(rgMap, rangeMap{numList[0], numList[1], numList[1] + numList[2] - 1})
	}
	for trackedIndex, tracked := range trackList {
		for _, rg := range rgMap {
			if tracked >= rg.startIndex && tracked <= rg.lastIndex {
				trackList[trackedIndex] = rg.startRange + tracked - rg.startIndex
			}
		}
	}
	min := 9223372036854775807
	for _, tracked := range trackList {
		if tracked < min {
			min = tracked
		}
	}
	fmt.Print(min)
}

func isInt(str string) (int, bool) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return -1, false
	}
	return n, true
}

func getFullNumber(str string, index int) (int, int) {
	n := 0
	for i := index; i < len(str); i++ {
		if int, b := isInt(string(str[i])); b {
			n = n*10 + int
		} else {
			return n, i
		}
	}
	return n, len(str)
}
