package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func colour(text, color string) string {
	return color + text + "\033[0m"
}

func main() {
	line := "Hello World"

	AsciiArt(line, "\033[1;33m")
	time.Sleep(1 * time.Second)
}

func AsciiArt(input, color string) {
	inputFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading font file:", err)
		return
	}

	lines := strings.Split(string(inputFile), "\n")

	for row := 0; row < 8; row++ {
		lineOutput := ""
		for _, char := range input {
			if char < ' ' || char > '~' {
				continue
			}

			index := row + (int(char-' ') * 9) + 1
			if index < len(lines) {
				lineOutput += lines[index]
			}
		}
		fmt.Println(colour(lineOutput, color))
		time.Sleep(1 * time.Second)

	}
}
