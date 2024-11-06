package storage

import (
	"auction-data-collector/internal/utils"
	"context"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func WriteToSheet(ctx context.Context, writeRange string, values [][]interface{}) error {
	err := godotenv.Load(".env")
	if err != nil {
		customErr := utils.NewCustomError("DataProcessor", "WriteToSheet", "FATAL", "failed to load .env file", err)
		utils.LogError(customErr)
		return customErr
	}

	jsonKeyFile := os.Getenv("GOOGLE_CREDENTIALS_PATH")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")

	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(jsonKeyFile))
	if err != nil {
		customErr := utils.NewCustomError("DataProcessor", "WriteToSheet", "FATAL", "unable to retrieve Sheets client", err)
		utils.LogError(customErr)
		return customErr
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		customErr := utils.NewCustomError("DataProcessor", "WriteToSheet", "FATAL", "unable to write data to sheet", err)
		utils.LogError(customErr)
		return customErr
	}

	return nil
}
