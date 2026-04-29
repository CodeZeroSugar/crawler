package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) (string, error) {
	htmlReader := strings.NewReader(html)
	htmlDoc, err := goquery.NewDocumentFromReader(htmlReader)
	if err != nil {
		return "", err
	}

	docSelection := htmlDoc.Selection

	var header string
	for i := 1; i <= 2; i++ {
		header = docSelection.Find(fmt.Sprintf("h%d", i)).Text()
		header = strings.TrimSpace(header)
		if header != "" {
			break
		}
	}

	return header, nil
}
