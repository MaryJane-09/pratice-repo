package main

import (
	"fmt"
	"time"
)

func main() {
	word := "hello world\n"

	for _, char := range word {
		fmt.Print(string(char))
		time.Sleep(200 * time.Millisecond)
	}
}

