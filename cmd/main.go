package main

import (
	"auction-data-collector/internal/config"
	"auction-data-collector/internal/processor"
	"auction-data-collector/internal/scrapers"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	auctionScraper := &scrapers.AuctionScraper{
		UserAgent:    cfg.Scraper.UserAgent,
		TimeInterval: cfg.Scraper.UpdateInterval,
		Ctx:          ctx,
	}

	proc := processor.NewProcessor(auctionScraper)
	if err := proc.ProcessAuctionData(ctx, cfg.Scraper.TehingukeskusURL); err != nil {
		log.Printf("Error processing auction data: %v", err)
		return
	}
}
