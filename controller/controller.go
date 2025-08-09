package controller

import (
	"net/http"
)

func RegisterRoutes() {
	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/js/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))
	registerHomeRoutes()
	registerContactRoutes()
	registerAboutRoutes()
	registerCompanyRoutes()
}
