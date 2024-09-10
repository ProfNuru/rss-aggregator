package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	// Get PORT form .env file
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is nor found in the environment")
	}

	// Create router
	router := chi.NewRouter()

	// Setup route
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Hook handlers to router
	v1Router := chi.NewRouter()
	// handleFunc allows all http methods
	// v1Router.HandleFunc("/healthz", handlerReadiness)

	// Only handle GET requests on this route
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)

	// Mount v1 router into main router
	// The full path then becomes /v1/healthz
	router.Mount("/v1", v1Router)

	// Connect router to http server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// Listen
	fmt.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
