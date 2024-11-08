package main

import (
	"hack3/db"
	"hack3/handlers"
	"hack3/middleware"
	"log"

	"github.com/savsgio/atreugo/v11"
)

func main() {
	err := db.ConnectMongo()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)

	// Use CORS middleware before Auth middleware
	server.UseBefore(middleware.CORSMiddleware())
	server.UseBefore(middleware.AuthMiddleware())

	server.POST("/api/notes/{uid}", handlers.CreateNote)
	server.GET("/api/notes/{id}", handlers.GetNote)
	server.GET("/api/notes", handlers.GetAllNotes)
	server.PUT("/api/notes/{id}", handlers.UpdateNote)
	server.DELETE("/api/notes/{id}", handlers.DeleteNote)
	server.POST("/plan", handlers.AddPlan)
	server.GET("/plan/{uid}", handlers.GetPlan)

	log.Println("Starting server on :8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
