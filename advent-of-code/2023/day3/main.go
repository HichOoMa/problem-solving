package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day3/puzzle.txt",
	)
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	count := 0
	for lineIndex, line := range lines {
		charIndex := 0
	charlopp:
		// for charIndex, str := range line {
		for charIndex < len(line) {
			str := line[charIndex]
			if _, b := isInt(string(str)); b {
				number, lastIndex := getFullNumber(line, charIndex)
				for i := lineIndex - 1; i < lineIndex+2; i++ {
					for j := charIndex - 1; j < lastIndex+1; j++ {
						if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[i]) {
							continue
						}
						if _, b := isInt(string(lines[i][j])); !b && string(lines[i][j]) != "." {
							count = count + number
							charIndex = lastIndex
							continue charlopp
						}
					}
				}
				charIndex = lastIndex
			}
			charIndex++
		}
	}
	fmt.Println(count)
}

type gear struct {
	vi      int
	hi      int
	numbers []int
}

var gears []gear

func part2(lines []string) {
	for lineIndex, line := range lines {
		charIndex := 0
	charlopp:
		for charIndex < len(line) {
			str := line[charIndex]
			if _, b := isInt(string(str)); b {
				number, lastIndex := getFullNumber(line, charIndex)
				for i := lineIndex - 1; i < lineIndex+2; i++ {
					for j := charIndex - 1; j < lastIndex+1; j++ {
						if i < 0 || j < 0 || i >= len(lines) || j >= len(lines[i]) {
							continue
						}
						if _, b := isInt(string(lines[i][j])); !b && string(lines[i][j]) == "*" {
							insertNum(i, j, number)
							charIndex = lastIndex
							continue charlopp
						}
					}
				}
				charIndex = lastIndex
			}
			charIndex++
		}
	}
	count := 0
	for _, g := range gears {
		if len(g.numbers) < 2 {
			continue
		}
		mul := 1
		for _, num := range g.numbers {
			mul = mul * num
		}
		count = count + mul
	}
	fmt.Println(count)
}

func isInt(str string) (int, bool) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return -1, false
	}
	return n, true
}

func insertNum(line int, index int, number int) {
	for i, g := range gears {
		if g.vi == line && g.hi == index {
			gears[i].numbers = append(gears[i].numbers, number)
			return
		}
	}
	gears = append(gears, gear{line, index, []int{number}})
}

// return number and the index of last degits of the number
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
