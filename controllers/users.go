package controllers

import (
	"github.com/VJSRE/lenslocked/views"
	"net/http"
)

type Users struct {
	Template struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Template.New.Execute(w, nil)
}
