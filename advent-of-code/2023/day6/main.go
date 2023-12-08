package main

import (
	"fmt"
	"strconv"
	"strings"

	pkg "github.com/HaithamBH/problem-solving/advent-of-code"
)

func main() {
	lines := pkg.FileIntoLines(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day6/puzzle.txt",
	)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	timeSlc := strings.Split(lines[0], ":")
	distanceSlc := strings.Split(lines[1], ":")

	index := 0
	timerList := []int{}
	for index < len(timeSlc[1]) {
		if string(timeSlc[1][index]) == " " {
			index++
			continue
		} else {
			num := 0
			for {
				if index == len(timeSlc[1]) || string(timeSlc[1][index]) == " " {
					timerList = append(timerList, num)
					num = 0
					break
				} else {
					n, _ := strconv.Atoi(string(timeSlc[1][index]))
					num = num*10 + n
					index++
					continue
				}
			}
		}
	}

	index = 0
	distanceList := []int{}
	for index < len(distanceSlc[1]) {
		if string(distanceSlc[1][index]) == " " {
			index++
			continue
		} else {
			num := 0
			for {
				if index == len(distanceSlc[1]) || string(distanceSlc[1][index]) == " " {
					distanceList = append(distanceList, num)
					num = 0
					break
				} else {
					n, _ := strconv.Atoi(string(distanceSlc[1][index]))
					num = num*10 + n
					index++
					continue
				}
			}
		}
	}

	count := 1
	for i := 0; i < len(timerList); i++ {
		recordCount := 0
		for timer := 1; timer < timerList[i]; timer++ {
			if d := timer * (timerList[i] - timer); d > distanceList[i] {
				recordCount++
			}
		}
		count = count * recordCount
	}
	fmt.Println(count)
}

func part2(lines []string) {
	timeSlc := strings.Split(lines[0], ":")
	distanceSlc := strings.Split(lines[1], ":")

	index := 0
	time := ""
	for index < len(timeSlc[1]) {
		if string(timeSlc[1][index]) == " " {
			index++
			continue
		} else {
			for {
				if index == len(timeSlc[1]) || string(timeSlc[1][index]) == " " {
					break
				} else {
					time = time + string(timeSlc[1][index])
					index++
					continue
				}
			}
		}
	}

	index = 0
	distance := ""
	for index < len(distanceSlc[1]) {
		if string(distanceSlc[1][index]) == " " {
			index++
			continue
		} else {
			for {
				if index == len(distanceSlc[1]) || string(distanceSlc[1][index]) == " " {
					break
				} else {
					distance = distance + string(distanceSlc[1][index])
					index++
					continue
				}
			}
		}
	}

	timeNum, _ := strconv.Atoi(time)
	distanceNum, _ := strconv.Atoi(distance)
	index = 1
	for {
		if d := index * (timeNum - index); d > distanceNum {
			break
		}
		index++
	}
	fmt.Println(timeNum - 2*index + 1)
}
