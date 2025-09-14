package controller

import (
	"encoding/json"
	"github.com/SiyuanSun0736/web-tutorial/model"
	"net/http"
)

func RegisterCompanyRoutes() {
	http.HandleFunc("/companies", handleCompany1)
}
func handleCompany1(w http.ResponseWriter, r *http.Request) {
	c := model.Company{
		ID:      1,
		Name:    "Innovations Tech",
		Country: "USA",
	}

	enc := json.NewEncoder(w)
	enc.Encode(c)
}
