package main

import (
	config "main/configs"
	"main/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	config.Connect()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	router.Get("/dogs", handlers.GetDogs)
	router.Get("/dogs/{id}", handlers.GetDog)
	router.Post("/dogs", handlers.CreateDog)
	router.Put("/dogs/{id}", handlers.UpdateDog)
	router.Delete("/dogs/{id}", handlers.DeleteDog)

	http.ListenAndServe(":3000", router)
}
