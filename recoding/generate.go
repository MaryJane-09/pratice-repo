package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	lines := Split(input)
	newline := true

	for _, seg := range lines {
		if seg != "" {
			newline = false
			break
		}
	}
	
		for newline {
			return strings.Repeat("\n", len(lines)-1)
		}
	
	
	var sb strings.Builder

	for i, seg := range lines {
		if seg == "" {
			if i > len(lines)-1 {
				sb.WriteString("\n")
			}
			continue
		}

		row := Render(seg, banner)

		for _, r := range row {
			//w := r + "\n"
			sb.WriteString(r + "\n")
		}
	}

	return sb.String()
}
