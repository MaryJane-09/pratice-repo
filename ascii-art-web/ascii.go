package main

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"net/http"
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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("input")
		bannerName := r.FormValue("banner")

		fmt.Println(input)
		fmt.Println(bannerName)

		banner, err := asciiart.BannerCheck("banners/" + bannerName)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = asciiart.ValidateInput(input)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		result := asciiart.GenerateArt(input, banner)

		err = tmpl.Execute(w, Data{
			Result: result,
		})
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
