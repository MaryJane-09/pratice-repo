package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Language struct {
	Name     string
	Title    string
	Messages map[string]string
	Base     int
}

func main() {
	runGame()
}

func runGame() {
	reader := bufio.NewReader(os.Stdin)

	languages := map[string]Language{
		"1": {
			Name:  "HEXA",
			Title: "HEXABETS",
			Base:  16,
			Messages: map[string]string{
				"a": "5315 451 35e9 2da14aa!!",
				"b": "333 1765d5f476a 43ea!!",
			},
		},
		"2": {
			Name:  "BINARY",
			Title: "BINABETS",
			Base:  2,
			Messages: map[string]string{
				"b": "101001100010101 10001010001",
				"i": "110110011011011010010001",
			},
		},
	}

	showIntro()

	for {
		lang := chooseLanguage(reader, languages)
		runMission(reader, lang)
	}
}

func showIntro() {
	fmt.Println("GREETINGS AGENT X")
	fmt.Println("WELCOME TO THE LEAGUE OF SHADOWS")
	fmt.Println("A MESSAGE AWAITS DECRYPTION")
	fmt.Println()
}

func chooseLanguage(reader *bufio.Reader, languages map[string]Language) Language {
	for {
		fmt.Println("SELECT LANGUAGE:")
		fmt.Println("1. HEXA")
		fmt.Println("2. BINARY")
		fmt.Print("> ")

		input := readLine(reader)

		if lang, ok := languages[input]; ok {
			return lang
		}

		fmt.Println("INVALID CHOICE.")
		fmt.Println()
	}
}

func runMission(reader *bufio.Reader, lang Language) {
	fmt.Printf("NICE CHOICE, AGENT.\n")
	fmt.Printf("PICK A SYMBOL FROM %s\n", lang.Title)
	fmt.Print("> ")

	key := readLine(reader)

	msg, ok := lang.Messages[key]
	if !ok {
		fmt.Println("UNKNOWN SYMBOL.")
		fmt.Println()
		return
	}

	fmt.Printf("%s '%s': %s\n\n", lang.Title, strings.ToUpper(key), msg)

	fmt.Println("DECRYPT MESSAGE? (yes/no)")
	fmt.Print("> ")

	answer := readLine(reader)

	switch answer {
	case "yes":
		decoded := decodeMessage(msg, lang.Base)
		fmt.Println("DE-ENCRYPTED:", capitalize(decoded))
		fmt.Println()

	case "no":
		fmt.Println("MISSION SKIPPED.")
		fmt.Println()

	default:
		fmt.Println("INVALID RESPONSE.")
		fmt.Println()
	}
}

func decodeMessage(input string, base int) string {
	tokens := tokenize(input)
	var result []string

	for _, token := range tokens {
		value, err := strconv.ParseInt(token, base, 64)
		if err == nil {
			result = append(result, strconv.FormatInt(value, 36))
		} else {
			result = append(result, token)
		}
	}

	return strings.Join(result, "")
}

func tokenize(input string) []string {
	var tokens []string
	var current strings.Builder

	for _, ch := range input {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			current.WriteRune(ch)
		} else {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func Capitalize(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
