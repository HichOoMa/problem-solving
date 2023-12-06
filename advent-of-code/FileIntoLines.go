package pkg

import (
	"fmt"
	"os"
	"strings"
)

func FileIntoLines(uri string) []string {
	fileBytes, err := os.ReadFile(uri)
	if err != nil {
		fmt.Print(err)
	}
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	return lines
}
