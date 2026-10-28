package main

import (
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	if err := config.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db := config.GetDB()

	// Close database connection on shutdown
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL connection:", err)
	}
	defer sqlDB.Close()

	// auto-migrate book model
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Init router
	router := mux.NewRouter()
	// register API routes
	routes.RegisterBookStoreRoutes(router)
	// config http server
	server := &http.Server{
		Addr:              ":9010",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	// start server
	log.Println("Starting server on port 9010")
	log.Panicln("MySQL connected successfully")
	log.Println("Database connected and migrated successfully")

	log.Fatal(server.ListenAndServe())
}
