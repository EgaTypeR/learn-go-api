package main

import (
	"Learn-Go-Api/database"
	"Learn-Go-Api/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	_, err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing DB:", err)
	}

	godotenv.Load(".env")
	fmt.Println("Connected to Database")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", utils.HandlerReadiness)
	v1Router.Get("/err", utils.HandlerErr)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv("PORT"),
	}
	log.Printf("Server is staring on port %v", srv.Addr)
	error := srv.ListenAndServe()
	if error != nil {
		log.Fatal("Error initiating server", error)
	}
}
