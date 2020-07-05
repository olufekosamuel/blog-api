package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/olufekosamuel/blog-api/handlers"
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

	file, err := thoth.Init("log")

	if err != nil {
		log.Fatal(err)
	}

	createTableErr := helpers.CreateTables()

	if createTableErr != nil {
		panic(createTableErr)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//setup router and middlewares with chi
	r := handlers.SetupRouter()

	port = fmt.Sprintf(":%s", port)

	fmt.Println(fmt.Sprintf("application is running on port %s", port))

	if err := file.Serve("/logs", "12345"); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(port, r); err != nil {
		file.Log(err)
	}
}
