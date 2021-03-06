package controllers

import (
	"net/http"
	"testing"
)

func Test_UserProfile_Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := executeRequest(req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("return status %v was expectiong %v", status, http.StatusOK)
	}
}
