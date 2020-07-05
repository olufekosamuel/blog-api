package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olufekosamuel/blog-api/auth"
	"github.com/olufekosamuel/blog-api/helpers"
	"github.com/olufekosamuel/blog-api/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json") // set response content type

	if (*r).Method == "POST" {

		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user) // decode the from json into our models.User struct

		msg, err := helpers.Validate(user) // validate request body

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, err.Error()), 400)
			return
		}

		if len(msg) != 0 {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, msg), 400)
			return
		}

		hashed, err := helpers.HashPassword(user.Password) // hash password
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, err), 400)
			return
		}

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, err), 400)
			return
		}
		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO users(email,firstname,lastname,password,createdat) VALUES('%s','%s','%s','%s','%s') RETURNING id`, user.Email, user.Firstname, user.Lastname, hashed, time.Now().Format(time.RFC3339))

		rows, err := db.Query(query)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, "user already exists"), 400)
			return
		}

		for rows.Next() {
			rows.Scan(&user.ID)
		}

		// create jwt
		tokenStr, err := auth.CreateToken(user.Email, user.ID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "problem generating token"), 400)
			return
		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.Response{
			Msg:    "User created successfully",
			Status: "success",
			Error:  false,
			Token:  tokenStr,
		})
		return
	}
	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "method not allowed"), 400)
	return
}
