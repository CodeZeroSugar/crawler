package main

import (
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	baseURL, _ := url.Parse(pageURL)
	heading, _ := getHeadingFromHTML(html)
	firstParagraph, _ := getFirstParagraphFromHTML(html)
	outLinks, _ := getURLsFromHTML(html, baseURL)
	imgURLs, _ := getImagesFromHTML(html, baseURL)

	return PageData{
		URL:            baseURL.String(),
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outLinks,
		ImageURLs:      imgURLs,
	}
}
