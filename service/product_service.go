package service

import (
	"sort"

	"github.com/denisrodrigues-code/lenovo-scraper-api/model"
)

func SortByCheapest(products []model.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].BasePrice < products[j].BasePrice
	})
}
