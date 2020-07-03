package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olufekosamuel/blog-api/controllers"
	"github.com/olufekosamuel/blog-api/helpers"
	"github.com/olufekosamuel/blog-api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(
		models.Response{
			Msg:    "Blog api works",
			Status: "success",
			Error:  false,
		},
	)
}

func main() {

	createTableErr := helpers.CreateTables()

	if createTableErr != nil {
		panic(createTableErr)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// endpoints
	http.HandleFunc("/", Index)
	http.HandleFunc("/v1/login", controllers.Login)
	http.HandleFunc("/v1/register", controllers.Register)

	port = fmt.Sprintf(":%s", port)

	fmt.Println(fmt.Sprintf("application is running on port %s", port))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
