package scraper

import (
	"fmt"
	"strconv"
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

	lastPage := 1
	doc.Find(".pagination li a").Each(func(i int, s *goquery.Selection) {
		pageText := strings.TrimSpace(s.Text())
		page, err := strconv.Atoi(pageText)
		if err == nil && page > lastPage {
			lastPage = page
		}
	})

	for page := 1; page <= lastPage; page++ {
		pageURL := fmt.Sprintf("%s?page=%d", baseURL, page)

		resp, err := client.Get(pageURL)
		if err != nil {
			continue
		}

		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()

		doc.Find(".thumbnail").Each(func(i int, s *goquery.Selection) {
			title := s.Find(".title").Text()
			if strings.Contains(strings.ToLower(title), "lenovo") {
				href, _ := s.Find(".title").Attr("href")
				links = append(links, "https://webscraper.io"+href)
			}
		})
	}

	return links, nil
}
