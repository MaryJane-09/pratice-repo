package main

import (
	"fmt"
	"html/template"
	"net/http"
)

	var tmpl = template.Must(template.ParseFiles("templates/index.html"))

	


func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r, "/ascii-art", http.StatusSeeOther)
	
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	fmt.Println(input)
	fmt.Println(banner)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("server running at port :8080")

	http.ListenAndServe(":8080", nil)
}
