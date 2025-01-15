package main

import (
	"fmt"
	"sort"
)

type pageReport struct {
	url   string
	count int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	var report []pageReport
	for url, count := range pages {
		report = append(report, pageReport{url, count})
	}

	sort.Slice(report, func(i, j int) bool {
		if report[i].count == report[j].count {
			return report[i].url < report[j].url
		}
		return report[i].count > report[j].count
	})

	for _, entry := range report {
		fmt.Printf("Found %d internal links to %s\n", entry.count, entry.url)
	}
}
