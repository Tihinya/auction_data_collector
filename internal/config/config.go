package config

import (
	"encoding/json"
	"os"

	"auction-data-collector/internal/utils"
)

type ScraperConfig struct {
	BaseURL        string `json:"base_url"`
	UpdateInterval int    `json:"update_interval"`
}

type ForestAPIConfig struct {
	BaseURL string `json:"base_url"`
	Timeout int    `json:"timeout"`
}

type StorageConfig struct {
	SpreadsheetID string `json:"spreadsheet_id"`
	WorksheetName string `json:"worksheet_name"`
}

type Config struct {
	Scraper   ScraperConfig   `json:"scraper"`
	ForestAPI ForestAPIConfig `json:"forest_api"`
	Storage   StorageConfig   `json:"storage"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		customErr := utils.NewCustomError("Config", "LoadConfig", "ERROR", "failed to open config file", err)
		utils.LogError(customErr)
		return nil, customErr
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		customErr := utils.NewCustomError("Config", "LoadConfig", "ERROR", "failed to decode config file", err)
		utils.LogError(customErr)
		return nil, customErr
	}

	return &config, nil
}
