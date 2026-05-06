package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) < 856 {
		return nil, errors.New("invalid format; insufficent lines") 
	}

	result := make(map[rune][]string)

	for ascii := ' '; ascii <= '~'; ascii++ {
		start := int(ascii-32)*9

		result[ascii] = lines[start+1 : start+9]
	}
	return result, nil
}

func main() {


	fmt.Println(LoadBanner("standard.txt"))
}
