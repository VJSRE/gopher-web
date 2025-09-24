package main

import (
	"fmt"
	"github.com/VJSRE/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
)

func executeTemplate(w http.ResponseWriter, filePath string) {
	t, err := views.Parse(filePath)
	if err != nil {
		fmt.Printf("parsing template: %v", err)
		http.Error(w, "There is an error while parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func fqlHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", fqlHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}
