package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Parse the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %v\n", err)
		return
	}

	// Parse the current URL
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL: %v\n", err)
		return
	}

	// Ensure the current URL is on the same domain as the base URL
	if baseURL.Host != currentURL.Host {
		return
	}

	// Resolve the current URL against the base URL
	resolvedURL := baseURL.ResolveReference(currentURL).String()

	// Normalize the URL
	normalizedURL := strings.TrimSuffix(resolvedURL, "/")

	// Check if the page has already been visited
	if _, visited := pages[normalizedURL]; visited {
		return
	}

	// Fetch the URL
	resp, err := http.Get(normalizedURL)
	if err != nil {
		fmt.Printf("error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Only proceed if the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: status code %d\n", resp.StatusCode)
		return
	}

	// Mark the page as visited
	pages[normalizedURL] = 1

	// Print the URL being crawled
	fmt.Printf("Crawling URL: %s\n", normalizedURL)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return
	}

	// Extract links from the page and recursively crawl them
	links := extractLinks(string(body))
	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}
}

func extractLinks(body string) []string {
	// Dummy implementation for extracting links from the page body
	// In a real implementation, you would parse the HTML and extract href attributes from <a> tags
	return []string{}
}
