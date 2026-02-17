package controller

import (
	"encoding/json"
	"net/http"

	"github.com/denisrodrigues-code/lenovo-scraper-api/model"
	"github.com/denisrodrigues-code/lenovo-scraper-api/scraper"
	"github.com/denisrodrigues-code/lenovo-scraper-api/service"
)

func LenovoHandler(w http.ResponseWriter, r *http.Request) {
	links, _ := scraper.GetLenovoProductLinks(
		"https://webscraper.io/test-sites/e-commerce/static/computers/laptops",
	)

	var products []model.Product

	for _, link := range links {
		product, err := scraper.GetProductDetails(link)
		if err == nil {
			products = append(products, *product)
		}
	}

	service.SortByCheapest(products)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
