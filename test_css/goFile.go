package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("test_css/assets/welcome.html"))

	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("login.html"))

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/login", login)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
