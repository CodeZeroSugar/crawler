package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("syntax: <url> <maxPages> <maxConcurrency>")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("syntax: <url> <maxPages> <maxConcurrency>")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid value for maxConcurrency")
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid value for maxPages")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Printf("starting crawl of: %s\n", baseURL)
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("failed to parse baseURL: %s", baseURL)
	}

	crawlerCfg := config{
		pages:              make(map[string]PageData, 0),
		maxPages:           maxPages,
		baseURL:            parsedBase,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	crawlerCfg.wg.Add(1)
	go crawlerCfg.crawlPage(baseURL)
	crawlerCfg.wg.Wait()

	for k, v := range crawlerCfg.pages {
		fmt.Printf("URL: %s, Data: %v\n\n", k, v)
	}

	err = writeJSONReport(crawlerCfg.pages, "report.json")
	if err != nil {
		fmt.Println("failed to write json report.")
	} else {
		fmt.Println("json report generated")
	}
}
