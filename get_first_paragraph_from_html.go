package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) (string, error) {
	htmlReader := strings.NewReader(html)
	htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return "", err
	}

	selection := htmlDoc.Selection

	mainParagraph := strings.TrimSpace(selection.Find("main").Find("p").First().Text())
	if mainParagraph == "" {
		return strings.TrimSpace(selection.Find("p").First().Text()), err
	}

	return mainParagraph, nil
}
