package scraper

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/denisrodrigues-code/lenovo-scraper-api/model"
	"github.com/denisrodrigues-code/lenovo-scraper-api/utils"
)

func GetProductDetails(url string) (*model.Product, error) {
	client := utils.NewClient()
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	product := &model.Product{
		Name:        strings.TrimSpace(doc.Find("h4").First().Text()),
		Brand:       "Lenovo",
		Description: doc.Find(".description").Text(),
		URL:         url,
	}

	doc.Find(".ratings .glyphicon-star").Each(func(i int, _ *goquery.Selection) {
		product.Rating++
	})

	reviewsText := doc.Find(".ratings p").Text()
	reviews, _ := strconv.Atoi(strings.Fields(reviewsText)[0])
	product.Reviews = reviews

	doc.Find(".swatches button").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		priceText := s.AttrOr("data-price", "")
		price, _ := strconv.ParseFloat(strings.Replace(priceText, "$", "", 1), 64)

		product.Storages = append(product.Storages, model.StorageOption{
			Size:  size,
			Price: price,
		})
	})

	return product, nil
}
