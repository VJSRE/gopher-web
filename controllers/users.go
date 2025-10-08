package controllers

import (
	"fmt"
	"github.com/VJSRE/lenslocked/models"
	"net/http"
)

type Users struct {
	Template struct {
		New Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Template.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created successfully:\n%s", user)
}
