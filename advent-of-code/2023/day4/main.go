package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	pkg "github.com/HaithamBH/problem-solving/advent-of-code"
)

func main() {
	lines := pkg.FileIntoLines("/home/hichoma/Dev/problem-solving/advent-of-code/2023/day4/puzzle.txt")

	// part1(lines)
	part2(lines)
}

func part1(lines []string) {
	count := 0
	for _, line := range lines {
		data := strings.Split(line, ":")
		isPlayNum := false
		index := 0
		playNums := []int{}
		winNums := []int{}
		for index < len(data[1]) {
			char := data[1][index]
			if string(char) == "|" {
				isPlayNum = true
				index++
				continue
			} else if string(char) == " " {
				index++
				continue
			} else {
				var num int
				if index < len(data[1])-1 && string(data[1][index+1]) == " " || index == len(data[1])-1 {
					num, _ = strconv.Atoi(string(data[1][index]))
				} else {
					num, _ = strconv.Atoi(data[1][index : index+2])
					index++
				}
				if isPlayNum {
					playNums = append(playNums, num)
				} else {
					winNums = append(winNums, num)
				}
			}
			index++
		}
		points := 0
		for _, pNum := range playNums {
			if slices.Contains(winNums, pNum) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
		count = count + points
	}
	fmt.Print(count)
}

func part2(lines []string) {
	scratchcards := make([]int, len(lines))

	for gameIndex, line := range lines {
		scratchcards[gameIndex] = scratchcards[gameIndex] + 1
		data := strings.Split(line, ":")
		isPlayNum := false
		index := 0
		playNums := []int{}
		winNums := []int{}
		for index < len(data[1]) {
			char := data[1][index]
			if string(char) == "|" {
				isPlayNum = true
				index++
				continue
			} else if string(char) == " " {
				index++
				continue
			} else {
				var num int
				if index < len(data[1])-1 && string(data[1][index+1]) == " " || index == len(data[1])-1 {
					num, _ = strconv.Atoi(string(data[1][index]))
				} else {
					num, _ = strconv.Atoi(data[1][index : index+2])
					index++
				}
				if isPlayNum {
					playNums = append(playNums, num)
				} else {
					winNums = append(winNums, num)
				}
			}
			index++
		}
		matches := 0
		for _, pNum := range playNums {
			if slices.Contains(winNums, pNum) {
				matches++
			}
		}
		for i := 1; i <= matches; i++ {
			scratchcards[gameIndex+i] = scratchcards[gameIndex+i] + scratchcards[gameIndex]
		}
	}
	score := 0
	for _, i := range scratchcards {
		score = score + i
	}
	fmt.Print(score)
}
