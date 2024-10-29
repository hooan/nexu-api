package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nexu-api/models"
	"nexu-api/repositories"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllBrands() ([]models.Brand, error) {
	args := m.Called()
	return args.Get(0).([]models.Brand), args.Error(1)
}

func (m *MockRepository) GetModelsByBrandID(brandID int) ([]models.Model, error) {
	args := m.Called(brandID)
	return args.Get(0).([]models.Model), args.Error(1)
}

func (m *MockRepository) UpdateModel(model models.Model) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockRepository) CreateBrand(brand models.Brand) (models.Brand, error) {
	args := m.Called(brand)
	return args.Get(0).(models.Brand), args.Error(1)
}

func (m *MockRepository) CreateModelForBrand(brandID int, model models.ModelDB) (models.ModelDB, error) {
	args := m.Called(brandID, model)
	return args.Get(0).(models.ModelDB), args.Error(1)
}

func (m *MockRepository) GetAllModels(filter models.Filter) ([]models.Model, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Model), args.Error(1)
}

func TestMain(m *testing.M) {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loadiang .env file")
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestGetBrandsHandler(t *testing.T) {
	mockRepo := new(MockRepository)
	repositories.ModelRepo = mockRepo

	mockBrands := []models.Brand{
		{ID: 1, Name: "Brand A", AveragePrice: 100000},
		{ID: 2, Name: "Brand B", AveragePrice: 200000},
	}
	mockRepo.On("GetAllBrands").Return(mockBrands, nil)

	req, err := http.NewRequest("GET", "/brands", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBrandsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []models.Brand
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, mockBrands, response)
}

func TestGetModelsHandler(t *testing.T) {
	mockRepo := new(MockRepository)
	repositories.ModelRepo = mockRepo

	mockModels := []models.Model{
		{ID: 1, Name: "Model A", AveragePrice: 120000},
		{ID: 2, Name: "Model B", AveragePrice: 200000},
	}
	mockRepo.On("GetAllModels", models.Filter{Greater: 120000, Lower: 500000}).Return(mockModels, nil)

	req, err := http.NewRequest("GET", "/models?greater=120000&lower=500000", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetModelsHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response []models.Model
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, mockModels, response)
}

func TestUpdateModelHandler(t *testing.T) {
	mockRepo := new(MockRepository)
	repositories.ModelRepo = mockRepo

	model := models.Model{ID: 1, Name: "Updated Model", AveragePrice: 150000}
	mockRepo.On("UpdateModel", model).Return(nil)

	body, err := json.Marshal(model)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/models/1", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/models/{id}", UpdateModelHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateBrandHandler(t *testing.T) {
	mockRepo := new(MockRepository)
	repositories.ModelRepo = mockRepo

	brand := models.Brand{Name: "New Brand"}
	createdBrand := models.Brand{ID: 1, Name: "New Brand"}
	mockRepo.On("CreateBrand", brand).Return(createdBrand, nil)

	body, err := json.Marshal(brand)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/brands", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBrandHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.Brand
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, createdBrand, response)
}
