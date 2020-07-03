package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/olufekosamuel/blog-api/controllers"
	"github.com/olufekosamuel/blog-api/helpers"
)

func Test(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Works", 200)
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
	http.HandleFunc("/", Test)
	http.HandleFunc("/v1/login", controllers.Login)
	http.HandleFunc("/v1/register", controllers.Register)

	port = fmt.Sprintf(":%s", port)

	fmt.Println(fmt.Sprintf("application is running on port %s", port))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
