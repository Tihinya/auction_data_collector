package processor

import (
	"auction-data-collector/internal/scrapers"
	"auction-data-collector/internal/storage"
	"context"
)

type Processor struct {
	scraper *scrapers.AuctionScraper
}

func NewProcessor(scraper *scrapers.AuctionScraper) *Processor {
	return &Processor{
		scraper: scraper,
	}
}

func (p *Processor) ProcessAuctionData(ctx context.Context, url string) error {
	data, err := p.scraper.URLsExtract(url, "div.auction.list-item", "a.auction-hover", "div.price")
	if err != nil {
		return err
	}

	var allData [][]interface{}
	for _, item := range data {
		itemData, err := p.scraper.DataExtract(item, "div.row", "div.label", "div.value")
		if err != nil {
			return err
		}
		allData = append(allData, itemData)
	}

	interfaceValues := make([][]interface{}, len(allData))
	for i, row := range allData {
		interfaceRow := make([]interface{}, len(row))
		for j, cell := range row {
			interfaceRow[j] = cell
		}
		interfaceValues[i] = interfaceRow
	}

	return storage.WriteToSheet(ctx, "Лист1!A2", interfaceValues)
}
