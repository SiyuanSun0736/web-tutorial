package controller

import (
	"html/template"
	"net/http"
)

func registerContactRoutes() {

	http.HandleFunc("/contact", handleContact)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "template/contact.html")
	t.ExecuteTemplate(w, "layout", "Contact us")
}
