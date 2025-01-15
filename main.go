package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	const (
		defaultMaxConcurrency = 10
		defaultMaxPages       = 100
	)

	if len(os.Args) < 2 {
		panic("Usage: ./crawler <baseURL> [maxConcurrency] [maxPages]")
	}

	rawBaseURL := os.Args[1]

	maxConcurrency := defaultMaxConcurrency
	if len(os.Args) > 2 {
		var err error
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic("Invalid value for maxConcurrency")
		}
	}

	maxPages := defaultMaxPages
	if len(os.Args) > 3 {
		var err error
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic("Invalid value for maxPages")
		}
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		panic(err)
	}

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("Starting crawl at base URL: %s\n", rawBaseURL)

	// Initial crawl without spawning a goroutine
	cfg.crawlPage(rawBaseURL)

	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}
