package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello MJ")
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}