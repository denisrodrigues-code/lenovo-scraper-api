package scraper

import (
	"math"
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

	image, _ := doc.Find("img.img-responsive").Attr("src")
	image = "https://webscraper.io" + image

	priceText := strings.TrimSpace(doc.Find("h4.price").Text())
	basePrice, _ := strconv.ParseFloat(strings.Replace(priceText, "$", "", 1), 64)

	product := &model.Product{
		Name:        strings.TrimSpace(doc.Find("h4.title").First().Text()),
		Brand:       "Lenovo",
		Description: doc.Find(".description").Text(),
		BasePrice:   basePrice,
		URL:         url,
		Image:       image,
	}

	doc.Find(".ws-icon-star").Each(func(i int, _ *goquery.Selection) {
		product.Rating++
	})

	reviewsText := doc.Find(".ratings p").Text()
	reviews, _ := strconv.Atoi(strings.Fields(reviewsText)[0])
	product.Reviews = reviews

	multipliers := []float64{1.0, 1.2, 1.5, 1.8}

	doc.Find(".swatches button").Each(func(i int, s *goquery.Selection) {
		if _, disabled := s.Attr("disabled"); disabled {
			return
		}

		size := s.Text()

		multiplier := 1.0
		if i < len(multipliers) {
			multiplier = multipliers[i]
		}

		calculatedPrice := math.Round(product.BasePrice*multiplier*100) / 100

		product.Storages = append(product.Storages, model.StorageOption{
			Size:  size,
			Price: calculatedPrice,
		})
	})

	return product, nil
}
