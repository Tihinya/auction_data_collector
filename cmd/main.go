package main

import (
	"auction-data-collector/internal/config"
	"auction-data-collector/internal/scrapers"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	auctionScraper := scrapers.AuctionScraper{
		UserAgent:    cfg.Scraper.UserAgent,
		TimeInterval: cfg.Scraper.UpdateInterval,
		Ctx:          ctx,
	}
	data, err := auctionScraper.URLsExtract(cfg.Scraper.TehingukeskusURL, "div.auction.list-item", "a.auction-hover", "div.price")
	if err != nil {
		log.Printf("Error scraping data: %v", err)
		return
	}

	var allData [][]map[string]string
	for _, item := range data {
		itemData, err := auctionScraper.DataExtract(item, "div.row", "div.label", "div.value")
		if err != nil {
			log.Printf("Error scraping data: %v", err)
			return
		}
		allData = append(allData, itemData)
	}
	fmt.Println(allData)
}
