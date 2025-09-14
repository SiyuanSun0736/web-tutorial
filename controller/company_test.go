package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SiyuanSun0736/web-tutorial/model"
)

func TestHandleCompanyCorrect(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/company", nil)
	w := httptest.NewRecorder()

	handleCompany1(w, r)

	result, _ := io.ReadAll(w.Result().Body)

	c := model.Company{}
	json.Unmarshal(result, &c)

	if c.ID != 1 {
		t.Errorf("Expected ID 1, got %d", c.ID)
	}
}
