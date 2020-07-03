package handlers

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

func SetupRouter() *chi.Mux {

	r := chi.NewRouter()

	setupMiddleware(r)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", controllers.Login)       //POST /login
		r.Post("/register", controllers.Register) //POST /register
		r.Post("/post", controllers.Post)         //POST /create
		r.Get("/post", controllers.Post)          //POST /create
		/*
			r.Get("/{phonenumber}", getContact)       //POST /contacts/0147344454
			r.Post("/", addContact)                   //POST /contacts
		*/
	})

	return r
}
