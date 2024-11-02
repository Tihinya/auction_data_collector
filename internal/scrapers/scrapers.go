package scrapers

type Scraper interface {
	URLsExtract() ([]map[string]string, error)
	DataExtract() ([]map[string]string, error)
}
