package main

import (
	"context"
	"log"

	"github.com/Jamescog/birth-buddy/apis/users"
	"github.com/Jamescog/birth-buddy/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

func NewServer() *Server {
	return &Server{
		http: gin.Default(),
	}
}

func (s *Server) InitDatabase() {
	client, err := ent.Open("sqlite3", "file:database.db?_fk=1")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	s.db = client
	log.Println("Database initialized successfully")
}

func (s *Server) CloseDatabase() {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			log.Printf("Failed to close database: %v", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}

func setupRoutes(r *gin.Engine, db *ent.Client) {
	api := r.Group("/api/v1")
	users.RegisterUserRoutes(api, db)
	// birthdays.RegisterBirthdayRoutes(api, db)
}
