# Auction Data Collector

## Overview
Auction Data Collector is an automated system for collecting and analyzing auction data from tehingukeskus.ee and enriching it with forest registry data. The system performs regular price updates, monitors auction endings, and stores the collected data in Google Sheets for further analysis.

## Features
- 🔄 Automated data collection from tehingukeskus.ee every 5 minutes
- 🌲 Forest registry data integration via API
- 📊 Google Sheets integration for data storage
- ⏰ Automatic price updates and auction end monitoring
- 🔍 Cadastral number-based forest data lookup
- 🛡️ Error handling and retry mechanisms

## Prerequisites
- Go 1.16 or higher
- Google Cloud Platform account (for Sheets API)
- Access to Forest Registry API
- Valid API credentials for both services

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/auction-data-collector.git
cd auction-data-collector
```

2. Set up your Google Sheets API credentials:
   - Create a project in Google Cloud Console
   - Enable Google Sheets API
   - Create service account credentials
   - Download the credentials JSON file
   - Place it in the `config` directory as `google_credentials.json`

## Configuration

1. Configure your settings in `config/config.yaml`:
```yaml
scraper:
  base_url: "https://www.tehingukeskus.ee/oksjonid-koik"
  update_interval: 300  # 5 minutes in seconds

forest_api:
  base_url: "https://register.metsad.ee/portaal/api/rest"
  timeout: 30

storage:
  spreadsheet_id: "your-spreadsheet-id"
  worksheet_name: "Auction Data"
```

2. Set up environment variables in `.env`:
```
GOOGLE_CREDENTIALS_PATH=config/google_credentials.json
FOREST_API_KEY=your-api-key
```

## Usage

### Makefile Commands

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Clean up the binary
make clean
```

1. Run the main collection service:
```bash
./auction-data-collector
```

2. Run specific components:
```bash
# Run only the scraper
go run ./internal/scrapers/auction_scraper.go

# Run only the forest data collector
go run ./internal/scrapers/forest_api.go
```

3. Monitor the logs in `logs/app.log`

## Project Structure
The project follows the standard Golang project layout, with the following directories:

```
auction-data-collector/
├── cmd/
│   └── main.go           # Main entry point of the application
│
├── internal/
│   ├── scrapers/
│   │   ├── auction_scraper.go   # Tehingukeskus.ee scraper
│   │   ├── forest_api.go        # Forest registry API client
│   │   └── scrapers.go          # Shared scraper functionality
│   │
│   ├── processors/
│   │   ├── data_processor.go    # Data cleaning and transformation
│   │   ├── update_service.go    # 5-minute update checker
│   │   └── processors.go        # Shared processor functionality
│   │
│   ├── storage/
│   │   ├── sheets_client.go     # Google Sheets integration
│   │   ├── models.go            # Data models/schemas
│   │   └── storage.go           # Shared storage functionality
│   │
│   ├── config/
│   │   └── config.go            # Configuration management
│   │
│   └── utils/
│       ├── errors.go            # Custom errors
│       └── utils.go             # Shared utility functions
│
├── pkg/
│   ├── scrapers/
│   │   ├── auction_scraper_test.go
│   │   ├── forest_api_test.go
│   │   └── scrapers_test.go
│   │
│   ├── processors/
│   │   ├── data_processor_test.go
│   │   ├── update_service_test.go
│   │   └── processors_test.go
│   │
│   └── storage/
│       ├── sheets_client_test.go
│       └── storage_test.go
│
├── config/
│   ├── config.yaml      # Application configuration
│   └── logging.yaml     # Logging configuration
│
├── docs/
│   ├── README.md        # Project documentation
│   ├── API.md           # API documentation
│   └── SETUP.md         # Setup instructions
│
├── scripts/
│   ├── setup.sh         # Setup script
│   └── run.sh           # Run script
│
├── .env.example         # Example environment variables
├── .gitignore           # Git ignore file
├── go.mod               # Go module definition
└── README.md            # Project README
```

## Development

### Running Tests
```bash
make test
```

### Adding New Features
1. Create a new branch for your feature
2. Write tests in the appropriate `pkg` directory
3. Implement your feature in the `internal` directory
4. Update documentation as needed
5. Submit a pull request

## Troubleshooting
Refer to the Troubleshooting section in the original README.

## Contributing
Refer to the Contributing section in the original README.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contact
Your Name - your.email@example.com
Project Link: https://github.com/yourusername/auction-data-collector

## Acknowledgments
- tehingukeskus.ee for auction data
- Forest Registry for forest data API
- Google Sheets API for data storage capabilities