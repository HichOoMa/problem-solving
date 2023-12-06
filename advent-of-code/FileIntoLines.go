package pkg

import (
	"fmt"
	"os"
	"strings"
)

func FileIntoLines(uri string) []string {
	fileBytes, err := os.ReadFile(
		"/home/hichoma/Dev/problem-solving/advent-of-code/2023/day3/puzzle.txt",
	)
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	return lines
}
