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
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
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

func HexaWords(words []string) string {
	var result []string

	for _, w := range words {
		val, err := strconv.ParseInt(w, 16, 64)
		if err == nil {
			result = append(result, strconv.FormatInt(val, 36))
		} else {
			result = append(result, w)
		}
	}

	return strings.Join(result, "")
}

func BinWords(words []string) string {
	var result []string

	for i := 0; i < len(words); i++ {
		val, err := strconv.ParseInt(words[i], 2, 64)
		if err == nil {
			result = append(result, strconv.FormatInt(val, 36))
			continue
		}
		result = append(result, words[i])

	}

	return strings.Join(result, "")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	HexMessages := map[string]string{
		"a": "HEXABET-'A': 33a 109586 2a4 654178d56",
		"b": "HEXABET-'B': 33a 109586 2a4 654178d56",
		"c": "HEXABET-'C': 33a 109586 2a4 654178d56",
		"d": "HEXABET-'D': 33a 109586 2a4 654178d56",
		"e": "HEXABET-'E': 33a 109586 2a4 654178d56",
		"f": "HEXABET-'F': 33a 109586 2a4 654178d56",
	}

	BinMessages := map[string]string{
		"b": "BINABET- '0': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
		"i": "BINABET- '10': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
		"n": "BINABET- '11': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
		"a": "BINABET- '101': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
		"r": "BINABET- '111': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
		"y": "BINABET- '1011': 1100111010 100001001010110000110 1010100100 11001010100000101111000110101010110",
	}

	fmt.Println("GREETINGS AGENT X . WELCOME TO THE LEAGUE OF SHADOWS\n A MESSAGE WILL BE SENT TO YOU IN A CODED LANGUAGE\n DON'T WORRY THE MACHINE WILL ALSO HELP YOU TO DE-ENCRYPT IT")

	for {
		fmt.Println("WHAT LANGUAGE DO YOU WANT THE MESSAGE IN:")
		fmt.Println("1. HEXA")
		fmt.Println("2. BINARY")
		fmt.Print("> ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))

		switch choice {

		case "1", "hexa":

			fmt.Print("NICE CHOICE")

			for {

				fmt.Println("PICK A LETTER WITHIN THE RANGE OF HEXABETS(hexadecimal alphabets)")
				fmt.Print("> ")

				letter, _ := reader.ReadString('\n')
				letter = strings.TrimSpace(strings.ToLower(letter))

				msg, ok := HexMessages[letter]
				if !ok {
					fmt.Println("YOU ARE TO PICK HEXABETS ONLY (hexadecimal alphabets)")
					fmt.Println("TRY AGAIN")
					fmt.Println(" ")

					continue
				}

				fmt.Println(msg)
				fmt.Println()

				fmt.Println("DO YOU WANT TO DE-ENCRYPT IT? (yes/no)")
				ans, _ := reader.ReadString('\n')
				ans = strings.TrimSpace(strings.ToLower(ans))

				if ans == "yes" {
					tokens := Tokenize(msg)
					fmt.Println()
					fmt.Printf("DE-ENCRYPTED: %v\n", HexaWords(tokens))
					fmt.Println()

				} else if ans == "no" {
					fmt.Println("SEE YOU LATER THEN")
					fmt.Println()
					return

				} else {
					fmt.Println("INVALID INPUT - LOGGED OUT")
					fmt.Println()
					return
				}
				break
			}

		case "2", "bin":
			fmt.Print("NICE CHOICE")

			for {

				fmt.Println("PICK A LETTER WITHIN THE RANGE OF LETTERs THAT MAKES UP THE LANGUAGE YOU ARE CURRENTLY SPEAKING")
				fmt.Print("> ")

				letter, _ := reader.ReadString('\n')
				letter = strings.TrimSpace(strings.ToLower(letter))

				msg, ok := BinMessages[letter]
				if !ok {
					fmt.Println("YOU ARE TO PICK BINABETS ONLY(letters that makes up the language)")
					fmt.Println("TRY AGAIN")
					fmt.Println(" ")

					continue
				}

				fmt.Println(msg)
				fmt.Println()

				fmt.Println("DO YOU WANT TO DE-ENCRYPT IT? (yes/no)")
				ans, _ := reader.ReadString('\n')
				ans = strings.TrimSpace(strings.ToLower(ans))

				if ans == "yes" {
					tokens := Tokenize(msg)
					fmt.Println()
					fmt.Printf("DE-ENCRYPTED: %v\n", BinWords(tokens))
					fmt.Println()

				} else if ans == "no" {
					fmt.Println("SEE YOU LATER THEN")
					fmt.Println()
					return

				} else {
					fmt.Println("INVALID INPUT - LOGGED OUT")
					fmt.Println()
					return
				}
				break
			}

		default:
			fmt.Println("INVALID CHOICE. ENTER 1 OR 2")
		}
	}
}
