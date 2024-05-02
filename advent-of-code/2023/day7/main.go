package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	pkg "github.com/HaithamBH/problem-solving/advent-of-code"
)

func main() {
	lines := pkg.FileIntoLines(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day7/puzzle.txt",
	)

	part1(lines)
}

type player struct {
	hand string
	typ  int
	pid  int
}

func part1(lines []string) {
	players := []player{}
	for _, line := range lines {
		slc := strings.Split(line, " ")
		hand, pidStr, slc := slc[0], slc[1], nil
		pid, _ := strconv.Atoi(pidStr)
		occLabels := make(map[string]int)
		for _, char := range hand {
			if occ, exist := occLabels[string(char)]; exist {
				occLabels[string(char)] = occ + 1
			} else {
				occLabels[string(char)] = 1
			}
		}
		typ := 0
		score := 0
		for key := range occLabels {
			score = score + int(math.Pow(float64(occLabels[key]), 2))
		}
		switch score {
		case 5:
			typ = 1
		case 7:
			typ = 2
		case 9:
			typ = 3
		case 11:
			typ = 4
		case 13:
			typ = 5
		case 17:
			typ = 6
		case 25:
			typ = 5
		}
		players = append(players, player{hand, typ, pid})
	}
	labels := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	count := 0
	for i := 0; i < len(players)-1; i++ {
		maxi := i
		for j := i + 1; j < len(players); j++ {
			if players[maxi].typ < players[j].typ {
				maxi = j
			} else if players[maxi].typ == players[j].typ {
				for charIndex, char := range players[maxi].hand {
					if IndexOf(labels, char) > IndexOf(labels, rune(players[j].hand[charIndex])) {
						maxi = j
						break
					} else if IndexOf(labels, char) < IndexOf(labels, rune(players[j].hand[charIndex])) {
						break
					}
				}
			}
		}
		players[i], players[maxi] = players[maxi], players[i]
		count = count + (len(players)-i)*players[i].pid
	}
	count = count + players[len(players)-1].pid
	fmt.Println(count)
}

func IndexOf(T []rune, x rune) int {
	for i := 0; i < len(T); i++ {
		if T[i] == x {
			return i
		}
	}
	return -1
}
