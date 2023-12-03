package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day1/puzzle.txt",
	)
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	part1(lines)
}

func part1(lines []string) {
	count := 0
	for _, str := range lines {
		numberList := []int{}
		for _, char := range str {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			numberList = append(numberList, n)
		}
		number := 0
		if len(numberList) == 1 {
			number = numberList[0] * 11
		} else if len(numberList) > 1 {
			number = numberList[0]*10 + numberList[len(numberList)-1]
		}
		count = count + number
	}
	fmt.Print(count)
}

func part2(lines []string) {
	numberStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	count := 0
	for _, str := range lines {
		numberList := []int{}
		for i, char := range str {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				for j := 3; j < 6; j++ {
					if i+j <= len(str) {
						index := slices.Index(numberStr, str[i:i+j])
						if index != -1 {
							numberList = append(numberList, index+1)
							break
						}
					}
				}
				continue
			}
			numberList = append(numberList, n)
		}
		fmt.Print(numberList)
		number := 0
		if len(numberList) == 1 {
			number = numberList[0] * 11
		} else if len(numberList) > 1 {
			number = numberList[0]*10 + numberList[len(numberList)-1]
		}
		count = count + number
	}
	fmt.Print(count)
}
