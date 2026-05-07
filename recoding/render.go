package main

import (
	"fmt"
	"os"
	"strings"
)

func Render(text string, banner map[rune][]string) []string {
	line := make([]string, 8)

	for row := range 8 {
		var lineBuilder strings.Builder
		for _, char := range text {
			lineBuilder.WriteString(banner[char][row])
		}

		line[row] = lineBuilder.String()
	}
	return line
}

func main() {

	input := os.Args[1]
	inpu, _ := LoadBanner("standard.txt")

	fmt.Print(strings.Join(Render(input, inpu), "\n"))
}
