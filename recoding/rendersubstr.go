package main

import (
	"strings"
)

func RenderSubstring(substr, color string, banner map[rune][]string) string {
	var word []string
	for j := 0; j < 8; j++ {
		words := ""
		for _, i := range substr {
			value, exist := banner[i]
			if exist {
				words += value[j]
			}
		}
		if color == "red" {
			word = append(word, "\033[33m"+words+"\033[0m")
		}
	}
	return strings.Join(word, "\n")
}
