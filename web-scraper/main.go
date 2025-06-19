package main

import (
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	router := chi.NewRouter() // create a new router object
	router.Use(cors.Handler(cors.Options{
			AllowedOrigins:	[]string{"https://*", "http://*"}, // Allow all origins 
			AllowedMethods: []string{"GET", "POST"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{},
			AllowCredentials: false,
			MaxAge: 300,
			Debug: true,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)


	server := &http.Server{
		Handler: router,
		Addr:	":" + port,
	}

	log.Printf("Server is starting on port %s...", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	if port == "" {
		log.Fatal("PORT is not found")
	} else {
		log.Default().Println("PORT is set to: ", port)
	}
}