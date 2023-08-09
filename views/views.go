package views

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("views/mahasiswa/test.html"))

	tmpl.Execute(w, nil)

}

func WE(w http.ResponseWriter, r *http.Request) {
	http.Handle("/views/assets/", http.StripPrefix("/views/assets/", http.FileServer(http.Dir("views/assets"))))
}
