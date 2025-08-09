package main

import (
	"github.com/SiyuanSun0736/web-tutorial/controller"
	"net/http"
)

func main() {
	controller.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
