package models

type Response struct {
	Message string `json:"message"`
}

type Brand struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	AveragePrice float32 `json:"average_price"`
}

type Model struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	AveragePrice float32 `json:"average_price"`
}

type ModelDB struct {
	ID           int     `json:"id"`
	BrandID      int     `json:"brand_id"`
	Name         string  `json:"name"`
	AveragePrice float32 `json:"average_price"`
}

type ModelDesc struct {
	ID           int     `json:"id"`
	Brand        string  `json:"brand_name"`
	Name         string  `json:"name"`
	AveragePrice float32 `json:"average_price"`
}

type Filter struct {
	Greater float32
	Lower   float32
}
