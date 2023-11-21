package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "https://*"},
		AllowedMethods:   []string{"GeT", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	v1Router := chi.NewRouter()
	v1Router.Get("/err", handlerErr)
	v1Router.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1Router)
	log.Printf("🚀 Server started on : %v", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
