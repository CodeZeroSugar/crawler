package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := make([]string, 0)
	for key := range pages {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	sortedData := make([]PageData, 0)
	for _, key := range keys {
		sortedData = append(sortedData, pages[key])
	}

	data, err := json.MarshalIndent(sortedData, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(filename, data, 0o644)

	return nil
}
