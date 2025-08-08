package main

import (
	"net/http"
	"text/template"
)

func main() {
	// 解析所有模板
	tmpl := template.Must(template.ParseFiles("layout.html", "home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 执行名为 "home.html" 的模板
		tmpl.ExecuteTemplate(w, "home.html", nil)
	})

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/js/", http.FileServer(http.Dir("wwwroot")))

	http.ListenAndServe(":8080", nil)
}
