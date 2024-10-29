package repositories

import (
	"database/sql"
	"log"
	"nexu-api/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	var errDB error
	db, errDB = sql.Open("postgres", dsn)
	if errDB != nil {
		log.Fatal(errDB)
	}

	if errDB = db.Ping(); errDB != nil {
		log.Fatal(errDB)
	}

}

type ModelRepository interface {
	GetAllBrands() ([]models.Brand, error)
	GetModelsByBrandID(brandID int) ([]models.Model, error)
	UpdateModel(model models.Model) error
	CreateBrand(brand models.Brand) (models.Brand, error)
	CreateModelForBrand(brandID int, model models.ModelDB) (models.ModelDB, error)
	GetAllModels(filter models.Filter) ([]models.Model, error)
}

type modelRepo struct{}

var ModelRepo ModelRepository = &modelRepo{}

func (r *modelRepo) GetAllBrands() ([]models.Brand, error) {
	rows, err := db.Query("SELECT brands.id id, brands.name name, ROUND(AVG(models.average_price), 2) average_price  FROM brands INNER JOIN models ON brands.id = models.brand_id GROUP BY brands.id, brands.name ORDER BY brands.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var brands []models.Brand
	for rows.Next() {
		var brand models.Brand
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.AveragePrice); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

func (r *modelRepo) CreateBrand(brand models.Brand) (models.Brand, error) {
	err := db.QueryRow("INSERT INTO brands (name) VALUES ($1) RETURNING id", brand.Name).Scan(&brand.ID)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"brands_name_key\"" {
			return models.Brand{ID: 0}, err
		}
		return models.Brand{}, err
	}
	return brand, nil
}

func (r *modelRepo) GetModelsByBrandID(brandID int) ([]models.Model, error) {
	rows, err := db.Query("SELECT id, name, average_price FROM models WHERE brand_id = $1", brandID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelsi []models.Model
	for rows.Next() {
		var model models.Model
		if err := rows.Scan(&model.ID, &model.Name, &model.AveragePrice); err != nil {
			return nil, err
		}
		modelsi = append(modelsi, model)
	}
	return modelsi, nil
}

func (r *modelRepo) CreateModelForBrand(brandID int, model models.ModelDB) (models.ModelDB, error) {
	err := db.QueryRow("INSERT INTO models (brand_id, name, average_price) VALUES ($1, $2, $3) RETURNING id", brandID, model.Name, model.AveragePrice).Scan(&model.ID)
	if err != nil {
		if err.Error() == "pq: insert or update on table \"models\" violates foreign key constraint \"models_brand_id_fkey\"" {
			return models.ModelDB{ID: 0}, err
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"models_brand_id_name_key\"" {
			return models.ModelDB{ID: -1}, err
		}
		return models.ModelDB{}, err
	}
	model.BrandID = brandID
	return model, nil
}

func (r *modelRepo) GetAllModels(filter models.Filter) ([]models.Model, error) {
	rows, err := db.Query("SELECT id, name, average_price FROM models WHERE average_price > $1 AND average_price < $2", filter.Greater, filter.Lower)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modelist []models.Model
	for rows.Next() {
		var model models.Model
		if err := rows.Scan(&model.ID, &model.Name, &model.AveragePrice); err != nil {
			return nil, err
		}
		modelist = append(modelist, model)
	}
	return modelist, nil
}

func (r *modelRepo) UpdateModel(model models.Model) error {
	_, err := db.Exec("UPDATE models SET average_price = $1 WHERE id = $2", model.AveragePrice, model.ID)
	return err
}
