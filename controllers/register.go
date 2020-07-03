package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/olufekosamuel/blog-api/auth"
	"github.com/olufekosamuel/blog-api/helpers"
	"github.com/olufekosamuel/blog-api/models"
	"github.com/lithammer/shortuuid"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "POST":

}