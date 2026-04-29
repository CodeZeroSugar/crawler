package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(strings.ToLower(parsedURL.Host+parsedURL.Path), "/"), nil
}
