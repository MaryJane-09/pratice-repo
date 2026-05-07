package main

import "strings"

func Split(input string) [] string {
	splitlines := strings.Split(input, "\\n")

	return splitlines
}