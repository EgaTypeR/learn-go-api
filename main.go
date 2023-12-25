package main

import (
	"example/Go-Api-Tutorial/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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
