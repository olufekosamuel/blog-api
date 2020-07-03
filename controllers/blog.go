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

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id, err := auth.ExtractTokenId(r)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 400)
		return
	}

	fmt.Println(id)

	if (*r).Method == "POST" {
		var post models.Post
		_ = json.NewDecoder(r.Body).Decode(&post)

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO posts(user_id,title,content,published_date,updatedat) VALUES('%s','%s','%s','%s','%s');`, id, post.Title, post.Content, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

		_, err = db.Exec(query)

		if err != nil {
			fmt.Println(err)
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, "an error occured while creating post"), 400)
			return
		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.Response{
			Msg:    "Post created successfully",
			Status: "success",
			Error:  false,
		})
		return

	}

}
