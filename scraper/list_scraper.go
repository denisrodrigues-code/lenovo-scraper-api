package scraper

import (
	"strings"

	"github.com/denisrodrigues-code/lenovo-scraper-api/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetLenovoProductLinks(baseURL string) ([]string, error) {
	client := utils.NewClient()
	var links []string

	resp, err := client.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	doc.Find(".thumbnail").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title").Text()
		if strings.Contains(strings.ToLower(title), "lenovo") {
			href, _ := s.Find(".title").Attr("href")
			links = append(links, "https://webscraper.io"+href)
		}
	})

	return links, nil
}
