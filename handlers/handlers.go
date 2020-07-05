package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/olufekosamuel/blog-api/auth"
	"github.com/olufekosamuel/blog-api/controllers"
)

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

func authRouter() http.Handler {
	r := chi.NewRouter()
	// Middleware with access rules for router.
	r.Use(checkJWT)
	r.Post("/post", controllers.CreatePost)   //POST /create post
	r.Delete("/post", controllers.DeletePost) //DELETE /delete a particular post in blog
	r.Put("/post", controllers.EditPost)      //PUT /edit a particular post in blog

	return r
}

/*
cleaned up code to check JWT token is still valid before making access into the route
*/
func checkJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		err := auth.TokenValid(r)

		if err != nil {
			http.Error(w, fmt.Sprintf(`{"status":"error","error":true,"msg":%s}`, "Unathorized"), 401)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func SetupRouter() *chi.Mux {

	r := chi.NewRouter()

	setupMiddleware(r)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", controllers.Login)              //POST /login
		r.Post("/register", controllers.Register)        //POST /register
		r.Get("/post", controllers.GetPost)              //GET /get all post in blog
		r.Get("/post/detail", controllers.GetPostDetail) //GET /get a post detail
		r.Mount("/", authRouter())
	})

	return r
}
