package main

import (
	"fmt"
	"github.com/VJSRE/lenslocked/controllers"
	"github.com/VJSRE/lenslocked/templates"
	"github.com/VJSRE/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		fmt.Printf("parsing template: %v", err)
		http.Error(w, "There is an error while parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", r)
}
