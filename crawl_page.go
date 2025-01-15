package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	if cfg.mu == nil {
		cfg.mu = &sync.Mutex{}
	}
	if cfg.concurrencyControl == nil {
		cfg.concurrencyControl = make(chan struct{}, 10)
	}
	if cfg.wg == nil {
		cfg.wg = &sync.WaitGroup{}
	}

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL: %v\n", err)
		return
	}

	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	resolvedURL := cfg.baseURL.ResolveReference(currentURL).String()
	normalizedURL := strings.TrimSuffix(resolvedURL, "/")

	cfg.mu.Lock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.mu.Unlock()
		return
	}
	cfg.pages[normalizedURL] = 1
	fmt.Printf("Crawling Page: %s\n", normalizedURL)
	cfg.mu.Unlock()

	resp, err := http.Get(normalizedURL)
	if err != nil {
		fmt.Printf("error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error: status code %d\n", resp.StatusCode)
		cfg.mu.Lock()
		delete(cfg.pages, normalizedURL)
		cfg.mu.Unlock()
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return
	}

	links := extractLinks(string(body))
	// fmt.Printf("Found %d links on %s\n", len(links), normalizedURL)
	for _, link := range links {
		cfg.wg.Add(1)
		go func(link string) {
			fmt.Printf("Crawling link: %s\n", link)
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(link)
			<-cfg.concurrencyControl
		}(link)
	}
}

func extractLinks(body string) []string {
	var links []string
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		fmt.Printf("error parsing HTML: %v\n", err)
		return links
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}
