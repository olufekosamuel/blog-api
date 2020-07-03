package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olufekosamuel/blog-api/auth"

	"github.com/olufekosamuel/blog-api/models"
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := auth.TokenValid(r)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 400)
	}

	if (*r).Method == "POST" {
		var post models.Post
		_ = json.NewDecoder(r.Body).Decode(&post)

	}

}
