package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
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
		baseURL:            parsedBase,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}

	crawlerCfg.wg.Add(1)
	go crawlerCfg.crawlPage(baseURL)
	crawlerCfg.wg.Wait()

	for k, v := range crawlerCfg.pages {
		fmt.Printf("URL: %s, Data: %v\n\n", k, v)
	}
}
