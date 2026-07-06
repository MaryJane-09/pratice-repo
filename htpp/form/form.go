package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	input := r.FormValue("input")
	fmt.Println(input)

	switch r.URL.Path {
	case "/":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Fprint(w, err)
		}
		tmpl.Execute(w, nil)

	default:
		fmt.Fprint(w, http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	fmt.Println("Received")
	http.ListenAndServe(":8080", nil)
}
