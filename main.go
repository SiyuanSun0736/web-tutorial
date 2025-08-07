package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello, World!")
}
func main() {

	http.HandleFunc("/process", process)

	http.ListenAndServe("localhost:8080", nil)

}
