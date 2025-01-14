package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	normalizedURL := strings.TrimPrefix(parsedURL.Host+parsedURL.Path, "www.")
	return normalizedURL, nil
}
