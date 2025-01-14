package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		baseURL := args[0]
		fmt.Printf("starting crawl of: %s\n", baseURL)
		pages := make(map[string]int)
		crawlPage(baseURL, baseURL, pages)
		for url, count := range pages {
			fmt.Printf("URL: %s, Count: %d\n", url, count)
		}
	}
}
