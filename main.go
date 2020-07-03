package main

import (
	"fmt"
	"net/http"
	"os"
)

func Test(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Works", 200)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// endpoints
	http.HandleFunc("/", Test)

	port = fmt.Sprintf(":%s", port)

	fmt.Println(fmt.Sprintf("application is running on port %s", port))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
