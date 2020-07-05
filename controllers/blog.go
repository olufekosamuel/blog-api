package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ichtrojan/thoth"
	"github.com/olufekosamuel/blog-api/auth"
	"github.com/olufekosamuel/blog-api/helpers"

	"github.com/olufekosamuel/blog-api/models"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "GET" {

		posts := make([]*models.Post, 0)

		file, err := thoth.Init("log")

		if err != nil {
			log.Fatal(err)
		}

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			file.Log(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`SELECT id, user_id, title, published_date, updatedat FROM posts`)
		rows, err := db.Query(query)

		if err != nil {
			file.Log(err)
		}

		defer rows.Close()

		for rows.Next() {
			post := new(models.Post)
			if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.PublishedAt, &post.UpdatedAt); err != nil {
				file.Log(err)
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

func GetPostDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "GET" {

		var post models.Post
		json.NewDecoder(r.Body).Decode(&post)

		if post.ID <= 0 {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Bad Request"), 400)
			return
		}

		file, err := thoth.Init("log")

		if err != nil {
			log.Fatal(err)
		}

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			file.Log(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`SELECT * FROM posts WHERE id=$1`)
		rows, err := db.Query(query, post.ID)

		if err != nil {
			file.Log(err)
		}

		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.PublishedAt, &post.UpdatedAt); err != nil {
				file.Log(err)
			}
		}

		if post.Title == "" {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Bad Request"), 400)
			return
		}

		// if there are no errors, return json response object
		json.NewEncoder(w).Encode(models.ResponsePost{
			Post:   post,
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

	file, err := thoth.Init("log")

	if err != nil {
		log.Fatal(err)
	}

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
			file.Log(err)
		}
		defer db.Close()

		query := `INSERT INTO posts(user_id,title,content,published_date,updatedat) VALUES($1,$2,$3,$4,$5);`
		_, err = db.Exec(query, new_id, post.Title, post.Content, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

		if err != nil {
			file.Log(err)
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":"%s"}`, err.Error()), 400)
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
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "PUT" {
		user_id, err := auth.ExtractTokenId(r)

		file, err := thoth.Init("log")

		if err != nil {
			log.Fatal(err)
		}

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

		db, err := helpers.InitDB() // connect to database
		if err != nil {
			file.Log(err)
		}
		defer db.Close()

		query := `UPDATE posts SET title = $1, content = $2, updatedat = $3 WHERE id = $4 AND user_id = $5 returning id`

		rows, err := db.Query(query, post.Title, post.Content, time.Now().Format(time.RFC3339), post.ID, new_user_id)

		if err != nil {
			file.Log(err)
			http.Error(w, fmt.Sprintf(`{"status":"error","msg":"%s"}`, err.Error()), 400)
			return
		}

		var id int

		for rows.Next() {
			rows.Scan(&id)
		}

		if id == 0 {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Bad Request"), 400)
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
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "DELETE" {

		file, err := thoth.Init("log")

		if err != nil {
			log.Fatal(err)
		}
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
			file.Log(err)
		}
		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM posts WHERE id = '%d' AND user_id = '%d' returning id `, post.ID, new_user_id)
		rows, err := db.Query(query)

		if err != nil {
			file.Log(err)
			http.Error(w, fmt.Sprintf(`{"status":"error","msg":"%s"}`, err.Error()), 400)
			return
		}

		var id int

		for rows.Next() {
			rows.Scan(&id)
		}

		if id == 0 {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Bad Request"), 400)
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
