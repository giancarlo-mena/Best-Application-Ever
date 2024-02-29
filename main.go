package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type PageData struct {
	PageTitle   string
	Description string
	Answer      string
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle:   "Guessing Game",
			Description: "Guess a number between 1 and 6.",
			Answer:      "Hello World!",
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
