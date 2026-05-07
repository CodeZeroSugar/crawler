package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedURL]; ok {
		return false
	}

	cfg.pages[normalizedURL] = PageData{}
	return true
}

func (cfg *config) checkPageLength() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) >= cfg.maxPages {
		return true
	}
	return false
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.checkPageLength() {
		return
	}

	parsedCurrent, _ := url.Parse(rawCurrentURL)
	if cfg.baseURL.Hostname() != parsedCurrent.Hostname() {
		return
	}

	normalizedCurrent, _ := normalizeURL(rawCurrentURL)

	if !cfg.addPageVisit(normalizedCurrent) {
		return
	}

	fmt.Printf("crawling %s\n", normalizedCurrent)
	htmlText, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("getHTML error: %s\n", err)
	}

	cfg.pages[normalizedCurrent] = extractPageData(htmlText, cfg.baseURL.String())

	urls, _ := getURLsFromHTML(htmlText, cfg.baseURL)

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
