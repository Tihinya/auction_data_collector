package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchHTML(ctx context.Context, url string, headers map[string]string, updateInterval int) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, NewCustomError("Utils", "FetchHTML", "ERROR", "failed to create new HTTP request", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Duration(updateInterval) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		customErr := NewCustomError("Utils", "FetchHTML", "ERROR", "failed to fetch HTML", err)
		LogError(customErr)
		return nil, customErr
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		customErr := NewCustomError("Utils", "FetchHTML", "ERROR", "failed to parse HTML document", err)
		LogError(customErr)
		return nil, customErr
	}

	return doc, nil
}
