package main

import (
	"encoding/json"
	"github.com/SiyuanSun0736/web-tutorial/middleware"
	"net/http"
)

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      1,
			Name:    "Google",
			Country: "USA",
		}
		enc := json.NewEncoder(w)
		enc.Encode(c)

	})
	http.ListenAndServe(":8080", &middleware.TimeoutMiddleware{
		Next: new(middleware.AuthMiddleware),
	})
}
