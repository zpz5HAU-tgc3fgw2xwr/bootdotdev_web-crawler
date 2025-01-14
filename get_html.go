package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	if contentType := resp.Header.Get("Content-Type"); contentType != "text/html" {
		return "", fmt.Errorf("error: content type is %s, not text/html", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
