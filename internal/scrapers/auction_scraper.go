package scrapers

import (
	"context"
	"fmt"
	"strings"

	"auction-data-collector/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

type AuctionScraper struct {
	UserAgent    string
	TimeInterval int
	Ctx          context.Context
}

func (s *AuctionScraper) URLsExtract(url, itemSelector, linkSelector, valueSelector string) ([][]string, error) {
	headers := map[string]string{"User-Agent": s.UserAgent}

	doc, err := utils.FetchHTML(s.Ctx, url, headers, s.TimeInterval)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML from %s: %w", url, err)
	}

	var dataItems [][]string

	doc.Find(itemSelector).Each(func(i int, s *goquery.Selection) {

		link, exists := s.Find(linkSelector).Attr("href")
		if !exists || strings.TrimSpace(link) == "" {
			return
		}

		value := strings.TrimSpace(s.Find(valueSelector).Text())

		var data []string
		data = append(data, link, value)

		dataItems = append(dataItems, data)
	})

	if len(dataItems) == 0 {
		customErr := utils.NewCustomError("URLScraper", "ExtractURLs", "WARNING", "no URLs found on page", nil)
		utils.LogError(customErr)
		return nil, customErr
	}

	return dataItems, nil
}

func (s *AuctionScraper) DataExtract(data []string, itemSelector, labelSelector, valueSelector string) ([]map[string]string, error) {
	headers := map[string]string{"User-Agent": s.UserAgent}
	var localItems []map[string]string

	item := map[string]string{data[0]: data[1]}
	localItems = append(localItems, item)

	doc, err := utils.FetchHTML(s.Ctx, data[0], headers, s.TimeInterval)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML from %s: %w", data[0], err)
	}

	doc.Find(itemSelector).Each(func(i int, s *goquery.Selection) {
		lable := strings.TrimSpace(s.Find(labelSelector).Text())
		if lable == "Lõpeb" {
			value := strings.TrimSpace(s.Find(valueSelector).Text())
			item := map[string]string{"Data": value}
			localItems = append(localItems, item)
		} else if lable == "Kogupindala" {
			value := strings.TrimSpace(s.Find(valueSelector).Text())
			item := map[string]string{"Size": value}
			localItems = append(localItems, item)
		} else if lable == "KatastritunnusPindala" {
			value := strings.TrimSpace(s.Find("span[data-holder='copy-text']").Text())
			item := map[string]string{"Location": value}
			localItems = append(localItems, item)
		}

	})

	return localItems, nil
}
