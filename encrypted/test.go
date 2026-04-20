package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(input string) []string {
	var tokens []string
	current := ""

	for _, ch := range input {

		if unicode.IsLetter(ch) {
			current += string(ch)

		} else if unicode.IsDigit(ch) {
			current += string(ch)

		} else {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, string(ch))

		}
	}
	if current != "" {
		tokens = append(tokens, current)
	}

	return tokens
}

func HexaDeca(words []string) string {
	var result []string

	for i := 0; i < len(words); i++ {
		val, err := strconv.ParseInt(words[i], 36, 64)
		if err == nil {
			result = append(result, strconv.FormatInt(val, 16))
			continue
		}
		result = append(result, words[i])

	}

	return strings.Join(result, "")
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Greetings Agent!!!\n got any encrypted puzzles for us today?\n 1.SURE\n or\n 2.NAHH")
	for {
		fmt.Print("> ")
		blocks, _ := reader.ReadString('\n')
		blocks = strings.TrimSpace(blocks)
		for {
			switch strings.ToLower(blocks) {

			case "1", "sure":
				fmt.Println(" ")
				fmt.Println("drop your encrypted message")
				fmt.Print("here: ___")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if input == "" {
					fmt.Println(" ")
					fmt.Println("you did not drop any message")
					fmt.Println("try again")
					fmt.Println(" ")
					continue
				}
				tokens := Tokenize(input)
				fmt.Println(" ")
				fmt.Printf("result: %v\n", HexaDeca(tokens))
				continue

			case "2", "nahh":

				fmt.Println("ok then see you later")
				return

			default:
				fmt.Println(" ")
				fmt.Println("it's either <sure> or <nahh>")
				fmt.Println("try again")
				fmt.Println(" ")
			}
			break
		}
	}
}
