package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

func TestCrawlPage(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<a href="/page1">Page 1</a>`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	baseURL, _ := url.Parse(server.URL)
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
		maxPages:           100,
	}
	cfg.crawlPage(server.URL)

	if len(cfg.pages) != 1 {
		t.Fatalf("expected 1 page, got %d", len(cfg.pages))
	}

	if count, exists := cfg.pages[server.URL]; !exists || count != 1 {
		t.Errorf("expected page %s to be visited once, got %d", server.URL, count)
	}
}

func TestCrawlPageNotFound(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	baseURL, _ := url.Parse(server.URL)
	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
		maxPages:           100,
	}
	cfg.crawlPage(server.URL)

	if len(cfg.pages) != 0 {
		t.Fatalf("expected 0 pages, got %d", len(cfg.pages))
	}
}
