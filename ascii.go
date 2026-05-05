package main

import(
	"fmt"
	"os"
	"strings"
)


func main() {
	data, _ := os.ReadFile("standard.txt")


	startline := strings.Split(string(data), "\n")

	//fmt.Println(startline)

	for i, char := range startline {
		if i >= 235 && i <= 243 {
			fmt.Println(char, i)
		}
		//fmt.Println(char,i)
	}
}