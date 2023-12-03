package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day2/puzzle.txt",
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
	for _, line := range lines {
		slices := strings.Split(line, ":")
		if slices[0] == "" {
			continue
		}
		sets := strings.Split(slices[1], ";")
		isPossible := true
		for _, set := range sets {
			blue, green, red := 0, 0, 0
			colors := strings.Split(set, ",")
			for _, color := range colors {
				items := strings.Split(color, " ")
				n, _ := strconv.Atoi(items[1])

				switch items[2] {
				case "blue":
					blue = int(n)
				case "green":
					green = int(n)
				case "red":
					red = int(n)
				default:
				}
				// fmt.Println(blue, green, red)
			}
			if blue > 14 || green > 13 || red > 12 {
				isPossible = false
				break
			}
		}
		if isPossible {
			id, _ := strconv.Atoi(slices[0][5:])
			count = count + int(id)
		}
	}
	fmt.Println(count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		slices := strings.Split(line, ":")
		if slices[0] == "" {
			continue
		}
		sets := strings.Split(slices[1], ";")
		blue, green, red := 0, 0, 0
		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, color := range colors {
				items := strings.Split(color, " ")
				n, _ := strconv.Atoi(items[1])

				switch items[2] {
				case "blue":
					if blue < int(n) {
						blue = int(n)
					}
				case "green":
					if green < int(n) {
						green = int(n)
					}
				case "red":
					if red < int(n) {
						red = int(n)
					}
				default:
				}
			}
		}
		count = count + blue*green*red
	}
	fmt.Println(count)
}
