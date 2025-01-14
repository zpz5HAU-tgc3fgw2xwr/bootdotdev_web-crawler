package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawlPage(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<a href="/page1">Page 1</a>`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	pages := make(map[string]int)
	crawlPage(server.URL, server.URL, pages)

	if len(pages) != 1 {
		t.Fatalf("expected 1 page, got %d", len(pages))
	}

	if count, exists := pages[server.URL]; !exists || count != 1 {
		t.Errorf("expected page %s to be visited once, got %d", server.URL, count)
	}
}

func TestCrawlPageNotFound(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	pages := make(map[string]int)
	crawlPage(server.URL, server.URL, pages)

	if len(pages) != 0 {
		t.Fatalf("expected 0 pages, got %d", len(pages))
	}
}
