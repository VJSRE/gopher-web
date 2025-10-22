package views

import (
	"bytes"
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField is not implemented")
		},
	},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("Error parseFS template: %w", err)
	}
	return Template{tpl}, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("Error parsing template: %w", err)
	}
	return Template{tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Println("cloning template:%v", err)
		http.Error(w, "Error while cloning template", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer // this will store the HTML files in the buffer which might cause performance issue.
	err = tpl.Execute(&buf, data)
	if err != nil {
		fmt.Printf("executing template: %v", err)
		http.Error(w, "There is an error while executing the template", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
