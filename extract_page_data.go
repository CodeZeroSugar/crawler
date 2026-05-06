package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) (PageData, error) {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}, err
	}

	heading, err := getHeadingFromHTML(html)
	if err != nil {
		return PageData{}, err
	}

	firstParagraph, err := getFirstParagraphFromHTML(html)
	if err != nil {
		return PageData{}, err
	}

	outLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		return PageData{}, err
	}

	imgURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		return PageData{}, err
	}

	return PageData{
		URL:            baseURL.String(),
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outLinks,
		ImageURLs:      imgURLs,
	}, nil
}
