package main

import (
	"fmt"
	"log"
	"net/http"
	"nexu-api/handlers"
	"nexu-api/repositories"
	"os"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Database connection
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	//config := &postgres.Config{MigrationsTable: "migrations", DatabaseName: os.Getenv("DB_NAME"), SchemaName: "public"}
	// Run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating migration driver: %v", err)
	}
	fmt.Println("Running migrations")
	m, err := migrate.NewWithDatabaseInstance("file://migrations", os.Getenv("DB_NAME"), driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Initialize the repository with the database connection
	repositories.InitDB(db)

	r := mux.NewRouter()

	r.HandleFunc("/brands", handlers.GetBrandsHandler).Methods("GET")
	r.HandleFunc("/brands/{id}/models", handlers.GetBrandModelsHandler).Methods("GET")
	r.HandleFunc("/brands", handlers.CreateBrandHandler).Methods("POST")
	r.HandleFunc("/brands/{id}/models", handlers.CreateBrandModelHandler).Methods("POST")
	r.HandleFunc("/models/{id}", handlers.UpdateModelHandler).Methods("PUT")
	r.HandleFunc("/models", handlers.GetModelsHandler).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
