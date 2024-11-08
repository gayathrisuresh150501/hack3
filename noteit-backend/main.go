package main

import (
	"hack3/config"
	"hack3/handlers"
	"log"

	"github.com/savsgio/atreugo/v11"
)

func main() {
	config.ConnectDB()

	config := atreugo.Config{
		Addr: ":8080",
	}

	server := atreugo.New(config)

	server.POST("/notes", handlers.CreateNote)
	server.GET("/notes", handlers.GetNotes)

	log.Println("Server started on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
