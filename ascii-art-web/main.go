package main

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Data struct {
	Result string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ascii-art", http.StatusSeeOther)

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		err := tmpl.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("input")
		if strings.ContainsAny(input, "\\n") {
			fmt.Println(true)
		}

		bannerName := r.FormValue("banner")
		fmt.Println(input)
		fmt.Println(bannerName)

		banner, err := asciiart.BannerCheck("banners/" + bannerName)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Printf("%q\n", input)
		fmt.Println(asciiart.GenerateArt(input, banner))

		_, err = asciiart.ValidateInput(input)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		result := asciiart.GenerateArt(input, banner)

		err = tmpl.Execute(w, Data{
			Result: result,
		})
		if err != nil {
			w.WriteHeader(500)
			return
		}

	}

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("server running at port :8080")

	http.ListenAndServe(":8080", nil)
}
