package main

import (
	"auction-data-collector/internal/config"
	"fmt"
	"log"
)

func main() {

	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	fmt.Println(cfg)
}
