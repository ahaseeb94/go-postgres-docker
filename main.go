package main

import (
	"go-postgres-docker/src/application"
	"go-postgres-docker/src/infrastructure/persistence"
	"log"
	"os"
)

func main() {
	// Load environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	appPort := os.Getenv("APP_PORT")

	// Initialize DB
	persistence.ConnectDB(host, port, user, password, dbName)
	persistence.Migrate()

	// Setup routes
	router := application.SetupRoutes()

	// Run the server
	log.Printf("Server running on port %s", appPort)
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
