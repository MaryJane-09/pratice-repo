package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello MJ")

	case "/about":
		fmt.Fprint(w, "This is a website")

	case "/contact":
		fmt.Fprint(w, "Contact us at nzekwechinaza0613@gmail.com")

	default:
		fmt.Fprint(w, http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", pageHandler)
	fmt.Println("Received")
	http.ListenAndServe(":8080", nil)
}
