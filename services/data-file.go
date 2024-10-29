package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type BrandData struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	AveragePrice int    `json:"average_price"`
	BrandName    string `json:"brand_name"`
}

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

	fillData(db)
}

func fillData(db *sql.DB) {
	file, err := os.ReadFile("models.json")
	if err != nil {
		panic(err)
	}

	var brandsData []BrandData
	if err := json.Unmarshal(file, &brandsData); err != nil {
		panic(err)
	}
	for _, brandData := range brandsData {

		var id int

		if db.QueryRow("SELECT id FROM brands WHERE name = $1", brandData.BrandName).Scan(&id) == sql.ErrNoRows {
			db.QueryRow("INSERT INTO brands (name) VALUES ($1) RETURNING id", brandData.BrandName).Scan(&id)
		}

		db.QueryRow("INSERT INTO models (brand_id, name, average_price) VALUES ($1, $2, $3)", id, brandData.Name, brandData.AveragePrice)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data inserted successfully")
}
