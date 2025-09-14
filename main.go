package main

import (
	"net/http"

	"github.com/SiyuanSun0736/web-tutorial/controller"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServeTLS("localhost:8080", "server.crt", "server.key", nil)
}
