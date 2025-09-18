package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome World !!</h1>")

}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
