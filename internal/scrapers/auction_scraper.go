package scrapers

import (
	"context"
	"fmt"
	"strings"

	"auction-data-collector/internal/processors"
	"auction-data-collector/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

type AuctionScraper struct {
	UserAgent    string
	TimeInterval int
	Ctx          context.Context
}

func (s *AuctionScraper) URLsExtract(url, itemSelector, linkSelector, valueSelector string) ([][]string, error) {
	headers := make(map[string]string, 1)
	headers["User-Agent"] = s.UserAgent

	doc, err := utils.FetchHTML(s.Ctx, url, headers, s.TimeInterval)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML from %s: %w", url, err)
	}

	dataItems := make([][]string, 0, 20)

	doc.Find(itemSelector).Each(func(i int, s *goquery.Selection) {
		link, exists := s.Find(linkSelector).Attr("href")
		if !exists || strings.TrimSpace(link) == "" {
			return
		}

		value := strings.TrimSpace(s.Find(valueSelector).Text())

		data := make([]string, 2)
		data[0] = value
		data[1] = link

		dataItems = append(dataItems, data)
	})

	if len(dataItems) == 0 {
		customErr := utils.NewCustomError("URLScraper", "ExtractURLs", "WARNING", "no URLs found on page", nil)
		utils.LogError(customErr)
		return nil, customErr
	}

	return dataItems, nil
}

func (s *AuctionScraper) DataExtract(data []string, itemSelector, labelSelector, valueSelector string) ([]interface{}, error) {
	headers := map[string]string{"User-Agent": s.UserAgent}
	var localItems []interface{}
	var cadastralCodes []string

	doc, err := utils.FetchHTML(s.Ctx, data[1], headers, s.TimeInterval)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML from %s: %w", data[0], err)
	}
	localItems = append(localItems, data[0], data[1])

	doc.Find(itemSelector).Each(func(i int, s *goquery.Selection) {
		label := strings.TrimSpace(s.Find(labelSelector).First().Text())
		value := strings.TrimSpace(s.Find(valueSelector).First().Text())
		if item, ok := processors.ProcessAuctionData(label, value); ok {
			localItems = append(localItems, item)
		}

		if label == "Katastritunnus" {
			value := s.Find("span[data-holder='copy-text']").Text()
			cadastralCodes = append(cadastralCodes, value)
		}
	})

	if len(cadastralCodes) > 0 {
		joinedCodes := strings.Join(cadastralCodes, "; ")
		localItems = append(localItems, joinedCodes)
	} else {
		localItems = append(localItems, "")
	}

	return localItems, nil
}
