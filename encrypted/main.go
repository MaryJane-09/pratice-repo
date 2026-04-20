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

	HexMessages := map[string] string {
		"a": "5315 451 35e9 2da14aa!!",
		"b": "333 1765d5f476a 43ea 18b253 80924 555a!!",
		"c": "a 17271e 70df 118c2e 142435 '2628dfa99e36 6dc0 7dd5'",
		"d": "5b926992 6bbbccdc9bdd5 29f 14bfa43 20627160616 5315 2be2716 9b26 258 40ae!!!",
		"e": "a 17271e 70df 118c2e 142435 '9542 118d86 2f3b5ac 196424e 2a4 18b253 92b151a6f7aae c64fa'",
		"f": "dbdc6 66311381 2a4 42c 333 130498c9629 184 258-1f408018e55 2a4 42c?",
	}

	BinMessages := map[string]string{
		"b": "101001100010101 10001010001 11010111101001 10110110100001010010101010 1010111110011110 11011001111010 1001010010001001010100100 11010111101001 11001011010111110100001011001111001100100!!!",
		"i": "110110011011011010010001100100110110 11000111011101000010 110001011001001010011 10100011010101000000001010 1100110011 10011000001001001100011001001011000101001 100001111101010 11000!",
		"n": "1100011001001100010010000010010 1100100110 101001111110111110100: 101110001111110010101 10011011110011010100 1010 11110100111111100011 111101100 1000010111000110010100011 1011001011011111111101100?",
		"a": "1010010111 1101010100 110011010 1010111110011110 101001111110101011111 1010010001111100 101101001101011100011111010100001000?",
		"r": "10111000100111011111001000 101110010011000101001 11110111000000110100 110000101 1010111110011110 11010111101001 10110111011010001110100010 '11111011101011001100111101011100110 101000011011111101111 10010011001100011001011110001100000110110101100 1101011010100101001110011011011100100011'.",
		"y": "1100011001001100010010000010010 1100100110 101001111110111110100: 10010 101111110 1010 10111110001111011100001101 101001111110011010101 10010110000000110110000011 1011000001100001011110100.",
	}

	fmt.Println("GREETINGS AGENT X . WELCOME TO THE LEAGUE OF SHADOWS\n A MESSAGE WILL BE SENT TO YOU IN A CODED LANGUAGE\n DON'T WORRY THE MACHINE WILL ALSO HELP YOU TO DE-ENCRYPT IT")
	fmt.Println()

	for {
		fmt.Println("WHAT LANGUAGE DO YOU WANT THE MESSAGE IN:")
		fmt.Println("1. HEXA")
		fmt.Println("2. BINARY")
		fmt.Print("> ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))

		switch choice {

		case "1", "hexa":

			fmt.Print("NICE CHOICE ")

			for {

				fmt.Println("PICK A LETTER WITHIN THE RANGE OF HEXABETS(hexadecimal alphabets)")
				fmt.Print("> ")

				letter, _ := reader.ReadString('\n')
				letter = strings.TrimSpace(strings.ToLower(letter))

				msg, ok := HexMessages[letter]
				if !ok {
					fmt.Println("YOU ARE TO PICK HEXABETS ONLY (hexadecimal alphabets)")
					fmt.Println("TRY AGAIN")
					fmt.Println()

					continue
				}

				fmt.Printf("HEXABET-'%v': " + msg, strings.ToUpper(letter))
				fmt.Println()

				fmt.Println("DO YOU WANT TO DE-ENCRYPT IT? (YES/NO)")
				fmt.Print("> ")
				ans, _ := reader.ReadString('\n')
				ans = strings.TrimSpace(strings.ToLower(ans))

				if ans == "yes" {
					tokens := Tokenize(msg)
					fmt.Println()
					fmt.Printf("DE-ENCRYPTED: %v\n", Capitalize(HexaWords(tokens)))
					fmt.Println()

				} else if ans == "no" {
					fmt.Println("SEE YOU LATER THEN")
					fmt.Println()
					continue

				} else {
					fmt.Println("INVALID INPUT - LOGGED OUT")
					fmt.Println()
					return
				}
				break
			}

		case "2", "bin":
			fmt.Print("NICE CHOICE ")

			for {

				fmt.Println("PICK A LETTER WITHIN THE RANGE OF LETTERS THAT MAKES UP THE LANGUAGE YOU ARE CURRENTLY SPEAKING")
				fmt.Print("> ")

				letter, _ := reader.ReadString('\n')
				letter = strings.TrimSpace(strings.ToLower(letter))

				msg, ok := BinMessages[letter]
				if !ok {
					fmt.Println("YOU ARE TO PICK BINABETS ONLY(letters that makes up the language)")
					fmt.Println("TRY AGAIN")
					fmt.Println()

					continue
				}

				fmt.Printf("BINABET-'%v': " + msg, strings.ToUpper(letter))
				fmt.Println()

				fmt.Println("DO YOU WANT TO DE-ENCRYPT IT? (YES/NO)")
				fmt.Print(">")
				ans, _ := reader.ReadString('\n')
				ans = strings.TrimSpace(strings.ToLower(ans))

				if ans == "yes" {
					tokens := Tokenize(msg)
					fmt.Println()
					fmt.Printf("DE-ENCRYPTED: %v\n", Capitalize(BinWords(tokens)))
					fmt.Println()

				} else if ans == "no" {
					fmt.Println("SEE YOU LATER THEN")
					fmt.Println()
					continue

				} else {
					fmt.Println("INVALID INPUT - LOGGED OUT")
					fmt.Println()
					return
				}
				break
			}

		default:
			fmt.Println("INVALID CHOICE. ENTER 1 OR 2")
			fmt.Println()
		}
	}
}

func Capitalize(str string) string {
	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, " ")
}
