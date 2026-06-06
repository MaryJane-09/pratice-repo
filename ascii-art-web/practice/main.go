package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Name    string
	Message string
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Method:", r.Method)
	// fmt.Println("Path:", r.URL.Path)
	// fmt.Println("Host:", r.Host)

	if r.Method == http.MethodPost {
		fmt.Println("Form submitted")
		fmt.Println(r.FormValue("input"))
		fmt.Println(r.FormValue("banner"))
	}

	data := Data{
		Name:    "MJ",
		Message: "Welcome to my webpage",
	}

	switch r.URL.Path {
	case "/":
		tmlp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tmlp.Execute(w, data)
	case "/about":
		fmt.Fprint(w, "this webpage is about printing hello MJ")
	case "/contact":
		fmt.Fprint(w, "contact us at 07036702434")
	default:
		http.NotFound(w, r)
	 }
}

func main() {
	http.HandleFunc("/", pageHandler)

	http.ListenAndServe(":8080", nil)
}
