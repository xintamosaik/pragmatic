package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Heading string
	Content string
}

func main() {
	tmpl := template.Must(template.ParseFiles(
		"base.html",
		"head.html",
		"navigation.html",
		"main.html",
	))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "My Page Title",
			Heading: "Welcome to My Website",
			Content: "This is a sample paragraph.",
		}
		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

    // Serve static files.
    fs := http.FileServer(http.Dir("dist"))
    http.Handle("/dist/", http.StripPrefix("/dist/", fs))

	log.Println("Listening on :8091...")
	if err := http.ListenAndServe(":8091", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
