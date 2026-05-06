package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return []string{}, err
	}

	urls := make([]string, 0)

	htmlDoc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		attrib, _ := s.Attr("src")
		if strings.TrimSpace(attrib) != "" {
			urls = append(urls, strings.TrimSpace(attrib))
		}
	})

	if len(urls) == 0 {
		return []string{}, nil
	}

	finalUrls := make([]string, 0)
	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			return []string{}, err
		}
		absURL := baseURL.ResolveReference(parsedURL)
		finalUrls = append(finalUrls, absURL.String())
	}

	return finalUrls, nil
}
