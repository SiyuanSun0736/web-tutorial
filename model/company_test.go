package model

import (
	"testing"
)

func TestCompanyTypeCorrect(t *testing.T) {
	c := Company{
		ID:      1,
		Name:    "Innovations Tech",
		Country: "USA",
	}

	companyType := c.GetCompanyType()

	if companyType != "Technology" {
		t.Errorf("Expected Technology, got %s", companyType)
	}
}
