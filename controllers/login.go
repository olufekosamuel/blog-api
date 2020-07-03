package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olufekosamuel/blog-api/auth"
	"github.com/olufekosamuel/blog-api/helpers"
	"github.com/olufekosamuel/blog-api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") // set response content type

	if (*r).Method == "POST" {

		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user) // decode the from json into our models.User struct

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`SELECT password FROM users WHERE email = '%s'`, user.Email)
		rows, err := db.Query(query)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var password string

		for rows.Next() {
			rows.Scan(&password)
		}

		fmt.Println(helpers.CheckPasswordHash(user.Password, password), password, user.Password)

		if helpers.CheckPasswordHash(user.Password, password) == true {
			// create jwt
			tokenStr, err := auth.CreateToken(user.Email)

			if err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "problem generating token"), 400)
				return
			}

			json.NewEncoder(w).Encode(models.Response{
				Status: "success",
				Error:  false,
				Token:  tokenStr,
			})
			return
		}

		json.NewEncoder(w).Encode(models.Response{
			Status: "failed",
			Error:  true,
			Msg:    "Invalid login credentials",
		})
		return
	}
	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "method not allowed"), 400)
	return
}
