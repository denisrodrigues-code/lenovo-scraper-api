package model

type StorageOption struct {
	Size  string  `json:"size"`
	Price float64 `json:"price"`
}

type Product struct {
	Name        string          `json:"name"`
	Brand       string          `json:"brand"`
	Description string          `json:"description"`
	Rating      int             `json:"rating"`
	Reviews     int             `json:"reviews"`
	Image       string          `json:"image"`
	URL         string          `json:"url"`
	BasePrice   float64         `json:"base_price"`
	Storages    []StorageOption `json:"storages"`
}
