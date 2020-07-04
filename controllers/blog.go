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

func GetPost(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "GET" {

		posts := make([]*models.Post, 0)

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`SELECT * FROM posts`)
		rows, err := db.Query(query)

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			post := new(models.Post)
			if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.PublishedAt, &post.UpdatedAt); err != nil {
				panic(err)
			}
			posts = append(posts, post)

		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.Response{
			Posts:  posts,
			Status: "success",
			Error:  false,
		})

		return

	}

	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Method not allowed"), 405)
	return

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id, err := auth.ExtractTokenId(r)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 401)
		return
	}

	var post models.Post
	var new_id int = int(id)
	_ = json.NewDecoder(r.Body).Decode(&post)

	if (*r).Method == "POST" {

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO posts(user_id,title,content,published_date,updatedat) VALUES('%d','%s','%s','%s','%s');`, new_id, post.Title, post.Content, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

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

	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Method not allowed"), 405)
	return

}

func EditPost(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "PUT" {
		user_id, err := auth.ExtractTokenId(r)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 401)
			return
		}

		/* instead of passing parameters through url, I just pass in ID */
		//id := chi.URLParam(r, "id")

		//new_id, err := strconv.Atoi(id)

		var new_user_id int = int(user_id)
		var post models.Post

		json.NewDecoder(r.Body).Decode(&post)

		fmt.Println(post)

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`UPDATE posts SET title= '%s', content='%s' WHERE id = '%d' AND user_id = '%d' `, post.Title, post.Content, post.ID, new_user_id)
		_, err = db.Exec(query)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","msg":"%s"}`, err.Error()), 400)
			return
		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.Response{
			Msg:    "Post updated successfully",
			Status: "success",
			Error:  false,
		})

		return

	}

	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Method not allowed"), 405)
	return
}

func DeletePost(w http.ResponseWriter, r *http.Request) {

	if (*r).Method == "DELETE" {
		w.Header().Set("content-type", "application/json")
		user_id, err := auth.ExtractTokenId(r)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 401)
			return
		}

		var post models.Post
		var new_user_id int = int(user_id)
		json.NewDecoder(r.Body).Decode(&post)

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			panic(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM posts WHERE id = '%d' AND user_id = '%d' `, post.ID, new_user_id)
		_, err = db.Exec(query)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","msg":"%s"}`, err.Error()), 400)
			return
		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.Response{
			Msg:    "Post deleted successfully",
			Status: "success",
			Error:  false,
		})

		return

	}

	http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Method not allowed"), 405)
	return

}
