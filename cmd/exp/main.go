package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Vijay Koneru",
		Age:  40,
		Meta: UserMeta{
			Visits: 1000,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
