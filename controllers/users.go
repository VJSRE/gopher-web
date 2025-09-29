package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Template struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Template.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
