package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle   string
	Description string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle:   "Guessing Game",
			Description: "Guess a number between 1 and 6.",
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
