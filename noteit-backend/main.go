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

	server.UseBefore(middleware.AuthMiddleware())
	server.UseBefore(middleware.CORSMiddleware())

	server.POST("/api/notes", handlers.CreateNote)
	server.GET("/api/notes/{id}", handlers.GetNote)
	server.PUT("/api/notes/{id}", handlers.UpdateNote)
	server.DELETE("/api/notes/{id}", handlers.DeleteNote)
	server.POST("/plan", handlers.AddPlan)
	server.GET("/plan/{uid}", handlers.GetPlan)

	log.Println("Starting server on :8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
