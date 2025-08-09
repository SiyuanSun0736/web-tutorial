package controller

import (
	"html/template"
	"net/http"
)

func registerHomeRoutes() {
	http.HandleFunc("/home", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "template/home.html")
	t.ExecuteTemplate(w, "layout", "Hello World")
}
