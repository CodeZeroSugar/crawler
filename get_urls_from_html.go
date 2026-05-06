package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return []string{}, err
	}

	urls := make([]string, 0)

	htmlDoc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		attrib, _ := s.Attr("href")
		urls = append(urls, strings.TrimSpace(attrib))
	})

	finalUrls := make([]string, 0)
	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			return []string{}, err
		}
		finalUrls = append(finalUrls, parsedURL.String())
	}

	return finalUrls, nil
}
